package services

import (
	"User_Service/proto"

	"regexp"
)


func AssureAllRequiredInformation(req *__.Credential) (status int, message string) {
	// Check each field individually and return a specific message if it's missing
	if req.Email == "" {
		return 1, "Missing email"
	}
	if req.Password == "" {
		return 1, "Missing password"
	}
	if req.Username == "" {
		return 1, "Missing username"
	}
	// If all required information is present
	status, message = validateEmail(req.Email)
	if (status != 0) {
		return status, message
	}

	status, message = validatePassword(req.Password)
	if (status != 0) {
		return status, message
	}

	status, message = validateUsername(req.Username)
	if (status != 0) {
		return status, message
	}

	return 0, "All required information is present"
}

func validateEmail(email string) (status int, message string) {
	// Check is it a proper email format
	// format of example@mail.com
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return 1, "Invalid email"
	}
	
	//TODO: Check if the email is already in use

	return 0, "Email is valid"
}

func validatePassword(password string) (status int, message string) {
	// Check if the password is strong enough
	// At least 8 characters long
	if len(password) < 8 {
		return 1, "Password is too short, must be at least 8 characters long"
	}
	// Contains at least one uppercase letter
	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return 1, "Password is missing an uppercase letter"
	}
	// Contains at least one lowercase letter
	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return 1, "Password is missing a lowercase letter"
	}
	// Contains at least one number
	if !regexp.MustCompile(`[0-9]`).MatchString(password) {
		return 1, "Password is missing a number"
	}
	// Contains at least one special character
	if !regexp.MustCompile(`[!@#$%^&*()_+{}|:<>?]`).MatchString(password) {
		return 1, "Password is missing a special character, (!@#$%^&*()_+{}|:<>?)"
	}
	return 0, "Password is valid"
}

func validateUsername(username string) (status int, message string) {
	// Check if the username is valid
	// At least 3 characters long
	if len(username) < 3 {
		return 1, "Username is too short"
	}
	// Contains only letters and numbers
	if !regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(username) {
		return 1, "Username contains invalid characters"
	}
	//TODO: Check if the username is already in use
	return 0, "Username is valid"
}