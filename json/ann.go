// Package json allows reading and writing complex structures as JSON.
package json

import (
	"encoding/json"
	"os"

	"github.com/pts-eduardoacuna/pachy-learning/learning"

	"gonum.org/v1/gonum/mat"
)

type neuralNetworkWeights struct {
	Data []float64 `json:"data"`
}

type neuralNetworkModel struct {
	Architecture []int                  `json:"architecture"`
	LearningRate float64                `json:"learningRate"`
	Weights      []neuralNetworkWeights `json:"weights"`
}

// WriteNeuralNetworkModel writes a neural network as a JSON in the given file.
func WriteNeuralNetworkModel(file *os.File, net *learning.NeuralNetwork) error {
	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")

	model := neuralNetworkToModel(net)

	return enc.Encode(model)
}

// ReadNeuralNetworkModel reads a neural network as a JSON from the given file.
func ReadNeuralNetworkModel(file *os.File) (*learning.NeuralNetwork, error) {
	dec := json.NewDecoder(file)

	var model neuralNetworkModel

	err := dec.Decode(&model)
	if err != nil {
		return nil, err
	}

	net, err := neuralNetworkFromModel(&model)
	if err != nil {
		return nil, err
	}

	return net, nil
}

func neuralNetworkFromModel(model *neuralNetworkModel) (*learning.NeuralNetwork, error) {
	net, err := learning.NewNeuralNetwork(model.LearningRate, model.Architecture)

	if err != nil {
		return nil, err
	}

	for i, ws := range model.Weights {
		net.Weights[i+1] = mat.NewDense(model.Architecture[i]+1, model.Architecture[i+1], ws.Data)
	}

	return net, nil
}

func neuralNetworkToModel(net *learning.NeuralNetwork) *neuralNetworkModel {
	weights := make([]neuralNetworkWeights, net.LayerCount-1)

	architecture := make([]int, net.LayerCount)
	architecture[0] = net.AttributesSize

	for i := 1; i < net.LayerCount; i++ {
		ws := net.Weights[i]
		rows, cols := ws.Dims()
		architecture[i] = cols
		weights[i-1].Data = make([]float64, rows*cols)
		for row := 0; row < rows; row++ {
			for col := 0; col < cols; col++ {
				weights[i-1].Data[cols*row+col] = ws.At(row, col)
			}
		}
	}

	model := neuralNetworkModel{
		Architecture: architecture,
		LearningRate: net.LearningRate,
		Weights:      weights,
	}

	return &model
}
