package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
	"math/rand"
	"time"
)

// Updated to generate a gradient, with a randomly generated R.

type Image struct {
	seed uint8
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 255, 255)
}
func (i Image) At(x, y int) color.Color {
	return color.RGBA{i.seed, uint8(x), uint8(y), 255}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	m := Image{uint8(rand.Intn(255))}
	pic.ShowImage(m)
}
