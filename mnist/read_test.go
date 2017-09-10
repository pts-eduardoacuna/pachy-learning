package mnist

import (
	"bytes"
	"encoding/binary"
	"os"
	"testing"
)

func TestImagesParsing(t *testing.T) {
	filename := "images_small_testfile"
	count := 2
	rows := 2
	cols := 3

	doImagesParsing(t, filename, count, rows, cols)
}

func TestBigImagesParsing(t *testing.T) {
	filename := "images_big_testfile"
	count := 60000
	rows := 28
	cols := 28

	doImagesParsing(t, filename, count, rows, cols)
}

func doImagesParsing(t *testing.T, filename string, count, rows, cols int) {
	err := createTestImagesFile(filename, count, rows, cols)
	if err != nil {
		t.Error("couldn't create test images file", err)
	}

	file, err := os.Open(filename)
	if err != nil {
		t.Error("couldn't open test images file", err)
	}
	defer file.Close()

	parser, err := NewImageParser(file)
	if err != nil {
		t.Error("couldn't create images parser from file", err)
	}

	if parser.Count != count {
		t.Error("parsed images count doesn't match", parser.Count, count)
	}

	if parser.Rows != rows {
		t.Error("parsed images rows doesn't match", parser.Rows, rows)
	}

	if parser.Columns != cols {
		t.Error("parsed images columns doesn't match", parser.Columns, cols)
	}

	for i := 0; i < count; i++ {
		pixels, err := parser.Parse()
		if err != nil {
			t.Error("couldn't parse image", err)
		}
		if len(pixels) != rows*cols {
			t.Error("parsed image size doesn't match", len(pixels), rows*cols)
		}
	}

	_, err = parser.Parse()
	if err == nil {
		t.Error("expecting error from parser")
	}
}

func TestLabelsParsing(t *testing.T) {
	filename := "labels_small_testfile"
	count := 10

	doLabelsParsing(t, filename, count)
}

func TestBigLabelsParsing(t *testing.T) {
	filename := "labels_big_testfile"
	count := 60000

	doLabelsParsing(t, filename, count)
}

func doLabelsParsing(t *testing.T, filename string, count int) {
	err := createTestLabelsFile(filename, count)
	if err != nil {
		t.Error("couldn't create test labels file", err)
	}

	file, err := os.Open(filename)
	if err != nil {
		t.Error("couldn't open test labels file", err)
	}
	defer file.Close()

	parser, err := NewLabelParser(file)
	if err != nil {
		t.Error("couldn't create labels parser from file", err)
	}

	if parser.Count != count {
		t.Error("parsed labels count doesn't match", parser.Count, count)
	}

	for i := 0; i < count; i++ {
		_, err := parser.Parse()
		if err != nil {
			t.Error("couldn't parse label", err)
		}
	}

	_, err = parser.Parse()
	if err == nil {
		t.Error("expecting error from parser")
	}
}

func createTestImagesFile(name string, count, rows, cols int) error {
	size := 4 + (count * rows * cols)
	data := make([]interface{}, size)
	data[0] = int32(666)
	data[1] = int32(count)
	data[2] = int32(rows)
	data[3] = int32(cols)
	for i := 4; i < size; i++ {
		data[i] = byte((i - 4) % 256)
	}

	return createFile(name, data)
}

func createTestLabelsFile(name string, count int) error {
	size := 2 + count
	data := make([]interface{}, size)
	data[0] = int32(666)
	data[1] = int32(count)
	for i := 2; i < size; i++ {
		data[i] = byte((i - 2) % 256)
	}

	return createFile(name, data)
}

func createFile(name string, data []interface{}) error {
	buff := new(bytes.Buffer)
	for _, v := range data {
		err := binary.Write(buff, binary.BigEndian, v)
		if err != nil {
			return err
		}
	}

	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(buff.Bytes())
	if err != nil {
		return err
	}

	return nil
}
