package handlers

import (
	"context"
	"log"
	"User_Service/proto"
	"User_Service/internal/services"

	"google.golang.org/grpc/status"
    "google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc"
)

type server struct {
    __.UserServiceServer
}

func RegisterMicroservicesProtoServer(s *grpc.Server) {
	__.RegisterUserServiceServer(s, &server{})
}

func (s *server) AuthUser(ctx context.Context, req *__.AuthCredential) (*__.Empty, error) {
    log.Printf("AuthUser called with email:%v, password: %v\n", req.Email, req.Password)

	// Authenticate the user with the email and password
	statusCode, success, message := services.AuthenticateUser(req.Email, req.Password)
	if statusCode != 0 {
		return nil, status.Errorf(codes.InvalidArgument, message.Error())
	}
	if !success {
		return &__.Empty{}, message
	}

	// Create a JWT token and send it back to the user
	jwtToken, err := services.CreateJWTToken(req.Email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create JWT token")
	}

	//  Create and send metadata with the JWT token
	md := metadata.Pairs("jwt_set_cookie", "jwt="+jwtToken+"; Path=/; HttpOnly; Secure")
	grpc.SendHeader(ctx, md)

    return &__.Empty{}, nil
}

func (s *server) CreateUser(ctx context.Context, req *__.Credential) (*__.Empty, error) {
    log.Printf("CreateUser called with email:%v, password: %v, username: %v\n", req.Email, req.Password, req.Username)

	// Send a verification email: + add the verification token to the database
	statusCode, message := services.SendVerificationEmail(req.Email, req.Username, req.Password)
	log.Printf("Status code: %v, Message: %v\n", statusCode, message)
	if statusCode == 2 {
		return nil, status.Errorf(codes.Internal, message)
	} else if statusCode == 1 {
		return nil, status.Errorf(codes.AlreadyExists, message)
	}
	// Store the user information in the database
    return &__.Empty{}, nil
}

func (s *server) VerifyUser(ctx context.Context, req *__.VerifyCredential) (*__.Empty, error) {
	log.Printf("VerifyUser called with email:%v, token: %v\n", req.Email, req.Token)

	// Verify the user with the token agianst the database
	statusCode, err := services.VerifyUser(req.Email, req.Token)
	if statusCode == 2 {
		return nil, status.Errorf(codes.Internal, err.Error())
	} else if statusCode == 1 {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	// Add the user to the database
	statusCode, err = services.AddUserToDB(req.Email, req.Token)
	if statusCode != 0 {
		log.Println("Error adding the user to the database: ", err)
		return nil, status.Errorf(codes.Internal, "Failed to add the user to the database")
	}

	// Delete the token from the database
	statusCode, err = services.DeleteVerificationToken(req.Email, req.Token)
	if statusCode != 0 {
		log.Println("Error deleting the token from the database: ", err)
		return nil, status.Errorf(codes.Internal, "Failed to delete the token from the database")
	}
	log.Println("Token deleted from the database")

	return &__.Empty{}, nil
}

func (s *server) VerifyStaff(ctx context.Context, req *__.EmailCredential) (*__.StaffCredential, error) {
	log.Printf("VerifyStaff called with email:%v\n", req.Email)

	// Verify the staff with the email agianst the database
	data, statusCode, err := services.VerifyStaff(req.Email)
	if statusCode == 2 {
		return nil, status.Errorf(codes.Internal, err.Error())
	} else if statusCode == 1 {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	return &__.StaffCredential{StaffID: data.StaffID, StaffName: data.StaffName}, nil
}