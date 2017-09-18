package image

import (
	"os"

	"gonum.org/v1/gonum/mat"
)

// ReadDataset reads a gonum matrix from a PNG file.
func ReadDataset(file *os.File) (*mat.Dense, error) {
	pixels, err := ReadPNG(file)
	if err != nil {
		return nil, err
	}

	floats := make([]float64, len(pixels))
	for i, v := range pixels {
		floats[i] = float64(v)
	}
	matrix := mat.NewDense(1, len(pixels), floats)

	return matrix, nil
}
