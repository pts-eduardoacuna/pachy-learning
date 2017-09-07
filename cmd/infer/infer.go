package main

import (
	"flag"
	"fmt"
)

func main() {
	// Handle command line arguments
	var modelDir string
	var attributesDir string
	var inferenceDir string

	flag.StringVar(&modelDir, "input-model", "", "The directory containing the model files.")
	flag.StringVar(&attributesDir, "input-attributes", "", "The directory containing the attributes files.")
	flag.StringVar(&inferenceDir, "output-inference", "", "The inference data output directory.")

	flag.Parse()

	fmt.Printf("Program `infer` called with arguments input-model=`%s`, input-attributes=`%s`, and output-inference=`%s`\n", modelDir, attributesDir, inferenceDir)
}
