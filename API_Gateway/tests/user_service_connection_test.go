package tests

import (
	"testing"
	"net/http"
	"io"
	"fmt"
	"API_Gateway/tests/helpers"
	
	"github.com/stretchr/testify/assert"
)

// type MockUserServiceClient struct {
// 	mock.Mock
// }


func TestAuthUserEndpoint(t *testing.T) {
	serverURL := "http://localhost:8080"

	helpers.Title("This is testing on LIVE server!")
	
	helpers.Action("Creating request to the server...")
	req, err := http.NewRequest("GET", serverURL+"/rest/user/login/shine", nil)
	assert.Nil(t, err)
	req.Header.Add("email", "test@gmail.com")
	req.Header.Add("password", "1234")
	helpers.Result(fmt.Sprint(`Request successfully created!
		target URL: `, serverURL+`/rest/user/login/shine
		Headers: `, req.Header))

	helpers.Action("Sending request to the server...")
	client := &http.Client{}
	resp, err := client.Do(req)
	assert.Nil(t, err)
	defer resp.Body.Close()		// defer means this line will be executed at the end of the function
	helpers.Result("Request sent successfully!")

	helpers.Action("Reading response from the server...")
	responseBody, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	expected := `{"valid":true}`
	assert.Equal(t, expected, string(responseBody))
	helpers.Result(fmt.Sprintf(`Response received successfully!
		Expected: %v
		Actual: %v`, expected, string(responseBody)))
}