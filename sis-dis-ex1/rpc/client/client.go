package main

import (
	"flag"
	"fmt"
	"nelson/util"
	"net/rpc"
	"time"
)

var (
	run = flag.Int("run", 100, "number of runs")
)

func main() {

	flag.Parse()
	imagePath := "./assets/Lenna.jpeg"

	serverEndpoint := "localhost:3030"

	// creates connection  rpc

	client, err := rpc.Dial("tcp", serverEndpoint)
	if err != nil {
		panic(err)
	}

	defer func(client *rpc.Client) {
		fmt.Println("Closing Connection")
		client.Close()
	}(client)

	img, err := util.OpenImage(imagePath)
	if err != nil {
		panic(err)
	}
	rawImg := util.Tensor2RawPixel(util.Image2Tensor(img))
	packet := new(util.Imagepacket)
	packet.Name = "Lenna"
	packet.Img = rawImg

	// call server
	var response util.Imagepacket
	for i := 0; i < *run; i++ {
		start := time.Now()
		err = client.Call("ImageServiceRpc.UpsideDown", packet, &response)
		if err != nil {
			fmt.Println("error while receiving")
			return
		}
		rtt := time.Since(start)
		fmt.Println(rtt.Nanoseconds())
	}
	// saves image locally
	// tensor := util.RawPixel2Tensor(response.Img)
	// err = util.SaveImage("./assets/saved.jpeg", util.Tensor2Image(tensor))
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
