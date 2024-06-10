package services

import (
	"nelson/util"
)

type Reques struct {
}

type ImageServiceRpc struct{}
type ImageService struct{}

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
