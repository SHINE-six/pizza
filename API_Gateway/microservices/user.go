package microservices

import (
	"API_Gateway/config"
	"API_Gateway/proto"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Setup
var cfg = config.LoadConfig()

/* AuthUser is a function that handles the authentication of a user.
It retrieves the email, password, and username from the request body
The email and password are retrieved from the request headers because they are typically sent as part of the authentication process.
The username is retrieved from the request parameters because it is often included in the URL path.
This function sets up a connection to the user service using gRPC, contacts the user service to authenticate the user, and retrieves the response.
Finally, it returns a JSON response indicating whether the user is valid or not. */

type AuthCredential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AuthUser(c *gin.Context) {
	log.Println("AuthUser")
	// Implementation
	var authCredential AuthCredential

	if err := c.ShouldBindJSON(&authCredential); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// set up connection to user service
	conn, err := grpc.NewClient(cfg.UserServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := __.NewUserServiceClient(conn)

	// contact the user service and print out its response
	var jwtToken metadata.MD
	_, err = client.AuthUser(context.Background(), &__.AuthCredential{Email: authCredential.Email, Password: authCredential.Password}, grpc.Header(&jwtToken))
	if err != nil {
		log.Printf("could not greet: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "User authenticated successfully", "jwt": jwtToken["jwt_set_cookie"]})
	}
}

type CreateUserRequestCredential struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(c *gin.Context) {
	// Implementation
	var createUserRequestCredential CreateUserRequestCredential

	if err := c.ShouldBindJSON(&createUserRequestCredential); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// set up connection to user service
	conn, err := grpc.NewClient(cfg.UserServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer conn.Close()
	client := __.NewUserServiceClient(conn)

	// contact the user service and print out its response
	_, err = client.CreateUser(context.Background(), &__.Credential{Username: createUserRequestCredential.Username, Email: createUserRequestCredential.Email, Password: createUserRequestCredential.Password})
	if err != nil {
		// check rpc error code
		// if rpc error code is already exists, return 409
		// if rpc error code is internal, return 500
		
		if status.Code(err) == codes.AlreadyExists {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
	}
}

func VerifyUser(c *gin.Context) {
	// Implementation
	email := c.Query("email")
	token := c.Query("token")
	log.Printf("Email: %v, Token: %v\n", email, token)

	// set up connection to user service
	conn, err := grpc.NewClient(cfg.UserServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := __.NewUserServiceClient(conn)

	// contact the user service and print out its response
	_, err = client.VerifyUser(context.Background(), &__.VerifyCredential{Email: email, Token: token})
	if err != nil {
		// check rpc error code
		// if rpc error code is not found, return 404
		// if rpc error code is internal, return 500
		if status.Code(err) == codes.NotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "User verified successfully"})
	}
}

func VerifyStaff(c *gin.Context) {
	// Implementation
	email := c.Query("email")
	
	// set up connection to user service
	conn, err := grpc.NewClient(cfg.UserServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := __.NewUserServiceClient(conn)

	// contact the user service and print out its response
	data, err := client.VerifyStaff(context.Background(), &__.EmailCredential{Email: email})
	if err != nil {
		// check rpc error code
		// if rpc error code is not found, return 404
		// if rpc error code is internal, return 500
		if status.Code(err) == codes.NotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Staff verified successfully", "data": data})
	}
}

