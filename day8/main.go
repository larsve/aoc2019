package main

import (
	"fmt"
	"os"
)

func main() {
	f, e := os.Open("input")
	if e != nil {
		fmt.Printf("Failed to open input map: %v\n", e)
	}
	defer f.Close()
	img, e := loadImage(25, 6, f)
	if e != nil {
		fmt.Printf("Failed to load input map: %v\n", e)
	}
	fmt.Printf("Image checksum: %v\n", img.checksum())
	if e = img.saveAsPng("spacemap.png"); e != nil {
		fmt.Printf("Failed to generate spave map image: %v\n", e)
		return
	}
	fmt.Println("Space map saved as spacemap.png")
}
