package tests

import (
	"API_Gateway/config"
	"API_Gateway/tests/helpers"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestLoadConfig tests the LoadConfig function
func TestLoadConfig(t *testing.T) {
	// change os path to the root
	os.Chdir("..")
	cfg := config.LoadConfig()

	helpers.Action("Loading config...")
	assert.Equal(t, ":8080", cfg.ServerPort, "ServerPort should be :8080")
	assert.Equal(t, "localhost:50051", cfg.UserServiceURL, "UserServiceURL should be localhost:50051")
	assert.NotNil(t, cfg)
	helpers.Result(fmt.Sprintf(`Config loaded successfully!
			ServerPort: %v 
			UserServiceURL: %v`, cfg.ServerPort, cfg.UserServiceURL))
}