package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"

	"nelson/services"
	"nelson/util"
)

func handleClient(conn net.Conn) {
	jsonEncoder := json.NewEncoder(conn)
	jsonDecoder := json.NewDecoder(conn)

	var data util.Imagepacket

	imgService := new(services.ImageService)

	defer func(conn net.Conn) {
		fmt.Println("Closing Connection")
		conn.Close()
	}(conn)

	for {

		err := jsonDecoder.Decode(&data)
		switch err {
		case nil:
			break
		case io.EOF:
			fmt.Println("Connection closed client side")
			return
		default:
			fmt.Println("error decoding data")
			fmt.Println(err)
			return
		}

		imgService.UpsideDown(data.Img)
		jsonEncoder.Encode(data)
	}

}

func runServer() {
	endpoint := "localhost:3030"
	addr, err := net.ResolveTCPAddr("tcp", endpoint)
	if err != nil {
		panic(err)
	}

	ln, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	fmt.Printf("Servidor online no endereço %v\n", addr)
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		go handleClient(conn)
	}
}

func main() {
	runServer()
}
