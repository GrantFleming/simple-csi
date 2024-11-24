package server

import (
	"errors"
	"fmt"
	"net"
	"strings"

	"google.golang.org/grpc"
	"grant.goose/csi/internal/server/testapi"
)

const unixPrefix string = "unix://"

func Listen(endpoint string) (net.Listener, error) {
	// endpoint is expected to be a ???
	if !strings.HasPrefix(endpoint, unixPrefix) {
		return nil, errors.New("Cannot listed on something that is not a unix socket")
	}

	addr := strings.TrimPrefix(endpoint, unixPrefix)

	return net.Listen("unix", addr)
}

func Start(endpoint string) error {
	// the listener allows you to await connections, then read/write bytes
	// it is very low-level
	// you could use this to read/write raw data over a socket!
	listener, err := Listen(endpoint)

	if err != nil {
		return err
	}

	// this works via the registration of handlers with the grpc server
	fmt.Print("Starting the grpc server")
	server := grpc.NewServer()

	testapi.RegisterTestServiceServer(server, testapi.TestApiServer{})

	server.Serve(listener)

	return nil
}
