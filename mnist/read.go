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

// LabelParser is a struct
type LabelParser struct {
	Count int
	file  *os.File
	buff  []byte
}

// NewLabelParser is a function
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

// Parse returns
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

// ImageParser is a struct
type ImageParser struct {
	Count   int
	Rows    int
	Columns int
	file    *os.File
	buff    []byte
}

// NewImageParser is a function
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

// Parse returns
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
