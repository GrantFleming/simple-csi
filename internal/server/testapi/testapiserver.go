package testapi

import (
	context "context"
	"fmt"
)

type TestApiServer struct {
	UnimplementedTestServiceServer
}

func (s TestApiServer) Ping(ctx context.Context, in *Request) (*Response, error) {
	if in.Ping == "ping" {
		fmt.Printf("Received: %s", in.Ping)
		return &Response{Pong: "pong"}, nil
	}
	return &Response{Pong: "bork"}, nil
}
