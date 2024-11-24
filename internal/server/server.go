package server

import (
	"errors"
	"net"
	"strings"

	"google.golang.org/grpc"
)

const unixPrefix string = "unix://"

var TestServiceDesc = grpc.ServiceDesc{
	ServiceName: "TestService",
	HandlerType: (any)(nil), // we don't really care about the handler type being checked for now
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    nil, // TODO we need to create a handler
		},
	},
	Streams:  []grpc.StreamDesc{}, // what is this used for?
	Metadata: "nothing",           // what is this used for?
}

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
	server := grpc.NewServer()

	server.RegisterService(&TestServiceDesc, nil) // TODO need to register an implementation?

	server.Serve(listener)

	return nil
}
