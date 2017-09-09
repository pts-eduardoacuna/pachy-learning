// Package learning allows training and inference with ML models.
package learning

import (
	"math"
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

// NeuralNetwork is a supervised learning model for classification.
type NeuralNetwork struct {
	Signals         []*mat.Dense
	Outputs         []*mat.Dense
	Weights         []*mat.Dense
	Deltas          []*mat.Dense
	Gradients       []*mat.Dense
	ErrorHistory    *mat.Dense
	LayerCount      int
	LearningRate    float64
	attributesSize  int
	predictionsSize int
	predictionError *mat.Dense
}

// NewNeuralNetwork creates a NeuralNetwork with the given learning rate and architecture.
//
// The architecture consists of at least two elements, where each element specifies the amount of
// nodes in the layers. The first number corresponds to the input layer and the last to the output
// layer.
func NewNeuralNetwork(learningRate float64, arch []int) *NeuralNetwork {
	layerCount := len(arch)

	signals := make([]*mat.Dense, layerCount)
	outputs := make([]*mat.Dense, layerCount)
	weights := make([]*mat.Dense, layerCount)
	deltas := make([]*mat.Dense, layerCount)
	gradients := make([]*mat.Dense, layerCount)

	outputs[0] = mat.NewDense(arch[0]+1, 1, nil)
	outputs[0].Set(0, 0, 1)
	for layer := 1; layer < layerCount; layer++ {
		signals[layer] = mat.NewDense(arch[layer], 1, nil)
		outputs[layer] = mat.NewDense(arch[layer]+1, 1, nil)
		outputs[layer].Set(0, 0, 1)
		weights[layer] = newRandomWeightMatrix(arch[layer-1]+1, arch[layer])
		deltas[layer] = mat.NewDense(arch[layer], 1, nil)
		gradients[layer] = mat.NewDense(arch[layer-1]+1, arch[layer], nil)
	}

	attributesSize := arch[0]
	predictionsSize := arch[len(arch)-1]

	return &NeuralNetwork{
		Signals:         signals,
		Outputs:         outputs,
		Weights:         weights,
		Deltas:          deltas,
		Gradients:       gradients,
		LayerCount:      layerCount,
		LearningRate:    learningRate,
		attributesSize:  attributesSize,
		predictionsSize: predictionsSize,
	}
}

// Train adjusts the parameters of a neural network to fit the attributes dataset with
// the targets dataset.
//
// Both datasets must have the same number of rows, and their columns should match the
// dimension of the first and last layer of the network.
func Train(net *NeuralNetwork, attributesSet, targetsSet *mat.Dense) {
	trainingSize, _ := attributesSet.Dims()
	attributesSize := net.attributesSize
	targetsSize := net.predictionsSize

	realOutput := mat.NewDense(targetsSize, 1, nil)
	errorHistory := mat.NewDense(trainingSize, targetsSize, nil)

	for row := 0; row < trainingSize; row++ {

		for attribute := 0; attribute < attributesSize; attribute++ {
			net.Outputs[0].Set(attribute+1, 0, attributesSet.At(row, attribute))
		}

		for target := 0; target < targetsSize; target++ {
			realOutput.Set(target, 0, targetsSet.Slice(row, row+1, 0, targetsSize).At(0, target))
		}

		forwardPropagation(net)
		backPropagation(net, realOutput)
		gradientDescentStep(net)

		learningError(errorHistory, row, realOutput, net.Outputs[net.LayerCount-1].Slice(1, targetsSize+1, 0, 1))
	}

	net.ErrorHistory = errorHistory
}

// Infer user the network to evaluate each row in the attributes dataset.
func Infer(net *NeuralNetwork, attributesSet *mat.Dense) *mat.Dense {
	datasetSize, _ := attributesSet.Dims()
	attributesSize := net.attributesSize
	predictionsSize := net.predictionsSize

	predictionsSet := mat.NewDense(datasetSize, predictionsSize, nil)

	for row := 0; row < datasetSize; row++ {

		for attribute := 0; attribute < attributesSize; attribute++ {
			net.Outputs[0].Set(attribute+1, 0, attributesSet.At(row, attribute))
		}

		forwardPropagation(net)

		for prediction := 0; prediction < predictionsSize; prediction++ {
			predictionsSet.Set(row, prediction, net.Outputs[net.LayerCount-1].At(prediction+1, 0))
		}

	}

	return predictionsSet
}

func newRandomWeightMatrix(rows, cols int) *mat.Dense {
	data := make([]float64, rows*cols)
	for i := range data {
		w := rand.Float64()
		if i%2 == 0 {
			w = -w
		}
		data[i] = w
	}

	return mat.NewDense(rows, cols, data)
}

func forwardPropagation(net *NeuralNetwork) {
	for layer := 1; layer < net.LayerCount; layer++ {
		computeSignalForward(net.Signals[layer], net.Weights[layer], net.Outputs[layer-1])
		computeOutputForward(net.Outputs[layer], net.Signals[layer])
	}
}

func backPropagation(net *NeuralNetwork, targets *mat.Dense) {
	var outputs *mat.Dense
	var deltas *mat.Dense
	var nextDeltas *mat.Dense
	var nextWeights *mat.Dense

	// Setup deltas in last layer
	outputs = net.Outputs[net.LayerCount-1]
	deltas = net.Deltas[net.LayerCount-1]
	rows, _ := deltas.Dims()
	for i := 0; i < rows; i++ {
		deltas.Set(i, 0, (outputs.At(i+1, 0)-targets.At(i, 0))*bernoulliVariance(outputs.At(i+1, 0)))
	}

	// Propagate error backwards
	for layer := net.LayerCount - 2; layer > 0; layer-- {
		deltas = net.Deltas[layer]
		nextDeltas = net.Deltas[layer+1]
		outputs = net.Outputs[layer]
		nextWeights = net.Weights[layer+1]

		rows, cols := nextWeights.Dims()
		deltas.Mul(nextWeights.Slice(1, rows, 0, cols).(*mat.Dense), nextDeltas)

		rows, _ = deltas.Dims()
		for i := 0; i < rows; i++ {
			deltas.Set(i, 0, deltas.At(i, 0)*bernoulliVariance(outputs.At(i+1, 0)))
		}
	}
}

func gradientDescentStep(net *NeuralNetwork) {
	var weights *mat.Dense
	var gradient *mat.Dense
	var deltas *mat.Dense
	var previousOutput *mat.Dense

	for layer := 1; layer < net.LayerCount; layer++ {
		weights = net.Weights[layer]
		gradient = net.Gradients[layer]
		deltas = net.Deltas[layer]
		previousOutput = net.Outputs[layer-1]

		gradient.Mul(previousOutput, deltas.T())
		gradient.Scale(net.LearningRate, gradient)
		weights.Sub(weights, gradient)
	}
}

func learningError(history *mat.Dense, row int, target, prediction mat.Matrix) {
	_, cols := history.Dims()
	for i := 0; i < cols; i++ {
		history.Set(row, i, 0.5*square(target.At(i, 0)-prediction.At(i, 0)))
	}
}

func computeSignalForward(signal, weight, output *mat.Dense) {
	signal.Mul(weight.T(), output)
}

func computeOutputForward(output, signal *mat.Dense) {
	output.Set(0, 0, 1)
	rows, _ := output.Dims()
	for i := 1; i < rows; i++ {
		output.Set(i, 0, logistic(signal.At(i-1, 0)))
	}
}

func logistic(z float64) float64 {
	return 1 / (1 + math.Exp(-z))
}

func bernoulliVariance(p float64) float64 {
	return p * (1 - p)
}

func square(x float64) float64 {
	return x * x
}
