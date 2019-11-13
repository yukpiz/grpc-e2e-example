package main

import (
	"fmt"
	"net"
	"os"

	"github.com/yukpiz/grpc-e2e-example/pb"
	"github.com/yukpiz/grpc-e2e-example/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	server := grpc.NewServer()

	svc := service.NewService()
	pb.RegisterExampleServiceServer(server, svc)

	reflection.Register(server)

	conn, err := net.Listen("tcp", ":1111")
	if err != nil {
		fmt.Printf("network I/O error: %v", err)
		os.Exit(1)
	}

	fmt.Println("...Waiting for localhost:1111")
	if err := server.Serve(conn); err != nil {
		fmt.Printf("serve error: %v", err)
		os.Exit(1)
	}
}
