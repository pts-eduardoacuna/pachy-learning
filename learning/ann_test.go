package learning

import (
	"math/rand"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func IsLearning(N int) (*mat.Dense, error) {
	net, err := NewNeuralNetwork(0.2, []int{1, 1})
	if err != nil {
		return nil, err
	}

	attributesSet := mat.NewDense(N, 1, nil)
	rows, _ := attributesSet.Dims()
	for i := 0; i < rows; i++ {
		attributesSet.Set(i, 0, rand.Float64())
	}

	targetsSet := mat.NewDense(N, 1, nil)
	targetsSet.Apply(func(i, j int, v float64) float64 {
		if v < 0.5 {
			return 0.0
		}
		return 1.0
	}, attributesSet)

	Infer(net, attributesSet)

	err = Train(net, attributesSet, targetsSet)
	if err != nil {
		return nil, err
	}

	return net.ErrorHistory, nil
}

func TestIsLearning(t *testing.T) {
	N := 1000
	errorHistory, err := IsLearning(N)

	if err != nil {
		t.Error("couldn't create and train a neural network", err)
	}

	errorA := errorHistory.At(0, 0)
	errorB := errorHistory.At((N-1)/2, 0)
	errorC := errorHistory.At(N-1, 0)

	if errorA <= errorB || errorB <= errorC {
		t.Error("couldn't make a neural network learn")
	}
}

func BenchmarkIsLearning(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsLearning(1000)
	}
}

func LearningXOR() (*mat.Dense, error) {
	net, err := NewNeuralNetwork(0.2, []int{2, 2, 2})
	if err != nil {
		return nil, err
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
		err = Train(net, attributesSet, targetsSet)
		if err != nil {
			return nil, err
		}
	}

	output, err := Infer(net, attributesSet)

	return output, err
}

func TestLearningXOR(t *testing.T) {
	predictionsSet, err := LearningXOR()

	if err != nil {
		t.Error("couldn't create and train a neural network", err)
	}

	if predictionsSet.At(0, 0) < predictionsSet.At(0, 1) {
		t.Error("couldn't solve XOR problem")
	}

	if predictionsSet.At(1, 0) > predictionsSet.At(1, 1) {
		t.Error("couldn't solve XOR problem")
	}

	if predictionsSet.At(2, 0) < predictionsSet.At(2, 1) {
		t.Error("couldn't solve XOR problem")
	}

	if predictionsSet.At(3, 0) > predictionsSet.At(3, 1) {
		t.Error("couldn't solve XOR problem")
	}
}

func BenchmarkLearningXOR(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LearningXOR()
	}
}
