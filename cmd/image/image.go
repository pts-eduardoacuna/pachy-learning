package main

import (
	"flag"
	"fmt"
)

func main() {
	// Handle command line arguments
	var inferenceDir string
	var graphDir string

	flag.StringVar(&inferenceDir, "input-inference", "", "The directory containing the inferences files.")
	flag.StringVar(&graphDir, "output-graph", "", "The graphs data output directory.")

	flag.Parse()

	fmt.Printf("Program `image` called with arguments input-inference=`%s` and output-graph=`%s`\n", inferenceDir, graphDir)
}
