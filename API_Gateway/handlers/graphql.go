package handlers

import (
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var graphql_server = "http://graphql_gateway:10070"

// HandleGraphql handles the GraphQL requests and routes them to the graphql server
func HandleGraphql(c *gin.Context) {
	// Routing the request to the GraphQL server on port 10070
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create a new request to the GraphQL server
	req, err := http.NewRequest("POST", graphql_server, strings.NewReader(string(body)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Set the headers for the request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.GetHeader("Authorization"))

	// Create a new client
	client := &http.Client{}

	// Send the request to the GraphQL server
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Read the response from the GraphQL server
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Send the response back to the client
	c.Data(http.StatusOK, "application/json", body)
}
