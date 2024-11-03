package services

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
	"User_Service/config"
	proto "User_Service/proto"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	log.Println("Initializing the postgresql_db.go file")
	cfg := config.LoadDBConfig()
	
	var err error
	db, err = sql.Open("postgres", cfg.PostgresDatabaseURL)
	if err != nil {
		log.Fatalf("Error opening the database connection: %v", err)
	}
	log.Println("Database connection opened successfully")
}

// CloseDB cleanly closes the database connection
func CloseDB() {
    if err := db.Close(); err != nil {
        log.Printf("Error closing database: %v", err)
    }
}

func addVerificationTokenToDB(email string, verificationToken string, username string, password string) (int, error) {
	log.Println("Adding the email and verification token to the database")
	// Add the email and verification token to the database table - verification_tokens
	// The table should have the following columns: email, token, username, password, created_at, updated_at

	// Check if the email already exists in the customers database
	var cus_emailExists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM public.customers WHERE email = $1)", email).Scan(&cus_emailExists)
	if err != nil {
		fmt.Println("Error checking if the email exists in the database: ", err)
		return 2, err
	}
	if cus_emailExists {
		fmt.Println("Email already exists in the database")
		return 1, errors.New("email already exists in the database")
	}

	// Check if the email already exists in the verification_tokens database
	var ver_emailExists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM public.verification_tokens WHERE email = $1)", email).Scan(&ver_emailExists)
	if err != nil {
		fmt.Println("Error checking if the email exists in the database: ", err)
		return 2, err
	}
	if ver_emailExists {
		fmt.Println("Email already exists in the verification_tokens database, updating the token")
		// Update the token in the database
		statusCode, err := updateVerificationToken(email, verificationToken)
		if statusCode != 0 {
			fmt.Println("Error updating the token in the database: ", err)
			return 2, err
		}
		return 0, nil
	}
	// hash the password
	passwordHash, _ := passwordHashing(password)
	_, err = db.Exec("INSERT INTO verification_tokens (email, token, username, password) VALUES ($1, $2, $3, $4)", email, verificationToken, username, passwordHash)
	if err != nil {
		fmt.Println("Error adding the email and verification token to the database: ", err)
		return 2, err
	}
	return 0, nil
}

func updateVerificationToken(email string, verificationToken string) (int, error) {
	fmt.Println("Updating the verification token in the database")
	// Update the token in the database and the time for updated_at
	_, err := db.Exec("UPDATE public.verification_tokens SET token = $1, updated_at = NOW() WHERE email = $2", verificationToken, email)
	if err != nil {
		fmt.Println("Error updating the token in the database: ", err)
		return 1, err
	}
	return 0, nil
}

func VerifyUser(email string, token string) (int, error) {
	fmt.Println("Verifying the user with the email and token")
	fmt.Println("Email: ", email, "\nToken: ", token)

	// Define a struct to hold query results
	type VerificationStatus struct {
		EmailExists bool
		TokenValid  bool
		TokenExpired bool
		UserInfo    struct {
			Email    string
			Username string
			Password string
		}
	}

	var result VerificationStatus

    // Query to check email existence, token validity, and expiration
	err := db.QueryRow(`
        SELECT 
            EXISTS(SELECT 1 FROM public.verification_tokens WHERE email = $1) AS EmailExists,
            EXISTS(SELECT 1 FROM public.verification_tokens WHERE email = $1 AND token = $2) AS TokenValid,
            EXISTS(SELECT 1 FROM public.verification_tokens WHERE email = $1 AND token = $2 AND updated_at < NOW() - INTERVAL '3 minutes') AS TokenExpired
    `, email, token).Scan(&result.EmailExists, &result.TokenValid, &result.TokenExpired)
    
	if err != nil {
        fmt.Println("Internal server error: ", err)
        return 2, err
    }

    if !result.EmailExists {
        fmt.Println("User does not exist")
        return 1, errors.New("user does not exist")
    } else if !result.TokenValid {
        fmt.Println("Invalid token")
        return 1, errors.New("invalid token")
    } else if result.TokenExpired {
        fmt.Println("Token has expired")
        return 1, errors.New("token expired")
    }
	fmt.Println("The token matches the stored token")



	return 0, nil
}

func DeleteVerificationToken(email string, token string) (int, error) {
	fmt.Println("Deleting the verification token from the database")
	// Delete the token from the database
	_, err := db.Exec("DELETE FROM public.verification_tokens WHERE email = $1 AND token = $2", email, token)
	if err != nil {
		fmt.Println("Error deleting the token from the database: ", err)
		return 2, err
	}
	return 0, nil
}

func AddUserToDB(email string, token string) (int, error) {
	fmt.Println("Adding the user to the database")
	// Take from the verification_tokens table
	var username, password string
	err := db.QueryRow("SELECT username, password FROM public.verification_tokens WHERE email = $1 AND token = $2", email, token).Scan(&username, &password)
	if err != nil {
		fmt.Println("Error getting the username and password from the database: ", err)
		return 2, err
	}
	// Add the user to the database
	_, err = db.Exec("INSERT INTO customers (email, username, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $4)", email, username, password, time.Now())
	if err != nil {
		fmt.Println("Error adding the user to the database: ", err)
		return 2, err
	}
	return 0, nil
}

func AuthenticateUser(email string, password string) (status int, success bool, err error) {
	log.Println("Authenticating the user with the email and password")

	// First, check if the user exists
	var userExists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM public.customers WHERE email = $1)", email).Scan(&userExists)
	if err != nil {
		fmt.Println("Error checking if the user exists: ", err)
		return 1, false, err
	}
	if !userExists {
		fmt.Println("User does not exist")
		return 0, false, errors.New("user does not exist")
	}

	// If the user exists, check if the password matches
	var passwordFromDB string
	// First, hash the password
	err = db.QueryRow("SELECT (password) FROM public.customers WHERE email = $1", email).Scan(&passwordFromDB)
	if err != nil {
		log.Println("Error checking if the password is correct: ", err)
		return 1, false, err
	}
	result := verifyPasswordHash(password, passwordFromDB)
	if !result {
		log.Println("Email and password do not match")
		return 0, false, errors.New("email and password do not match")
	}
	
	log.Println("Email and password match")
	return 0, true, nil // Assuming 0 is the success code
}

func getUsernameFromEmail(email string) string {
	log.Println("Getting the username from the email")
	// Get the username from the email
	var username string
	err := db.QueryRow("SELECT username FROM public.customers WHERE email = $1", email).Scan(&username)
	if err != nil {
		fmt.Println("Error getting the username from the email: ", err)
		return ""
	}
	return username
}

func getCustomerIDFromEmail(email string) uint {
	log.Println("Getting the customer ID from the email")
	// Get the customer ID from the email
	var customerID uint
	err := db.QueryRow("SELECT id FROM public.customers WHERE email = $1", email).Scan(&customerID)
	if err != nil {
		fmt.Println("Error getting the customer ID from the email: ", err)
		return 0
	}
	return customerID
}

func VerifyStaff(email string) (*proto.StaffCredential, int, error) {
	log.Println("Verifying the staff with the email")
	// Verify the staff with the email
	var staffCredential = &proto.StaffCredential{}

	err := db.QueryRow("SELECT id, name FROM delivery_staffs WHERE email = $1", email).Scan(&staffCredential.StaffID, &staffCredential.StaffName)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Staff does not exist")
			return staffCredential, 1, errors.New("staff does not exist")
		} else {
			fmt.Println("Error checking if the staff exists: ", err)
			return staffCredential, 2, err
		}
	}

	log.Printf("Staff ID: %v, Staff Name: %v\n", staffCredential.StaffID, staffCredential.StaffName)

	return staffCredential, 0, nil
}