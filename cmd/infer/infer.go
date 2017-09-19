package main

import (
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/pts-eduardoacuna/pachy-learning/gob"
	"github.com/pts-eduardoacuna/pachy-learning/image"
	"github.com/pts-eduardoacuna/pachy-learning/learning"
	"github.com/pts-eduardoacuna/pachy-learning/log"
)

const (
	modelFilename = "model"
)

func main() {
	var err error

	// Handle command line arguments
	var modelDir string
	var digitsPngDir string
	var inferenceDir string

	flag.StringVar(&modelDir, "input-model", "", "The directory containing the model file.")
	flag.StringVar(&digitsPngDir, "input-digits-png", "", "The directory containing the image files.")
	flag.StringVar(&inferenceDir, "output-inference", "", "The inference data output directory.")

	flag.Parse()

	// Initialize logger
	log.ToFile(filepath.Join(inferenceDir, "log"))
	defer log.Close()

	// Process model
	modelPath := filepath.Join(modelDir, modelFilename)

	log.Printf("opening neural network model: path=%v", modelPath)
	modelFile, err := os.Open(modelPath)
	if err != nil {
		log.Fatalf("there were problems opening the model file: error=%v", err)
	}
	defer modelFile.Close()

	net := &learning.NeuralNetwork{}
	err = gob.ReadBinaryObject(modelFile, net)
	if err != nil {
		log.Fatalf("there were problems reading the model file: error=%v", err)
	}

	log.Printf("the network has a testing error of %f", net.TestingError)

	log.Printf("traversing images directory: path=%v", digitsPngDir)
	if err := filepath.Walk(digitsPngDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		inpath := filepath.Join(digitsPngDir, info.Name())
		log.Printf("opening file %s", inpath)
		file, err := os.Open(inpath)
		if err != nil {
			return err
		}
		defer file.Close()

		log.Printf("reading image")
		image, err := image.ReadDataset(file)
		if err != nil {
			return err
		}

		log.Printf("infering value")
		prediction, err := learning.Infer(net, image)
		if err != nil {
			return err
		}

		decoded := learning.DecodeTargets(prediction.RawRowView(0))
		log.Printf("the network thinks is a %v", decoded)

		outpath := filepath.Join(inferenceDir, info.Name())
		log.Printf("writing prediction to file: path=%v", outpath)

		str := strconv.FormatFloat(decoded, 'f', 0, 64)
		if err := ioutil.WriteFile(outpath, []byte(str), 0644); err != nil {
			log.Fatalf("there was a problem writing to a file: error=%v", err)
		}

		return nil

	}); err != nil {
		log.Fatalf("there was a problem traversing the image directory: path=%s, error=%v", digitsPngDir, err)
	}
}
