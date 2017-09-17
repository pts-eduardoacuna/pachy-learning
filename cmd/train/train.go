package main

import (
	"flag"
	"fmt"
)

func main() {
	// Handle command line arguments
	var mnistCsvDir string
	var analysisDir string
	var modelDir string

	flag.StringVar(&mnistCsvDir, "input-mnist-csv", "", "The directory containing the MNIST files as CSV.")
	flag.StringVar(&analysisDir, "input-analysis", "", "The directory containing the ANN model analysis.")
	flag.StringVar(&modelDir, "output-model", "", "The model output directory.")

	flag.Parse()

	fmt.Printf("Program `train` called with arguments input-mnist-csv=`%s`, input-analysis=`%s`, and output-model=`%s`\n", mnistCsvDir, analysisDir, modelDir)
}
