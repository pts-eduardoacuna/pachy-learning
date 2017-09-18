package main

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/pts-eduardoacuna/pachy-learning/csnv"
	"github.com/pts-eduardoacuna/pachy-learning/gob"
	"github.com/pts-eduardoacuna/pachy-learning/learning"
	"github.com/pts-eduardoacuna/pachy-learning/log"

	"gonum.org/v1/gonum/mat"
)

const (
	trainingFilename = "train"
	testingFilename  = "test"
	analysisFilename = "validation"
	modelFilename    = "model"
)

func processDataset(filepath string) (*mat.Dense, *mat.Dense, error) {
	// Open input file
	log.Printf("opening input file: path=%v", filepath)
	file, err := os.Open(filepath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	// Read its content
	log.Printf("reading %v as CSV", filepath)
	attributes, targets, err := csnv.ReadDataset(file)

	return attributes, targets, err
}

func main() {
	var err error

	// Handle command line arguments
	var mnistCsvDir string
	var analysisDir string
	var modelDir string

	flag.StringVar(&mnistCsvDir, "input-mnist-csv", ".", "The directory containing the MNIST files as CSV.")
	flag.StringVar(&analysisDir, "input-analysis", ".", "The directory containing the ANN model analysis.")
	flag.StringVar(&modelDir, "output-model", ".", "The model output directory.")

	flag.Parse()

	// Initialize logger
	log.ToFile(filepath.Join(modelDir, "log"))
	defer log.Close()

	// Process training and testing datasets
	trainingPath := filepath.Join(mnistCsvDir, trainingFilename)
	testingPath := filepath.Join(mnistCsvDir, testingFilename)
	analysisPath := filepath.Join(analysisDir, analysisFilename)
	modelPath := filepath.Join(modelDir, modelFilename)

	trainingAttributes, trainingTargets, err := processDataset(trainingPath)
	if err != nil {
		log.Fatalf("there were problems processing the training dataset: error=%v", err)
	}
	if trainingAttributes == nil || trainingTargets == nil {
		log.Fatalf("there were problems processing the training dataset: nil matrices")
	}

	testingAttributes, testingTargets, err := processDataset(testingPath)
	if err != nil {
		log.Fatalf("there were problems processing the testing dataset: error=%v", err)
	}
	if testingAttributes == nil || testingTargets == nil {
		log.Fatalf("there were problems processing the testing dataset: nil matrices")
	}

	log.Printf("reading neural network analysis: path=%v", analysisPath)
	analysis, err := learning.ReadAnalysis(analysisPath)
	if err != nil {
		log.Fatalf("there were problems processing the analysis data: error=%v", err)
	}

	rate := analysis.Best.LearningRate
	arch := analysis.Best.Architecture
	log.Printf("creating neural network: learning rate=%f, architecture=%v", rate, arch)
	net, err := learning.NewNeuralNetwork(rate, arch)
	if err != nil {
		log.Fatalf("there were problems creating the neural network: error=%v", err)
	}

	log.Printf("training neural network")
	err = learning.Train(net, trainingAttributes, trainingTargets)
	if err != nil {
		log.Fatalf("there were problems training the neural network: error=%v", err)
	}

	log.Printf("computing the testing error")
	testingError, err := learning.ComputeError(net, testingAttributes, testingTargets)
	if err != nil {
		log.Fatalf("there were problems computing the testing error: error=%v", err)
	}

	log.Printf("network testing error=%f", testingError)
	net.TestingError = testingError

	log.Printf("serializing model: path=%v", modelPath)
	file, err := os.Create(modelPath)
	if err != nil {
		log.Fatalf("there were problems creating the model file: error=%v", err)
	}
	defer file.Close()

	err = gob.WriteBinaryObject(file, net)
	if err != nil {
		log.Fatalf("there were problems creating the model file: error=%v", err)
	}
}
