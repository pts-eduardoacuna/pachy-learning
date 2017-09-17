package main

import (
	"flag"
	"fmt"
)

func main() {
	// Handle command line arguments
	var digitsPngDir string
	var digitsCsvDir string

	flag.StringVar(&digitsPngDir, "input-digits-png", "", "The directory containing the image files as PNG.")
	flag.StringVar(&digitsCsvDir, "output-digits-csv", "", "The digits as CSV output directory.")

	flag.Parse()

	fmt.Printf("Program `image` called with arguments input-digits-png=`%s` and output-digits-csv=`%s`\n", digitsPngDir, digitsCsvDir)
}
