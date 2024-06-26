package services

import (
	"context"
	"log"
	pb "nelson/grpc/imageserial"
	"nelson/util"
)

type Reques struct {
}

type ImageServiceRpc struct{}
type ImageService struct{}
type ImageServiceGrpc struct {
	pb.UnimplementedImageServer
}

func (t *ImageService) UpsideDown(pixels [][]util.RawPixel) {
	for i := 0; i < len(pixels); i++ {
		tr := pixels[i]
		for j := 0; j < len(tr)/2; j++ {
			k := len(tr) - j - 1
			tr[j], tr[k] = tr[k], tr[j]
		}
	}
}

func (t *ImageServiceRpc) UpsideDown(req util.Imagepacket, resp *util.Imagepacket) error {
	for i := 0; i < len(req.Img); i++ {
		tr := req.Img[i]
		for j := 0; j < len(tr)/2; j++ {
			k := len(tr) - j - 1
			tr[j], tr[k] = tr[k], tr[j]
		}
	}

	resp.Img = req.Img
	resp.Name = req.Name
	return nil
}

func (t *ImageServiceGrpc) UpsideDownImage(ctx context.Context, in *pb.ImageRequest) (*pb.ImageResponse, error) {
	log.Printf("Received image %s", in.GetName())
	img := in.GetImage()
	rawPixel := util.ImageData2RawPixel(img)
	upsidedown := util.UpsideDown(rawPixel)

	img = util.RawPixel2ImageData(upsidedown)
	return &pb.ImageResponse{Name: in.GetName(), Image: img}, nil
}
