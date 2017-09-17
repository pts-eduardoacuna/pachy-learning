package csnv

import (
	"encoding/csv"
	"math/rand"
	"os"
	"testing"
)

func TestFloats(t *testing.T) {
	filename := "floats_small_testfile"
	count := 3
	fields := 4

	data := createFloatData(count, fields)
	doFloatsWrite(t, filename, data)
	doFloatsRead(t, filename, data)
}

func TestAllFloats(t *testing.T) {
	filename := "all_floats_small_testfile"
	count := 3
	fields := 4

	data := createFloatData(count, fields)
	doAllFloatsWrite(t, filename, data)
	doAllFloatsRead(t, filename, data)
}

func TestBigFloats(t *testing.T) {
	filename := "floats_big_testfile"
	count := 6000
	fields := 78

	data := createFloatData(count, fields)
	doFloatsWrite(t, filename, data)
	doFloatsRead(t, filename, data)
}

func TestBigAllFloats(t *testing.T) {
	filename := "all_floats_big_testfile"
	count := 6000
	fields := 78

	data := createFloatData(count, fields)
	doAllFloatsWrite(t, filename, data)
	doAllFloatsRead(t, filename, data)
}

func TestInts(t *testing.T) {
	filename := "ints_small_testfile"
	count := 3
	fields := 4

	data := createIntData(count, fields)
	doIntsWrite(t, filename, data)
	doIntsRead(t, filename, data)
}

func TestAllInts(t *testing.T) {
	filename := "all_ints_small_testfile"
	count := 3
	fields := 4

	data := createIntData(count, fields)
	doAllIntsWrite(t, filename, data)
	doAllIntsRead(t, filename, data)
}

func TestBigInts(t *testing.T) {
	filename := "ints_big_testfile"
	count := 6000
	fields := 78

	data := createIntData(count, fields)
	doIntsWrite(t, filename, data)
	doIntsRead(t, filename, data)
}

func TestAllBigInts(t *testing.T) {
	filename := "all_ints_big_testfile"
	count := 6000
	fields := 78

	data := createIntData(count, fields)
	doAllIntsWrite(t, filename, data)
	doAllIntsRead(t, filename, data)
}

func createFloatData(count, fields int) [][]float64 {
	data := make([][]float64, count)
	for i := range data {
		entry := make([]float64, fields)
		for j := range entry {
			entry[j] = rand.Float64()
		}
		data[i] = entry
	}
	return data
}

func createIntData(count, fields int) [][]int {
	data := make([][]int, count)
	for i := range data {
		entry := make([]int, fields)
		for j := range entry {
			entry[j] = rand.Int()
		}
		data[i] = entry
	}
	return data
}

func doFloatsWrite(t *testing.T, filename string, data [][]float64) {
	file, err := os.Create(filename)
	if err != nil {
		t.Error("couldn't create testfile", filename, err)
	}
	defer file.Close()

	w := csv.NewWriter(file)

	for _, entry := range data {
		err = WriteFloats(w, entry)
		if err != nil {
			t.Error("couldn't write float entry", entry, err)
		}
	}

	w.Flush()
}

func doAllFloatsWrite(t *testing.T, filename string, data [][]float64) {
	file, err := os.Create(filename)
	if err != nil {
		t.Error("couldn't create testfile", filename, err)
	}
	defer file.Close()

	w := csv.NewWriter(file)

	err = WriteAllFloats(w, data)
	if err != nil {
		t.Error("couldn't write float entries", data, err)
	}

	w.Flush()
}

func doIntsWrite(t *testing.T, filename string, data [][]int) {
	file, err := os.Create(filename)
	if err != nil {
		t.Error("couldn't create testfile", filename, err)
	}
	defer file.Close()

	w := csv.NewWriter(file)

	for _, entry := range data {
		err = WriteInts(w, entry)
		if err != nil {
			t.Error("couldn't write ints entry", entry, err)
		}
	}

	w.Flush()
}

func doAllIntsWrite(t *testing.T, filename string, data [][]int) {
	file, err := os.Create(filename)
	if err != nil {
		t.Error("couldn't create testfile", filename, err)
	}
	defer file.Close()

	w := csv.NewWriter(file)

	err = WriteAllInts(w, data)
	if err != nil {
		t.Error("couldn't write int entries", data, err)
	}

	w.Flush()
}

func doFloatsRead(t *testing.T, filename string, data [][]float64) {
	file, err := os.Open(filename)
	if err != nil {
		t.Error("couldn't open testfile", filename, err)
	}
	defer file.Close()

	r := csv.NewReader(file)

	for _, expectedEntry := range data {
		entry, err := ReadFloats(r)
		if err != nil {
			t.Error("couldn't read float entry", err)
		}

		if len(entry) != len(expectedEntry) {
			t.Error("entry size doesn't match", len(entry), len(expectedEntry))
		} else {
			for i, value := range entry {
				if value != expectedEntry[i] {
					t.Error("entry value doesn't match", value, expectedEntry[i])
				}
			}
		}
	}

	_, err = ReadFloats(r)
	if err == nil {
		t.Error("expecting error from ReadFloats")
	}
}

func doAllFloatsRead(t *testing.T, filename string, data [][]float64) {
	file, err := os.Open(filename)
	if err != nil {
		t.Error("couldn't open testfile", filename, err)
	}
	defer file.Close()

	r := csv.NewReader(file)

	allEntries, err := ReadAllFloats(r)
	if err != nil {
		t.Error("couldn't read float entries", err)
	}

	for i, expected := range data {
		if len(expected) != len(allEntries[i]) {
			t.Error("entry size doesn't match", len(allEntries[i]), len(expected))
		}
		for j := range expected {
			if expected[j] != allEntries[i][j] {
				t.Error("entry value doesn't match", allEntries[i][j], expected[j])
			}
		}
	}
}

func doIntsRead(t *testing.T, filename string, data [][]int) {
	file, err := os.Open(filename)
	if err != nil {
		t.Error("couldn't open testfile", filename, err)
	}
	defer file.Close()

	r := csv.NewReader(file)

	for _, expectedEntry := range data {
		entry, err := ReadInts(r)
		if err != nil {
			t.Error("couldn't read float entry", err)
		}

		if len(entry) != len(expectedEntry) {
			t.Error("entry size doesn't match", len(entry), len(expectedEntry))
		} else {
			for i, value := range entry {
				if value != expectedEntry[i] {
					t.Error("entry value doesn't match", value, expectedEntry[i])
				}
			}
		}
	}

	_, err = ReadInts(r)
	if err == nil {
		t.Error("expecting error from ReadFloats")
	}
}

func doAllIntsRead(t *testing.T, filename string, data [][]int) {
	file, err := os.Open(filename)
	if err != nil {
		t.Error("couldn't open testfile", filename, err)
	}
	defer file.Close()

	r := csv.NewReader(file)

	allEntries, err := ReadAllInts(r)
	if err != nil {
		t.Error("couldn't read int entries", err)
	}

	for i, expected := range data {
		if len(expected) != len(allEntries[i]) {
			t.Error("entry size doesn't match", len(allEntries[i]), len(expected))
		}
		for j := range expected {
			if expected[j] != allEntries[i][j] {
				t.Error("entry value doesn't match", allEntries[i][j], expected[j])
			}
		}
	}
}
