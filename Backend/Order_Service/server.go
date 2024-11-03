package main

import (
	"Order_Service/db"
	"Order_Service/graph"
	"context"
	"log"
	"net/http"
	"os"
	"Order_Service/internal/jwt"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "10072"

func main() {
	db.Connect()
	defer db.Close()
	DB := db.GetDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DB: DB}}))

	// Middleware to extract headers and add to context
	headerMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extracting the token
			token := r.Header.Get("Authorization")
			
			ctx := context.WithValue(r.Context(), jwt.TokenKey, token)
			log.Printf("Token outside: %s", token)
			// Call the next handler with the new context
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
	
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", headerMiddleware(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
