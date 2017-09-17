// Package gob allows serializing and deserializing structures.
package gob

import (
	"encoding/gob"
	"os"
)

// WriteBinaryObject writes a gob of the given object in the given file.
func WriteBinaryObject(file *os.File, obj interface{}) error {
	enc := gob.NewEncoder(file)
	return enc.Encode(obj)
}

// ReadBinaryObject reads the gob file and puts its content on the given object.
func ReadBinaryObject(file *os.File, obj interface{}) error {
	dec := gob.NewDecoder(file)
	return dec.Decode(obj)
}
