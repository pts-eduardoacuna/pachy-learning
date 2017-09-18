package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pts-eduardoacuna/pachy-learning/csnv"
	"github.com/pts-eduardoacuna/pachy-learning/log"
	"github.com/pts-eduardoacuna/pachy-learning/mnist"
)

const (
	trainingFilename       = "train"
	trainingImagesFilename = "train-images-idx3-ubyte"
	trainingLabelsFilename = "train-labels-idx1-ubyte"
	testingFilename        = "test"
	testingImagesFilename  = "t10k-images-idx3-ubyte"
	testingLabelsFilename  = "t10k-labels-idx1-ubyte"
)

func processDataset(inputImagesName, inputLabelsName, outputName string) error {
	// Open input files
	log.Printf("opening input file: path=%s", inputImagesName)
	inputImages, err := os.Open(inputImagesName)
	if err != nil {
		return err
	}
	defer inputImages.Close()

	log.Print("opening input file: path=%s", inputLabelsName)
	inputLabels, err := os.Open(inputLabelsName)
	if err != nil {
		return err
	}
	defer inputLabels.Close()

	// Create output files
	log.Print("creating output file: path=%s", outputName)
	output, err := os.Create(outputName)
	if err != nil {
		return err
	}
	defer output.Close()

	// Make input parsers
	log.Print("making image parser")
	imageParser, err := mnist.NewImageParser(inputImages)
	if err != nil {
		return err
	}

	log.Print("making labels parser")
	labelParser, err := mnist.NewLabelParser(inputLabels)
	if err != nil {
		return err
	}

	if imageParser.Count != labelParser.Count {
		return fmt.Errorf("expecting count for %d images and %d labels to be the same", imageParser.Count, labelParser.Count)
	}

	// Make output writer
	log.Print("making output CSV writer")
	outputWriter := csv.NewWriter(output)

	// Transform inputs into outputs
	log.Print("processing images with labels")
	for i := 0; i < imageParser.Count; i++ {
		img, err := imageParser.Parse()
		if err != nil {
			return err
		}
		lbl, err := labelParser.Parse()
		if err != nil {
			return err
		}
		entry := append(img, lbl)
		csnv.WriteInts(outputWriter, entry)
		log.Printf("entry %d written (%d%%)", i+1, int(100.0*float64(i+1)/float64(imageParser.Count)))
	}

	outputWriter.Flush()

	return nil
}

func main() {
	var err error

	// Handle command line arguments
	var mnistDir string
	var mnistCsvDir string

	flag.StringVar(&mnistDir, "input-mnist", ".", "The directory containing the uncompressed MNIST data files.")
	flag.StringVar(&mnistCsvDir, "output-mnist-csv", ".", "The mnist-csv data output directory.")

	flag.Parse()

	log.ToFile(filepath.Join(mnistCsvDir, "log"))
	defer log.Close()

	// Process training and testing datasets
	trainingImagesPath := filepath.Join(mnistDir, trainingImagesFilename)
	trainingLabelsPath := filepath.Join(mnistDir, trainingLabelsFilename)
	trainingPath := filepath.Join(mnistCsvDir, trainingFilename)

	testingImagesPath := filepath.Join(mnistDir, testingImagesFilename)
	testingLabelsPath := filepath.Join(mnistDir, testingLabelsFilename)
	testingPath := filepath.Join(mnistCsvDir, testingFilename)

	log.Print("processing training dataset")
	err = processDataset(trainingImagesPath, trainingLabelsPath, trainingPath)
	if err != nil {
		log.Fatalf("there were problems processing the training dataset, this is essential for training a model: error=%v", err)
	}

	log.Print("processing testing dataset")
	err = processDataset(testingImagesPath, testingLabelsPath, testingPath)
	if err != nil {
		log.Fatalf("there were problems processing the testing dataset, this is important for measuring a model's peformance: error=%v", err)
	}
}
