package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/pts-eduardoacuna/pachy-learning/csnv"
	"github.com/pts-eduardoacuna/pachy-learning/mnist"
)

var logger *log.Logger

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
	logger.Print("opening input file", inputImagesName)
	inputImages, err := os.Open(inputImagesName)
	if err != nil {
		return err
	}
	defer inputImages.Close()

	logger.Print("opening input file", inputLabelsName)
	inputLabels, err := os.Open(inputLabelsName)
	if err != nil {
		return err
	}
	defer inputLabels.Close()

	// Create output files
	logger.Print("creating output file", outputName)
	output, err := os.Create(outputName)
	if err != nil {
		return err
	}
	defer output.Close()

	// Make input parsers
	logger.Print("making image parser")
	imageParser, err := mnist.NewImageParser(inputImages)
	if err != nil {
		return err
	}

	logger.Print("making labels parser")
	labelParser, err := mnist.NewLabelParser(inputLabels)
	if err != nil {
		return err
	}

	if imageParser.Count != labelParser.Count {
		return fmt.Errorf("expecting count for %d images and %d labels to be the same", imageParser.Count, labelParser.Count)
	}

	// Make output writer
	logger.Print("making output CSV writer")
	outputWriter := csv.NewWriter(output)

	// Transform inputs into outputs
	logger.Print("processing images with labels")
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
		logger.Printf("entry %d written (%d%%)", i+1, int(100.0*float64(i+1)/float64(imageParser.Count)))
	}
	return nil
}

func main() {
	var err error

	// Handle command line arguments
	var mnistDir string
	var mnistCsvDir string

	flag.StringVar(&mnistDir, "input-mnist", "", "The directory containing the uncompressed MNIST data files.")
	flag.StringVar(&mnistCsvDir, "output-mnist-csv", "", "The mnist-csv data output directory.")

	flag.Parse()

	// Initialize logger
	logfile, err := os.Create(filepath.Join(mnistCsvDir, "log"))
	if err != nil {
		log.Fatal("couldn't create logfile")
	}
	defer logfile.Close()

	logger = createLogger(logfile)

	// Process training and testing datasets
	trainingImagesPath := filepath.Join(mnistDir, trainingImagesFilename)
	trainingLabelsPath := filepath.Join(mnistDir, trainingLabelsFilename)
	trainingPath := filepath.Join(mnistCsvDir, trainingFilename)

	testingImagesPath := filepath.Join(mnistDir, testingImagesFilename)
	testingLabelsPath := filepath.Join(mnistDir, testingLabelsFilename)
	testingPath := filepath.Join(mnistCsvDir, testingFilename)

	logger.Print("processing training dataset")
	err = processDataset(trainingImagesPath, trainingLabelsPath, trainingPath)
	if err != nil {
		logger.Println("there were problems processing the training dataset, this is essential for training a model", err)
	}

	logger.Print("processing testing dataset")
	err = processDataset(testingImagesPath, testingLabelsPath, testingPath)
	if err != nil {
		logger.Println("there were problems processing the testing dataset, this is important for measuring a model's peformance", err)
	}
}

func createLogger(file *os.File) *log.Logger {
	return log.New(file, "❯❯ ", log.LstdFlags|log.Lshortfile)
}
