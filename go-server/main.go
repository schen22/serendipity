package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"serendipity/go-server/chat"

	"google.golang.org/grpc"
)

// RunServer registers gRPC service and run server
func RunServer(ctx context.Context, srv chat.ChatServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()
	chat.RegisterChatServiceServer(server, srv)

	// start gRPC server
	log.Println("starting gRPC server...")
	return server.Serve(listen)
}

func main() {
	if err := RunServer(context.Background(), chat.NewChatServiceServer(), "3000"); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
