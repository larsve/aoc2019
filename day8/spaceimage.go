package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"math"
	"os"
)

type spaceImage struct {
	layers []*spaceImageLayer
	width  int
	height int
}

func (img *spaceImage) checksum() int {
	minZeroDigits := math.MaxInt32
	var minLayer *spaceImageLayer
	for _, l := range img.layers {
		layerZeroCount := l.distibution[0]
		if layerZeroCount < minZeroDigits {
			minZeroDigits = layerZeroCount
			minLayer = l
		}
	}
	return minLayer.distibution[1] * minLayer.distibution[2]
}

func (img *spaceImage) saveAsPng(filename string) error {
	f, e := os.Create(filename)
	if e != nil {
		return e
	}

	topLeft := image.Point{0, 0}
	bottomRight := image.Point{img.width, img.height}
	bmp := image.NewRGBA(image.Rectangle{topLeft, bottomRight})

	// Draw layers in reverse order, starting with the last layer
	for i := len(img.layers) - 1; i >= 0; i-- {
		l := img.layers[i]
		for address, pixel := range l.rawData {
			x := address % img.width
			y := address / img.width
			switch pixel {
			case 0:
				bmp.Set(x, y, color.Black)
			case 1:
				bmp.Set(x, y, color.White)
			}
		}
	}
	return png.Encode(f, bmp)
}

func loadImage(imgWidth, imgHeight int, imageSource io.Reader) (*spaceImage, error) {
	img := &spaceImage{layers: []*spaceImageLayer{}, width: imgWidth, height: imgHeight}
	layerSize := imgWidth * imgHeight
	readbuffer := make([]byte, layerSize)
	for {
		read, err := imageSource.Read(readbuffer)
		if read != layerSize {
			if len(img.layers) == 0 {
				return nil, fmt.Errorf("Not enough data, expected %v, but got %v reading layer %v", layerSize, read, len(img.layers)+1)
			}
			fmt.Printf("Warning! Layer %v skipped since it only contained %v bytes out of the %v expected\n", len(img.layers)+1, read, layerSize)
			break
		}
		img.layers = append(img.layers, newLayer(readbuffer))
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
	}
	fmt.Printf("%v layers loaded..", len(img.layers))

	return img, nil
}
