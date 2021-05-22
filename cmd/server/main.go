package main

import (
	"fmt"
	p "github.com/grozauf/grpc_message_service/pkg/message_service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
)

type MessageServer struct {
	p.UnimplementedMessageServiceServer
}

var port = ":8080"

func (MessageServer) SayIt(ctx context.Context, r *p.RequestMS) (*p.ResponseMS, error) {
	fmt.Println("Request Text:", r.Text)
	fmt.Println("Request SubText:", r.Subtext)
	response := &p.ResponseMS{
		Text: r.Text,
		Subtext: "Got it!",
	}
	return response, nil
}

func main() {
	fmt.Println("version 0.2.0 in new repository")
	server := grpc.NewServer()
	var messageServer MessageServer
	p.RegisterMessageServiceServer(server, messageServer)
	listen, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Serving requests...")
	server.Serve(listen)
}
