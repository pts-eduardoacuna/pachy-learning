// Package mnist allows working with images and labels MNIST binary files.
package mnist

import (
	"bytes"
	"encoding/binary"
	"os"
)

type labelsHeader struct {
	MagicNumber int32
	Count       int32
}

// LabelParser holds the state of a label parsing process.
type LabelParser struct {
	Count int
	file  *os.File
	buff  []byte
}

// NewLabelParser creates a label parser from the given file.
func NewLabelParser(file *os.File) (*LabelParser, error) {
	labelsHeader := labelsHeader{}

	err := parseHeader(file, 8, &labelsHeader)
	if err != nil {
		return nil, err
	}

	buff := make([]byte, 1)
	parser := &LabelParser{
		Count: int(labelsHeader.Count),
		file:  file,
		buff:  buff,
	}

	return parser, nil
}

// Parse reads a label from the parser file.
//
// The returned value is in the range [0,9].
func (p *LabelParser) Parse() (int, error) {
	_, err := p.file.Read(p.buff)
	if err != nil {
		return 0, err
	}
	return int(p.buff[0]), nil
}

type imagesHeader struct {
	MagicNumber int32
	Count       int32
	Rows        int32
	Columns     int32
}

// ImageParser holds the state of an image parsing process.
type ImageParser struct {
	Count   int
	Rows    int
	Columns int
	file    *os.File
	buff    []byte
}

// NewImageParser creates an image parser from the given file.
func NewImageParser(file *os.File) (*ImageParser, error) {
	imagesHeader := imagesHeader{}

	err := parseHeader(file, 16, &imagesHeader)
	if err != nil {
		return nil, err
	}

	count := int(imagesHeader.Count)
	rows := int(imagesHeader.Rows)
	columns := int(imagesHeader.Columns)
	buff := make([]byte, rows*columns)
	parser := &ImageParser{
		Count:   count,
		Rows:    rows,
		Columns: columns,
		file:    file,
		buff:    buff,
	}

	return parser, nil
}

// Parse reads an image from the parser file and returns it as a slice of integers.
//
// The elements of the slice are the image's pixels grayscale values in the range [0,255].
func (p *ImageParser) Parse() ([]int, error) {
	_, err := p.file.Read(p.buff)
	if err != nil {
		return nil, err
	}

	elms := make([]int, p.Rows*p.Columns)
	for i, b := range p.buff {
		elms[i] = int(b)
	}
	return elms, nil
}

func parseHeader(file *os.File, size int, header interface{}) error {
	bs := make([]byte, size)
	_, err := file.Read(bs)
	if err != nil {
		return err
	}
	buffer := bytes.NewReader(bs)
	err = binary.Read(buffer, binary.BigEndian, header)
	if err != nil {
		return err
	}
	return nil
}
