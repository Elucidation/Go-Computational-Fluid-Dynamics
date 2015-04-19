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

func updatePNG(m *image.RGBA, arr [][]float64, maxInitVal float64) {
	// a_max := max(arr[:])
	a_max := maxInitVal
	for i := range arr {
		for j := range arr[i] {
			if arr[i][j] > a_max {
				// Red is too big
				m.Set(j, i, color.RGBA{255, 0, 0, 255})
			} else if arr[i][j] <= 0 {
				// Green is zero
				m.Set(j, i, color.RGBA{0, 255, 0, 255})
			} else {
				val := uint8((arr[i][j] / a_max) * 255)
				m.Set(j, i, color.RGBA{val, val, val, 255})
			}
			// diffuse_val := uint8((arr[i] / maxInitVal) * 255)
		}
	}
}

func writePNG(m *image.RGBA, filename string) {
	w, _ := os.Create(filename)
	defer w.Close()
	png.Encode(w, m) //Encode writes the Image m to w in PNG format.
}
