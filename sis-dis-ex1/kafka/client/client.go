package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "nelson/grpc/imageserial"
	"nelson/util"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	run  = flag.Int("run", 100, "number of runs")
	conc = flag.Bool("conc", false, "concurrent mode")
)

func init() {
	flag.Parse()
}

func main() {
	// Set up a connection to the server

	conn, err := grpc.NewClient(
		*addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(7194304)),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(7194304)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewImageClient(conn)

	//loadsImage
	img, err := util.OpenImage(util.ImagePath)
	if err != nil {
		panic(err)
	}
	/*
		imgTensor := util.Image2Tensor(img)
		rawImage := util.Tensor2RawPixel(imgTensor)
		protoImage := util.RawPixel2ImageData(rawImage)
	*/
	imgBytes, err := util.Image2Bytes(img)
	if err != nil {
		panic(err)
	}
	reqData := &pb.ImageRequestGray{Name: "lena.jpg", Data: imgBytes, Conc: *conc}
	var response *pb.ImageResponseGray

	for i := 0; i < *run; i++ {
		start := time.Now()
		response, err = client.GrayScaleImage(context.Background(), reqData)

		if err != nil {
			log.Fatalf("Failed to receive: %v", err)
		}
		rtt := time.Since(start).Nanoseconds()
		fmt.Println(rtt)
	}

	img, err = util.Bytes2Image(response.GetData())
	if err != nil {
		panic(err)
	}

	err = util.SaveImage(fmt.Sprintf(util.OutputPath, response.Name), img)
	if err != nil {
		panic(err)
	}

	// Contact the server and print out its response.
	// ctx := context.Background()

	/*
		for i := 0; i < *run; i++ {
			start := time.Now()
			_, err := client.UpsideDownImage(ctx, &pb.ImageRequest{Name: "lena.jpg", Image: protoImage})
			if err != nil {
				log.Fatalf("Failed to receive: %v", err)
			}
			rtt := time.Since(start).Nanoseconds()
			fmt.Println(rtt)
		}
	*/
	// rcvImg := util.ImageData2RawPixel(r.GetImage())
	// rcvTensor := util.RawPixel2Tensor(rcvImg)
	// rcvImage := util.Tensor2Image(rcvTensor)
	// util.SaveImage(fmt.Sprintf(util.OutputPath, r.GetName()), rcvImage)
	// log.Printf("Received Image: %s", r.GetName())
}
