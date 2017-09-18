package csnv

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/pts-eduardoacuna/pachy-learning/learning"

	"gonum.org/v1/gonum/mat"
)

// ReadDataset reads a numeric CSV file and computes the attributes and targets as gonum's matrices.
func ReadDataset(file *os.File) (*mat.Dense, *mat.Dense, error) {
	r := csv.NewReader(file)
	records, err := ReadAllFloats(r)
	if err != nil {
		return nil, nil, err
	}

	rows := len(records)
	if rows < 1 {
		return nil, nil, fmt.Errorf("expecting at least one record in file")
	}
	columns := len(records[0])
	for _, record := range records {
		if len(record) != columns {
			return nil, nil, fmt.Errorf("malformed records in file, expecting the same amount of entries per record")
		}
	}

	// Construct gonum's dense matrices for attributes and targets
	attributes := mat.NewDense(rows, columns-1, nil)
	targets := mat.NewDense(rows, 10, nil)

	for row, record := range records {
		attributes.SetRow(row, learning.EncodeAttributes(record[0:len(record)-1]))
		targets.SetRow(row, learning.EncodeTarget(record[len(record)-1]))
	}

	return attributes, targets, nil
}
