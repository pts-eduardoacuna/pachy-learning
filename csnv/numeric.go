package csnv

import (
	"encoding/csv"
	"strconv"
)

func convertToFloatsRecord(strs []string) ([]float64, error) {
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

func convertFromFloatsRecord(floats []float64) []string {
	strs := make([]string, len(floats))
	for i, value := range floats {
		strs[i] = strconv.FormatFloat(value, 'f', -1, 64)
	}

	return strs
}

func convertToIntsRecord(strs []string) ([]int, error) {
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

func convertFromIntsRecord(ints []int) []string {
	strs := make([]string, len(ints))
	for i, value := range ints {
		strs[i] = strconv.FormatInt(int64(value), 10)
	}

	return strs
}

// ReadFloats is a wrapper around the standard CSV Read for records containing float64.
func ReadFloats(r *csv.Reader) ([]float64, error) {
	strs, err := r.Read()
	if err != nil {
		return nil, err
	}
	floats, err := convertToFloatsRecord(strs)

	return floats, err
}

// ReadAllFloats is a wrapper around the standard CSV ReadAll for records containing float64.
func ReadAllFloats(r *csv.Reader) ([][]float64, error) {
	allStrs, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	allFloats := make([][]float64, len(allStrs))
	for i, strs := range allStrs {
		floats, err := convertToFloatsRecord(strs)
		if err != nil {
			return nil, err
		}
		allFloats[i] = floats
	}

	return allFloats, nil
}

// WriteFloats is a wrapper around the standard CSV Write for records containing float64.
func WriteFloats(w *csv.Writer, floats []float64) error {
	strs := convertFromFloatsRecord(floats)

	return w.Write(strs)
}

// WriteAllFloats is a wrapper around the standard CSV WriteAll for records containing float64.
func WriteAllFloats(w *csv.Writer, allFloats [][]float64) error {
	allStrs := make([][]string, len(allFloats))
	for i, floats := range allFloats {
		strs := convertFromFloatsRecord(floats)
		allStrs[i] = strs
	}

	return w.WriteAll(allStrs)
}

// ReadInts is a wrapper around the standard CSV Read for records containing int.
func ReadInts(r *csv.Reader) ([]int, error) {
	strs, err := r.Read()
	if err != nil {
		return nil, err
	}
	ints, err := convertToIntsRecord(strs)

	return ints, err
}

// ReadAllInts is a wrapper around the standard CSV ReadAll for records containing int.
func ReadAllInts(r *csv.Reader) ([][]int, error) {
	allStrs, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	allInts := make([][]int, len(allStrs))
	for i, strs := range allStrs {
		ints, err := convertToIntsRecord(strs)
		if err != nil {
			return nil, err
		}
		allInts[i] = ints
	}

	return allInts, nil
}

// WriteInts is a wrapper around the standard CSV Write for records containing int.
func WriteInts(w *csv.Writer, ints []int) error {
	strs := convertFromIntsRecord(ints)

	return w.Write(strs)
}

// WriteAllInts is a wrapper around the standard CSV WriteAll for records containing int.
func WriteAllInts(w *csv.Writer, allInts [][]int) error {
	allStrs := make([][]string, len(allInts))
	for i, ints := range allInts {
		strs := convertFromIntsRecord(ints)
		allStrs[i] = strs
	}

	return w.WriteAll(allStrs)
}
