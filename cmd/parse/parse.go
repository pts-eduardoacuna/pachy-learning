package main

import (
	"flag"
)

func main() {
	// Handle command line arguments
	var mnistDir string
	var trainingDir string

	flag.StringVar(&mnistDir, "input-mnist", "", "The directory containing the uncompressed MNIST data files.")
	flag.StringVar(&trainingDir, "output-training", "", "The training data output directory.")

	flag.Parse()
}
