package main

import (
	"encoding/json"
	"fmt"
	"nelson/util"
	"net"
	"time"
)

func main() {

	imagePath := "./assets/Lenna.jpeg"

	serverEndpoint := "localhost:3030"

	// creates connection
	addr, err := net.ResolveTCPAddr("tcp", serverEndpoint)
	if err != nil {
		panic(err)
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		panic(err)
	}

	defer func(conn net.Conn) {
		fmt.Println("Closing Connection")
		conn.Close()
	}(conn)

	// prepare image
	img, err := util.OpenImage(imagePath)
	if err != nil {
		panic(err)
	}
	rawImg := util.Tensor2RawPixel(util.Image2Tensor(img))
	packet := new(util.Imagepacket)
	packet.Name = "Lenna"
	packet.Img = rawImg

	// encoder and decoder to send and receive through connection
	encoder := json.NewEncoder(conn)
	decoder := json.NewDecoder(conn)

	// sends package
	for i := 0; i < 10000; i++ {
		start := time.Now()
		encoder.Encode(packet)
		// receive package
		var response util.Imagepacket
		err = decoder.Decode(&response)

		if err != nil {
			fmt.Println("error while receiving")
			return
		}
		rtt := time.Since(start)
		fmt.Println(rtt.Nanoseconds())
		// saves image locally
		// tensor := util.RawPixel2Tensor(response.Img)
		// err = util.SaveImage("./assets/saved.jpeg", util.Tensor2Image(tensor))
		// if err != nil {
		// 	fmt.Println(err)
		// }
	}

}
