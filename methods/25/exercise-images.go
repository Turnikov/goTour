package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	w, h  int
	color int
}

func (p Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (p Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, p.w, p.h)
}

func (p Image) At(x, y int) color.Color {
	v := uint8((x + y) / 2)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{255, 255, 255}
	pic.ShowImage(m)
}
