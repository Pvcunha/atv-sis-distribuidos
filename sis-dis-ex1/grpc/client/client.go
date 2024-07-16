/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
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
)

func main() {
	flag.Parse()
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

	imgTensor := util.Image2Tensor(img)
	rawImage := util.Tensor2RawPixel(imgTensor)
	protoImage := util.RawPixel2ImageData(rawImage)
	// Contact the server and print out its response.
	ctx := context.Background()

	for i := 0; i < *run; i++ {
		start := time.Now()
		_, err := client.UpsideDownImage(ctx, &pb.ImageRequest{Name: "lena.jpg", Image: protoImage})
		if err != nil {
			log.Fatalf("Failed to receive: %v", err)
		}
		rtt := time.Since(start).Nanoseconds()
		fmt.Println(rtt)
	}

	// rcvImg := util.ImageData2RawPixel(r.GetImage())
	// rcvTensor := util.RawPixel2Tensor(rcvImg)
	// rcvImage := util.Tensor2Image(rcvTensor)
	// util.SaveImage(fmt.Sprintf(util.OutputPath, r.GetName()), rcvImage)
	// log.Printf("Received Image: %s", r.GetName())
}
