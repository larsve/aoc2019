package main

import (
	"os"
	"testing"

	"aoc2019/common"
)

func TestImageLoad(t *testing.T) {
	f, e := os.Open("input")
	common.Assert(t, e == nil, "Failed to open input file: %v", e)
	defer f.Close()
	img, e := loadImage(25, 6, f)
	common.Assert(t, e == nil, "Failed to load image: %v", e)
	common.Assert(t, img != nil, "Failed to load image, no image returned")
}

func TestImageChecksum(t *testing.T) {
	f, e := os.Open("input")
	common.Assert(t, e == nil, "Failed to open input file: %v", e)
	defer f.Close()
	img, e := loadImage(25, 6, f)
	common.Assert(t, e == nil, "Failed to load image: %v", e)
	common.Assert(t, img != nil, "Failed to load image, no image returned")
	cs := img.checksum()
	common.Assert(t, cs == 2318, "Invalid space image checksum, got %v, expected 2318", cs)
}

func TestImageSameAsPng(t *testing.T) {
	f, e := os.Open("input")
	common.Assert(t, e == nil, "Failed to open input file: %v", e)
	defer f.Close()
	img, e := loadImage(25, 6, f)
	common.Assert(t, e == nil, "Failed to load image: %v", e)
	common.Assert(t, img != nil, "Failed to load image, no image returned")
	e = img.saveAsPng("test.png")
	common.Assert(t, e == nil, "Failed to generate PNG image: %v", e)
}
