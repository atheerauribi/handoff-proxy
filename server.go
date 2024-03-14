package main

import (
	"context"
	"fmt"
	"log"
	"net"

	// Import the generated gRPC code for the service
	pb "github.com/atheerauribi/handoff-proxy/proto"
	"google.golang.org/grpc"
)

// Assume your service is named proxyServer and has a method Add
type proxyServer struct {
	pb.UnimplementedCalculatorServer
	calcClient pb.CalculatorClient
}

// NewProxyServer creates a new proxy server instance
func NewProxyServer(calcAddr string) (*proxyServer, error) {
	// Connect to server2
	conn, err := grpc.Dial(calcAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	calcClient := pb.NewCalculatorClient(conn)

	return &proxyServer{
		calcClient: calcClient,
	}, nil
}

// Add is a proxy method that forwards the request to server2 and returns the response
func (p *proxyServer) Add(ctx context.Context, req *pb.AddRequest) (*pb.OperationResponse, error) {
	// Forward the request to Steve and get the response
	resp, err := p.calcClient.Add(ctx, req)
	if err != nil {
		return nil, err
	}

	// Return the response from Steve back to the client (Bob)
	return resp, nil
}

// Divide is a proxy method that forwards the request to server2 and returns the response
func (p *proxyServer) Divide(ctx context.Context, req *pb.DivideRequest) (*pb.OperationResponse, error) {
	// Forward the request to Steve and get the response
	resp, err := p.calcClient.Divide(ctx, req)
	if err != nil {
		return nil, err
	}

	// Return the response from Steve back to the client (Bob)
	return resp, nil
}

func RunServer() error {
	// Address of the server2 server
	server2addr := "localhost:8889"

	// Create the proxy server instance
	proxy, err := NewProxyServer(server2addr)
	if err != nil {
		log.Fatalf("failed to create proxy: %v", err)
	}

	// Listen for incoming connections from server1
	fmt.Println("Running Server\nListening on port 50052...")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50052))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	// Create a new gRPC server
	s := grpc.NewServer(opts...)

	// Register the proxy server as a SteveService server
	pb.RegisterCalculatorServer(s, proxy)

	// Start serving requests
	return s.Serve(lis)
}

// RunServer runs the gRPC server
// func RunServer() error {
// 	fmt.Println("Running Server\nListening on port 8889...")
// 	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8889))
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}
// 	var opts []grpc.ServerOption

// 	s := grpc.NewServer(opts...)
// 	pb.RegisterCalculatorServer(s, &calculatorServer{})
// 	return s.Serve(lis)
// }
