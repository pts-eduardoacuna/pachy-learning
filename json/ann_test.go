package json

import (
	"os"
	"testing"

	"github.com/pts-eduardoacuna/pachy-learning/learning"

	"gonum.org/v1/gonum/mat"
)

func TestReadWriteNeuralNetworkModel(t *testing.T) {
	filename := "nn_model_testfile"

	net, input, output, err := createTrainedDummyNetwork()
	if err != nil {
		t.Error("couldn't create and train a dummy neural network", err)
	}

	doModelWrite(t, filename, net)
	doModelRead(t, filename, input, output)
}

func doModelWrite(t *testing.T, filename string, net *learning.NeuralNetwork) {
	file, err := os.Create(filename)
	if err != nil {
		t.Error("couldn't create testfile", err)
	}
	defer file.Close()

	err = WriteNeuralNetworkModel(file, net)
	if err != nil {
		t.Error("couldn't write neural network model", err)
	}
}

func doModelRead(t *testing.T, filename string, input, output *mat.Dense) {
	file, err := os.Open(filename)
	if err != nil {
		t.Error("couldn't open testfile", err)
	}
	defer file.Close()

	net, err := ReadNeuralNetworkModel(file)
	if err != nil {
		t.Error("couldn't read neural network model", err)
	}

	verifyNeuralNetworkOutput(t, net, input, output)
}

func verifyNeuralNetworkOutput(t *testing.T, net *learning.NeuralNetwork, input, expected *mat.Dense) {
	output, err := learning.Infer(net, input)
	if err != nil {
		t.Error("couldn't make inference", err)
	}

	rows, cols := output.Dims()
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if output.At(i, j) != expected.At(i, j) {
				t.Error("expecting inference to match")
			}
		}
	}
}

func createTrainedDummyNetwork() (*learning.NeuralNetwork, *mat.Dense, *mat.Dense, error) {
	net, err := learning.NewNeuralNetwork(0.2, []int{2, 2, 2})
	if err != nil {
		return nil, nil, nil, err
	}

	attributesSet := mat.NewDense(4, 2, []float64{
		0, 0,
		0, 1,
		1, 0,
		1, 1,
	})

	targetsSet := mat.NewDense(4, 2, []float64{
		1, 0,
		0, 1,
		1, 0,
		0, 1,
	})

	for i := 0; i < 1000; i++ {
		err = learning.Train(net, attributesSet, targetsSet)
		if err != nil {
			return nil, nil, nil, err
		}
	}

	output, err := learning.Infer(net, attributesSet)

	return net, attributesSet, output, err
}
