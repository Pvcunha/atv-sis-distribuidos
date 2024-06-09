package main

import (
	"encoding/gob"
	"fmt"
	"io"
	"net"

	"nelson/services"
	"nelson/util"
)

func handleClient(conn net.Conn) {
	// encoder := json.NewEncoder(conn)
	// decoder := json.NewDecoder(conn)
	encoder := gob.NewEncoder(conn)
	decoder := gob.NewDecoder(conn)
	var data util.Imagepacket

	imgService := new(services.ImageService)

	defer func(conn net.Conn) {
		fmt.Println("Closing Connection")
		conn.Close()
	}(conn)

	for {

		err := decoder.Decode(&data)
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
		encoder.Encode(data)
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
