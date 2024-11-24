package main

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grant.goose/csi/internal/server/testapi"
)

func main() {

	con, err := grpc.Dial("unix:///Users/grant/some.sock", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		os.Exit(1)
	}

	defer con.Close()
	req := testapi.Request{Ping: "ping"}

	client := testapi.NewTestServiceClient(con)

	res, reserr := client.Ping(context.Background(), &req)

	if reserr != nil {
		fmt.Printf("Error: %s", reserr.Error())
		os.Exit(1)
	}

	fmt.Printf("%s", res.Pong)
}
