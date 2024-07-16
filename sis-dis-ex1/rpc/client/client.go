package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"image"
	"nelson/util"
	"net/rpc"
)

var (
	run = flag.Int("run", 100, "number of runs")
)

func main() {
	gob.Register(image.YCbCr{})
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
	// rawImg := util.Tensor2RawPixel(util.Image2Tensor(img))
	// packet := new(util.Imagepacket)
	// packet.Name = "Lenna"
	// packet.Img = rawImg

	imgBytes, err := util.Image2Bytes(img)
	if err != nil {
		fmt.Println("error while converting image to bytes")
		panic(err)
	}

	packet := util.Packet{
		Name: "Lenna",
		Data: imgBytes,
	}

	// call server
	var response util.Packet

	fmt.Println("Sending image")
	err = client.Call("ImageServiceRpc.GrayScale", packet, &response)
	if err != nil {
		fmt.Println(err)
		panic("error while receiving")
	}

	// for i := 0; i < *run; i++ {
	// 	start := time.Now()
	// 	err = client.Call("ImageServiceRpc.UpsideDown", packet, &response)
	// 	if err != nil {
	// 		fmt.Println("error while receiving")
	// 		return
	// 	}
	// 	rtt := time.Since(start)
	// 	fmt.Println(rtt.Nanoseconds())
	// }
	// saves image locally
	// tensor := util.RawPixel2Tensor(response.Img)

	img, err = util.Bytes2Image(response.Data)
	if err != nil {
		fmt.Println("error while converting from bytes to img", err)
	}
	err = util.SaveImage(fmt.Sprintf("./assets/%s-saved.jpeg", response.Name), img)
	if err != nil {
		fmt.Println(err)
	}
}
