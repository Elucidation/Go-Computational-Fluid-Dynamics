package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

var (
	white color.Color = color.RGBA{255, 255, 255, 255}
	black color.Color = color.RGBA{0, 0, 0, 255}
	blue  color.Color = color.RGBA{0, 0, 255, 255}
)

func initPNG(rows int, cols int) *image.RGBA {

	m := image.NewRGBA(image.Rect(0, 0, rows, cols)) //*NRGBA (image.Image interface)

	// fill m in blue
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	return m
}

func updatePNG(m *image.RGBA, row int, arr []float64) {
	a_max := max(arr[:])
	for i := range arr {
		val := uint8(arr[i] / a_max * 255)
		m.Set(row, i, color.RGBA{val, val, val, 255})
	}
}

func writePNG(m *image.RGBA, filename string) {
	w, _ := os.Create(filename)
	defer w.Close()
	png.Encode(w, m) //Encode writes the Image m to w in PNG format.
}
