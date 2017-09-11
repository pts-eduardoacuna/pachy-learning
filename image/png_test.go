package image

import (
	"os"
	"testing"
)

func TestPNG(t *testing.T) {
	filename := "glider_testfile.png"
	pixels, rows, cols := createDummyImageData()
	doWritePNG(t, filename, rows, cols, pixels)
	doReadPNG(t, filename, pixels)
}

func createDummyImageData() ([]int, int, int) {
	glider := []int{
		0, 1, 0,
		0, 0, 1,
		1, 1, 1,
	}

	return glider, 3, 3
}

func doWritePNG(t *testing.T, filename string, rows, cols int, pixels []int) {
	file, err := os.Create(filename)
	if err != nil {
		t.Error("couldn't create testfile", err)
	}
	defer file.Close()

	err = WritePNG(file, rows, cols, pixels)
	if err != nil {
		t.Error("couldn't create image testfile", err)
	}
}

func doReadPNG(t *testing.T, filename string, expected []int) {
	file, err := os.Open(filename)
	if err != nil {
		t.Error("couldn't open testfile", err)
	}
	defer file.Close()

	pixels, err := ReadPNG(file)
	if err != nil {
		t.Error("couldn't open image testfile", err)
	}

	if len(pixels) != len(expected) {
		t.Error("expecting images to be the same size", len(expected), len(pixels))
	}
	for i, value := range pixels {
		if value != expected[i] {
			t.Error("unexpected value in image testfile", value, expected[i])
		}
	}
}
