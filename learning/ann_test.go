package learning

import (
	"math/rand"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func IsLearning(N int) *mat.Dense {
	net := NewNeuralNetwork(0.2, []int{1, 1})

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

	Train(net, attributesSet, targetsSet)

	return net.ErrorHistory
}

func TestIsLearning(t *testing.T) {
	N := 1000
	errorHistory := IsLearning(N)

	if errorHistory == nil {
		t.Fail()
	}

	errorA := errorHistory.At(0, 0)
	errorB := errorHistory.At((N-1)/2, 0)
	errorC := errorHistory.At(N-1, 0)

	if errorA <= errorB || errorB <= errorC {
		t.Fail()
	}
}

func BenchmarkIsLearning(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsLearning(1000)
	}
}

func LearningXOR() *mat.Dense {
	net := NewNeuralNetwork(0.2, []int{2, 2, 2})

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
		Train(net, attributesSet, targetsSet)
	}

	return Infer(net, attributesSet)
}

func TestLearningXOR(t *testing.T) {
	predictionsSet := LearningXOR()

	if predictionsSet.At(0, 0) < predictionsSet.At(0, 1) {
		t.Fail()
	}

	if predictionsSet.At(1, 0) > predictionsSet.At(1, 1) {
		t.Fail()
	}

	if predictionsSet.At(2, 0) < predictionsSet.At(2, 1) {
		t.Fail()
	}

	if predictionsSet.At(3, 0) > predictionsSet.At(3, 1) {
		t.Fail()
	}
}

func BenchmarkLearningXOR(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LearningXOR()
	}
}
