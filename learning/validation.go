package learning

import (
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

// SplitTrainingValidation chooses a random sample of training attributes for constructing a validation set.
func SplitTrainingValidation(validationPercent int, inputs, outputs *mat.Dense) (*mat.Dense, *mat.Dense, *mat.Dense, *mat.Dense) {
	trainingSize, inputSize := inputs.Dims()
	_, outputSize := outputs.Dims()

	validationSize := validationPercent * trainingSize / 100

	randomRows := rand.Perm(trainingSize)
	trainingRows := randomRows[0 : trainingSize-validationSize]
	validationRows := randomRows[trainingSize-validationSize : trainingSize]

	tInputs := mat.NewDense(trainingSize-validationSize, inputSize, nil)
	tOutputs := mat.NewDense(trainingSize-validationSize, outputSize, nil)
	vInputs := mat.NewDense(validationSize, inputSize, nil)
	vOutputs := mat.NewDense(validationSize, outputSize, nil)

	for i, row := range trainingRows {
		tInputs.SetRow(i, inputs.RawRowView(row))
		tOutputs.SetRow(i, outputs.RawRowView(row))
	}

	for i, row := range validationRows {
		vInputs.SetRow(i, inputs.RawRowView(row))
		vOutputs.SetRow(i, outputs.RawRowView(row))
	}

	return tInputs, tOutputs, vInputs, vOutputs
}

// ValidateNeuralNetwork runs a simple training and inference check on a training and validation set.
func ValidateNeuralNetwork(tInputs, tOutputs, vInputs, vOutputs *mat.Dense, rate float64, arch []int) (float64, float64, error) {
	net, err := NewNeuralNetwork(rate, arch)
	if err != nil {
		return 0, 0, err
	}

	err = Train(net, tInputs, tOutputs)
	if err != nil {
		return 0, 0, err
	}

	tError, err := ComputeError(net, tInputs, tOutputs)
	if err != nil {
		return 0, 0, err
	}

	vError, err := ComputeError(net, vInputs, vOutputs)
	if err != nil {
		return 0, 0, err
	}

	return tError, vError, nil
}

// ComputeError checks how many unsuccessful predictions the neural network makes.
func ComputeError(net *NeuralNetwork, inputs, expected *mat.Dense) (float64, error) {
	predictions, err := Infer(net, inputs)
	if err != nil {
		return 0, err
	}

	rows, _ := expected.Dims()
	score := 0
	for i := 0; i < rows; i++ {
		correct := DecodeTargets(expected.RawRowView(i))
		predicted := DecodeTargets(predictions.RawRowView(i))

		if correct == predicted {
			score++
		}
	}

	incorrect := 1.0 - (float64(score) / float64(rows))

	return incorrect, nil
}
