package main

import (
	"flag"
	"fmt"
)

func main() {
	// Handle command line arguments
	var trainingDir string
	var modelDir string

	flag.StringVar(&trainingDir, "input-training", "", "The directory containing the training data.")
	flag.StringVar(&modelDir, "output-model", "", "The model output directory.")

	flag.Parse()

	fmt.Printf("Program `train` called with arguments input-training=`%s` and output-model=`%s`\n", trainingDir, modelDir)
}
