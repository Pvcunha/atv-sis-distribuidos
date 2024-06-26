package util

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"

	pb "nelson/grpc/imageserial"
)

func OpenImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fi, _ := f.Stat()
	fmt.Println(fi.Name())
	//defer f.Close()
	img, err := jpeg.Decode(f)
	if err != nil {
		fmt.Println("Decoding error:", err.Error())
		return nil, err
	}

	return img, nil
}

func SaveImage(path string, img image.Image) error {
	f, err := os.Create(path)
	if err != nil {
		return errors.New("failed saving image")
	}
	defer f.Close()

	err = jpeg.Encode(f, img, nil)
	if err != nil {
		return errors.New("failed saving image")
	}

	return nil
}
func Image2Tensor(img image.Image) [][]color.Color {
	size := img.Bounds().Size()
	var pixels [][]color.Color
	// put pixels into two three two dimensional array
	for i := 0; i < size.X; i++ {
		var y []color.Color
		for j := 0; j < size.Y; j++ {
			y = append(y, img.At(i, j))
		}
		pixels = append(pixels, y)
	}
	return pixels
}

func Tensor2Image(pixels [][]color.Color) image.Image {
	rect := image.Rect(0, 0, len(pixels), len(pixels[0]))
	nImg := image.NewRGBA(rect)

	for x := 0; x < len(pixels); x++ {
		for y := 0; y < len(pixels[0]); y++ {
			q := pixels[x]
			if q == nil {
				continue
			}
			p := pixels[x][y]
			if p == nil {
				continue
			}
			original, ok := color.RGBAModel.Convert(p).(color.RGBA)
			if ok {
				nImg.Set(x, y, original)
			}
		}
	}

	return nImg
}

func Tensor2RawPixel(pixels [][]color.Color) [][]RawPixel {

	var result [][]RawPixel
	// put pixels into two three two dimensional array
	for i := 0; i < len(pixels); i++ {
		var y []RawPixel
		for j := 0; j < len(pixels[0]); j++ {
			r, g, b, a := pixels[i][j].RGBA()
			y = append(y, RawPixel{R: r, G: g, B: b, A: a})
		}
		result = append(result, y)
	}
	return result
}

func RawPixel2Tensor(pixels [][]RawPixel) [][]color.Color {
	var result [][]color.Color

	for i := 0; i < len(pixels); i++ {
		var y []color.Color
		for j := 0; j < len(pixels[0]); j++ {
			var r, g, b, a uint32
			r, g, b, a = pixels[i][j].Get()
			// create a color.Color tensor
			p := color.NRGBA{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: uint8(a >> 8),
			}
			y = append(y, p)
		}
		result = append(result, y)
	}
	return result
}

func RawPixel2ImageData(pixels [][]RawPixel) *pb.ImageData {
	var result pb.ImageData

	for i := 0; i < len(pixels); i++ {
		row := make([]*pb.RawPixel, 0)
		for j := 0; j < len(pixels[0]); j++ {
			r, g, b, a := pixels[i][j].Get()
			row = append(row, &pb.RawPixel{R: r, G: g, B: b, A: a})
		}
		result.Rows = append(result.Rows, &pb.RawPixelRow{Pixels: row})
	}
	return &result
}

func ImageData2RawPixel(data *pb.ImageData) [][]RawPixel {
	var result [][]RawPixel

	for i := 0; i < len(data.Rows); i++ {
		var y []RawPixel
		for j := 0; j < len(data.Rows[0].Pixels); j++ {
			r := data.Rows[i].Pixels[j].GetR()
			g := data.Rows[i].Pixels[j].GetG()
			b := data.Rows[i].Pixels[j].GetB()
			a := data.Rows[i].Pixels[j].GetA()
			y = append(y, RawPixel{R: r, G: g, B: b, A: a})
		}
		result = append(result, y)
	}
	return result
}

func UpsideDown(pixels [][]RawPixel) [][]RawPixel {
	for i := 0; i < len(pixels); i++ {
		tr := pixels[i]
		for j := 0; j < len(tr)/2; j++ {
			k := len(tr) - j - 1
			tr[j], tr[k] = tr[k], tr[j]
		}
	}
	return pixels
}
