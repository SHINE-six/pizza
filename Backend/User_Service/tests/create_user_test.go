package tests

// import (
// 	"User_Service/internal/services"
// 	"User_Service/tests/helpers"
// 	"log"

// 	// "User_Service/internal/handlers"
// 	"User_Service/proto"
// 	"os"
// 	"testing"

// 	// "fmt"

// 	// "net/http"

// 	"github.com/stretchr/testify/assert"
// 	// "github.com/stretchr/testify/mock"
// )

// // func TestPostgresqlDB(t *testing.T) {
// // 	// Change os path to the server main.go file
// // 	os.Chdir("..")
// // 	assert := assert.New(t)

// // 	helpers.Title("Testing the function *addVerificationTokenToDB*")

// // 	helpers.Action("Call to function to add verification token to the database")
// // 	resp, err := services.addVerificationTokenToDB("test@mail.com", "1234")
// // 	assert.Equal(0, resp)
// // 	assert.Nil(err)
// // 	helpers.Result("Verification token added to the database successfully")
// // }

// func TestCreateUser(t *testing.T) {
// 	// Change os path to the server main.go file
// 	os.Chdir("..")
// 	assert := assert.New(t)

// 	helpers.Title("Testing the function *AssureAllRequiredInformation*")

// 	helpers.Action("Call to function with all required information")
// 	dataToSend := &__.Credential{
// 		Email: "desmondfoo0610@gmail.com",
// 		Password: "1234aA@6",
// 		Username: "shine",
// 	}
// 	resp, err := services.AssureAllRequiredInformation(dataToSend)
// 	assert.Equal(0, resp)
// 	assert.Equal("All required information is present", err)
// 	helpers.Result("User Created Successfully")


// 	// helpers.Action("Call to function with missing required information, example missing Username")
// 	// dataToSend = &__.Credential{
// 	// 	Email: "test@gmail.com",
// 	// 	Password: "1234",
// 	// 	Username: "",
// 	// }
// 	// resp, err = services.AssureAllRequiredInformation(dataToSend)
// 	// assert.Equal(1, resp)
// 	// assert.Equal("Missing username", err)
// 	// helpers.Result("Create User with not all required information")
// }

// func TestVerifyStaff(t *testing.T) {
// 	// Change os path to the server main.go file
// 	os.Chdir("..")
// 	assert := assert.New(t)

// 	helpers.Title("Testing the function *VerifyStaff*")

// 	helpers.Action("Call to function to verify staff")
// 	data, resp, err := services.VerifyStaff("potato@mail.com")
// 	log.Println("data: ", data)
// 	assert.Equal(0, resp)
// 	assert.Nil(err)
// 	helpers.Result("Staff Verified Successfully")
// }


// // Test the function CreateJWTToken
// func TestCreateJWTToken(t *testing.T) {
// 	// Change os path to the server main.go file
// 	os.Chdir("..")
	
// 	assert := assert.New(t)
	
// 	helpers.Title("Testing the function *CreateJWTToken*")

// 	helpers.Action("Call to function to create a JWT token")
// 	token, err := services.CreateJWTToken("desmondfoo0610@gmail.com")
// 	assert.Nil(err)
// 	assert.NotEmpty(token)
// 	log.Println("Token: ", token)
// 	helpers.Result("JWT token created successfully")
// } 



// // func TestPasswordHasing(t *testing.T) {
// // 	os.Chdir("../cmd/server")
// // 	assert := assert.New(t)

// // 	helpers.Title("Testing the function *HashPassword*")

// // 	helpers.Action("Call to function to hash a password")
// // 	hashedPassword, err := services.PasswordHashing("Grape@123")
// // 	assert.Nil(err)
// // 	assert.NotEmpty(hashedPassword)
// // 	log.Println("Hashed Password: ", hashedPassword)
// // 	helpers.Result("Password hashed successfully")
// // }