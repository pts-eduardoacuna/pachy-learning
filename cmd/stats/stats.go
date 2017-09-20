package main

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/pts-eduardoacuna/pachy-learning/csnv"
	"github.com/pts-eduardoacuna/pachy-learning/learning"
	"github.com/pts-eduardoacuna/pachy-learning/log"

	"gonum.org/v1/gonum/mat"
)

const (
	trainingFilename = "train"
	analysisFilename = "validation"
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
	attributes, targets, err := csnv.ReadDataset(file)

	return attributes, targets, nil
}

func validateModels(inputs, outputs *mat.Dense, rates []float64, hiddenLayers [][]int) ([]learning.AnalysisValidation, error) {
	log.Printf("validating models")
	_, inputSize := inputs.Dims()
	_, outputSize := outputs.Dims()

	percent := 30
	totalModels := len(rates) * len(hiddenLayers)
	modelCount := 0

	log.Printf("randomly separating training (%d%%) and validation (%d%%) data", 100-percent, percent)

	tInputs, tOutputs, vInputs, vOutputs := learning.SplitTrainingValidation(percent, inputs, outputs)

	validation := make([]learning.AnalysisValidation, len(rates))
	var err error

	for rateIdx, rate := range rates {
		validation[rateIdx].LearningRate = rate
		validation[rateIdx].Models = make([]learning.AnalysisValidationResult, len(hiddenLayers))

		for hiddenIdx, hidden := range hiddenLayers {
			arch := append([]int{inputSize}, hidden...)
			arch = append(arch, outputSize)
			modelCount++

			log.Printf("validating model (%d/%d) with learning rate=%f and architecture=%v", modelCount, totalModels, rate, arch)

			tError, vError, err := learning.ValidateNeuralNetwork(tInputs, tOutputs, vInputs, vOutputs, rate, arch)
			if err != nil {
				return nil, err
			}

			log.Printf("training error = %f\t\tvalidation error = %f", tError, vError)
			validation[rateIdx].Models[hiddenIdx].Architecture = arch
			validation[rateIdx].Models[hiddenIdx].TrainingError = tError
			validation[rateIdx].Models[hiddenIdx].ValidationError = vError
		}
	}

	return validation, err
}

func main() {
	var err error

	// Handle command line arguments
	var mnistCsvDir string
	var analysisDir string

	flag.StringVar(&mnistCsvDir, "input-mnist-csv", ".", "The directory containing the MNIST data files as CSV.")
	flag.StringVar(&analysisDir, "output-analysis", ".", "The model analysis output directory.")

	flag.Parse()

	// Initialize logger
	log.ToFile(filepath.Join(analysisDir, "log"))
	defer log.Close()

	// Process training and testing datasets
	trainingPath := filepath.Join(mnistCsvDir, trainingFilename)
	analysisPath := filepath.Join(analysisDir, analysisFilename)

	trainingAttributes, trainingTargets, err := processDataset(trainingPath)
	if err != nil {
		log.Fatalf("there were problems processing the training dataset: error=%v", err)
	}

	// Define learning rates and models to try (keep them in ascending order)
	learningRates := []float64{
		//0.01,
		//0.1,
		//0.15,
		0.2,
		//0.25,
		//0.3,
		//0.35,
		//0.4,
		//0.45,
		//0.5,
		//0.55,
		//0.6,
		//0.65,
		//0.7,
		//0.75,
	}

	hiddenLayers := [][]int{
		// architecture   ---  complexity (~ #synapses)
		//[]int{},                   // ~ 7,840
		//[]int{10},                 // ~ 7,940
		//[]int{10, 10},             // ~ 8,040
		//[]int{10, 10, 10},         // ~ 8,140
		//[]int{20},                 // ~ 15,880
		//[]int{20, 20},             // ~ 16,280
		//[]int{20, 20, 20},         // ~ 16,680
		// []int{
		// 	20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
		// 	20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
		// }, // ~ 39,480
		//[]int{50},                 // ~ 39,700
		//[]int{50, 50},             // ~ 42,200
		//[]int{50, 50, 50},         // ~ 44,700
		//[]int{90},                 // ~ 71,460
		//[]int{90, 90},             // ~ 79,560
		//[]int{90, 90, 90},         // ~ 87,660
		//[]int{90, 90, 90, 90},     // ~ 95,760
		[]int{387}, // ~ 307,278
		//[]int{387, 387},           // ~ 457,047
		//[]int{387, 387, 387},      // ~ 606,816
		//[]int{784}, // ~ 622,496
		//[]int{387, 387, 387, 387}, // ~ 756,585
	}

	validation, err := validateModels(trainingAttributes, trainingTargets, learningRates, hiddenLayers)
	if err != nil {
		log.Fatalf("there were problems validating the models: error=%v", err)
	}

	log.Printf("selecting the best model based on the validation error")
	model, err := learning.SelectBestModel(validation)
	if err != nil {
		log.Fatalf("there were problems selecting the best model: error=%v", err)
	}
	log.Printf("best model has learning rate = %f and architecture = %v", model.LearningRate, model.Architecture)

	analysis := learning.NewAnalysis(validation, model)

	err = learning.WriteAnalysis(analysis, analysisPath)
	if err != nil {
		log.Fatalf("there were problems writing the model analysis: error=%v", err)
	}
}
