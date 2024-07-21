package services

import (
	"context"
	"fmt"
	"image"
	"image/color"
	"log"
	pb "nelson/grpc/imageserial"
	"nelson/util"
)

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

func (t *ImageServiceRpc) Echo(req util.Packet, resp *util.Packet) error {
	fmt.Println("echo image")
	resp.Data = req.Data
	resp.Name = req.Name
	return nil
}

func (t *ImageServiceRpc) GrayScale(req util.Packet, resp *util.Packet) error {
	fmt.Println("gray scale image")
	img, err := util.Bytes2Image(req.Data)
	if err != nil {
		return err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	imgSet := image.NewRGBA(bounds)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			oldColor := img.At(x, y)
			r, g, b, _ := oldColor.RGBA()

			var gray float32 = (float32(r) * 0.3) + (float32(g) * 0.59) + (float32(b) * 0.11)
			pixel := color.Gray{uint8(gray / 256)}
			imgSet.Set(x, y, pixel)
		}
	}

	bytes, err := util.Image2Bytes(imgSet)
	if err != nil {
		return err
	}
	resp.Data = bytes
	resp.Name = req.Name + "_gray"
	return nil
}

func (t *ImageServiceGrpc) GrayScaleImage(ctx context.Context, in *pb.ImageRequestGray) (*pb.ImageResponseGray, error) {
	log.Printf("Received image %s", in.GetName())
	img := in.GetImage()
	bytes := util.ImageDataGray2Bytes(img)
	grayscale, _ := util.GrayScale(bytes)

	img = util.Bytes2ImageDataGray(grayscale)

	return &pb.ImageResponseGray{Name: in.GetName(), Image: img}, nil
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
