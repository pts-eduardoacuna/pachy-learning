// Package image allows reading and writing PNG images via int slices.
package image

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

// WritePNG writes a PNG image from the given pixel data.
func WritePNG(file *os.File, rows, cols int, pixels []int) error {
	img := image.NewGray(image.Rect(0, 0, rows, cols))

	for i, value := range pixels {
		if value == 0 {
			img.Set(i%cols, i/rows, color.Gray{uint8(255)})
		} else {
			img.Set(i%cols, i/rows, color.Gray{uint8(0)})
		}
	}

	return png.Encode(file, img)
}

// ReadPNG reads a PNG image and computes it's pixel data.
func ReadPNG(file *os.File) ([]int, error) {
	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}

	rect := img.Bounds()
	cols := rect.Dx()
	rows := rect.Dy()
	pixels := make([]int, rows*cols)

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			color := img.At(col, row)
			r, g, b, _ := color.RGBA()
			if r == 0 && g == 0 && b == 0 {
				pixels[cols*row+col] = 1
			} else {
				pixels[cols*row+col] = 0
			}
		}
	}

	return pixels, nil
}
