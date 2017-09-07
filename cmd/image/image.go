package main

import (
	"flag"
)

func main() {
	// Handle command line arguments
	var inferenceDir string
	var graphDir string

	flag.StringVar(&inferenceDir, "input-inference", "", "The directory containing the inferences files.")
	flag.StringVar(&graphDir, "output-graph", "", "The graphs data output directory.")

	flag.Parse()
}
