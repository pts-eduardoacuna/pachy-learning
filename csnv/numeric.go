package csnv

import (
	"encoding/csv"
	"strconv"
)

// ReadFloats is a wrapper around the standard CSV reader for records containing float64.
func ReadFloats(r *csv.Reader) ([]float64, error) {
	strs, err := r.Read()
	if err != nil {
		return nil, err
	}
	floats := make([]float64, len(strs))
	for i, str := range strs {
		value, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, err
		}
		floats[i] = value
	}
	return floats, nil
}

// WriteFloats is a wrapper around the standard CSV writer for record containing float64.
func WriteFloats(w *csv.Writer, floats []float64) error {
	strs := make([]string, len(floats))
	for i, value := range floats {
		strs[i] = strconv.FormatFloat(value, 'f', -1, 64)
	}

	return w.Write(strs)
}

// ReadInts is a wrapper around the standard CSV reader for records containing int.
func ReadInts(r *csv.Reader) ([]int, error) {
	strs, err := r.Read()
	if err != nil {
		return nil, err
	}
	ints := make([]int, len(strs))
	for i, str := range strs {
		value, err := strconv.ParseInt(str, 10, 0)
		if err != nil {
			return nil, err
		}
		ints[i] = int(value)
	}
	return ints, nil
}

// WriteInts is a wrapper around the standard CSV writer for record containing int.
func WriteInts(w *csv.Writer, ints []int) error {
	strs := make([]string, len(ints))
	for i, value := range ints {
		strs[i] = strconv.FormatInt(int64(value), 10)
	}

	return w.Write(strs)
}
