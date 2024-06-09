package main

import (
	"fmt"
	"nelson/services"
	"net"
	"net/rpc"
)

func runServer() {
	endpoint := "localhost:3030"

	imgService := new(services.ImageServiceRpc)

	server := rpc.NewServer()
	err := server.Register(imgService)

	if err != nil {
		panic(err)
	}

	ln, err := net.Listen("tcp", endpoint)
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	fmt.Println("Manda pro pai")
	server.Accept(ln)
}

func main() {
	runServer()
}
