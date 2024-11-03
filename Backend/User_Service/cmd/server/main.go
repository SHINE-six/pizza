package main

import (
    "log"
    "net"
    "User_Service/config"
    "User_Service/internal/handlers"

    "google.golang.org/grpc"
)


func setUpServer() *grpc.Server {
    return grpc.NewServer()
}

func main() {
    // Load the configuration
    cfg := config.LoadServerConfig()

    // Listen for incoming connections
    lis, err := net.Listen(cfg.NetworkType, cfg.ServerPort)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    
    // Set up the server
    s := setUpServer()

    // Register the microservices proto server
    handlers.RegisterMicroservicesProtoServer(s)

    // Print the server address
    log.Printf("server listening at %v", lis.Addr())

    // Serve the server on the listener
    err = s.Serve(lis)
    if err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
