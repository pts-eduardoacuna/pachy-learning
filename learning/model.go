package learning

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
)

// Analysis holds the results of training a variety of neural network architectures with different learning rates.
//
// It holds the result of the best architecture and learning rate according to a validation analysis, and also
// the data associated with each model that was analyzed.
type Analysis struct {
	Data []AnalysisValidation `json:"validation"`
	Best Model                `json:"model"`
}

// NewAnalysis creates an Analysis reference from the given validation data and selected model.
func NewAnalysis(validationData []AnalysisValidation, selected Model) *Analysis {
	return &Analysis{
		Data: validationData,
		Best: selected,
	}
}

// Model describes a description of a neural network.
type Model struct {
	LearningRate float64 `json:"learningRate"`
	Architecture []int   `json:"architecture"`
}

// NewModel creates a neural network Model from the given learning rate and architecture.
func NewModel(rate float64, arch []int) Model {
	return Model{
		LearningRate: rate,
		Architecture: arch,
	}
}

// AnalysisValidation holds the validation results associated with a learning rate.
type AnalysisValidation struct {
	LearningRate float64                    `json:"learningRate"`
	Models       []AnalysisValidationResult `json:"models"`
}

// AnalysisValidationResult hold the training and validation error associated with a neural network architecture.
type AnalysisValidationResult struct {
	Architecture    []int   `json:"architecture"`
	TrainingError   float64 `json:"trainingError"`
	ValidationError float64 `json:"validationError"`
}

// SelectBestModel receives a slice of AnalysisValidation and returns the model which has the minimum validation error.
func SelectBestModel(validationData []AnalysisValidation) (Model, error) {
	rate := 0.0
	arch := []int{}
	minerr := math.MaxFloat64

	for _, rateGroup := range validationData {
		for _, resultGroup := range rateGroup.Models {
			if resultGroup.ValidationError < minerr {
				rate = rateGroup.LearningRate
				arch = resultGroup.Architecture
				minerr = resultGroup.ValidationError
			}
		}
	}

	model := NewModel(rate, arch)

	if minerr == math.MaxFloat64 {
		return model, fmt.Errorf("all the models stink with validation error %f", minerr)
	}

	return model, nil
}

// WriteAnalysis writes a neural network analysis in a JSON encoding to a file.
func WriteAnalysis(analysis *Analysis, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	bytes, err := json.MarshalIndent(analysis, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}

// ReadAnalysis reads a neural network analysis in a JSON encoding from a file.
func ReadAnalysis(path string) (*Analysis, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	bytes := []byte{}
	_, err = file.Read(bytes)
	if err != nil {
		return nil, err
	}

	analysis := &Analysis{}
	err = json.Unmarshal(bytes, analysis)
	if err != nil {
		return nil, err
	}

	return analysis, nil
}
