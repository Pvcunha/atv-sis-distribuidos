package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"nelson/util"
	"net"
	"time"
)

var (
	run = flag.Int("run", 100, "number of runs")
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
	encoder := gob.NewEncoder(conn)
	decoder := gob.NewDecoder(conn)
	var response util.Imagepacket

	// sends package
	for i := 0; i < *run; i++ {
		start := time.Now()
		encoder.Encode(packet)
		// receive package
		err = decoder.Decode(&response)
		rtt := time.Since(start)

		if err != nil {
			fmt.Println("error while receiving")
			return
		}
		fmt.Println(rtt.Nanoseconds())
		// saves image locally
		// tensor := util.RawPixel2Tensor(response.Img)
		// err = util.SaveImage("./assets/saved.jpeg", util.Tensor2Image(tensor))
		// if err != nil {
		// 	fmt.Println(err)
		// }
	}

}
