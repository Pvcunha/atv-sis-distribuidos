package services

import (
	"nelson/util"
)

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
