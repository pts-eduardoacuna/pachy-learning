package main

import (
	"flag"
	"fmt"
)

func main() {
	// Handle command line arguments
	var mnistDir string
	var trainingDir string

	flag.StringVar(&mnistDir, "input-mnist", "", "The directory containing the uncompressed MNIST data files.")
	flag.StringVar(&trainingDir, "output-training", "", "The training data output directory.")

	flag.Parse()

	fmt.Printf("Program `parse` called with arguments input-mnist=`%s` and output-training=`%s`\n", mnistDir, trainingDir)
}
