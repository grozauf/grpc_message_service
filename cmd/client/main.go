package main

import (
	"fmt"
	p "github.com/grozauf/message_service/pkg/message_service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var port = ":8080"

func AboutToSayIt(ctx context.Context, m p.MessageServiceClient,
	text string) (*p.ResponseMS, error) {
	request := &p.RequestMS{
		Text: text,
		Subtext: "New Message!",
	}
	r, err := m.SayIt(ctx, request)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func main() {
	fmt.Println("version 0.15.0 after tag creation")
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Dial:", err)
		return
	}
	client := p.NewMessageServiceClient(conn)
	r, err := AboutToSayIt(context.Background(), client, "My Message!")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Response Text:", r.Text)
	fmt.Println("Response SubText:", r.Subtext)
}
