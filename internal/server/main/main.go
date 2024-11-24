package main

import "grant.goose/csi/internal/server"

func main() {
	server.Start("unix:///Users/grant/some.sock")
}
