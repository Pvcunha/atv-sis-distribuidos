package main

import (
	"fmt"
	"nelson/util"
	"net/rpc"
)

func main() {

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
	err = client.Call("ImgEditor.UpsideDown", packet, &response)

	if err != nil {
		fmt.Println("error while receiving")
		return
	}

	// saves image locally
	tensor := util.RawPixel2Tensor(response.Img)
	err = util.SaveImage("./assets/saved.jpeg", util.Tensor2Image(tensor))
	if err != nil {
		fmt.Println(err)
	}
}
