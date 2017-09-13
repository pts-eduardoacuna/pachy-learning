package gob

import (
	"os"
	"testing"

	"github.com/pts-eduardoacuna/pachy-learning/learning"

	"gonum.org/v1/gonum/mat"
)

func TestBinaryObject(t *testing.T) {
	filename := "neural_network_testfile"

	netBefore, err := learning.NewNeuralNetwork(0.2, []int{4, 2, 1})
	if err != nil {
		t.Error("couldn't create a neural network", err)
	}

	netAfter := &learning.NeuralNetwork{}

	doBinaryObjectWrite(t, filename, netBefore)
	doBinaryObjectRead(t, filename, netAfter)
	doBinaryObjectComparison(t, netBefore, netAfter)
}

func doBinaryObjectWrite(t *testing.T, filename string, net *learning.NeuralNetwork) {
	file, err := os.Create(filename)
	if err != nil {
		t.Error("couldn't create testfile", err)
	}
	defer file.Close()

	err = WriteBinaryObject(file, net)
	if err != nil {
		t.Error("couldn't write binary file", err)
	}
}

func doBinaryObjectRead(t *testing.T, filename string, net *learning.NeuralNetwork) {
	file, err := os.Open(filename)
	if err != nil {
		t.Error("couldn't open testfile", err)
	}
	defer file.Close()

	err = ReadBinaryObject(file, net)
	if err != nil {
		t.Error("couldn't read binary file", err)
	}
}

func doBinaryObjectComparison(t *testing.T, before, after *learning.NeuralNetwork) {
	if before == nil || after == nil {
		t.Error("unexpected nil neural network")
	}

	// Compare every detail of the neural networks
	if before.LayerCount != after.LayerCount {
		t.Error("expecting networks LayerCount to be the same", before.LayerCount, after.LayerCount)
	}

	if before.LearningRate != after.LearningRate {
		t.Error("expecting networks LearningRate to be the same", before.LearningRate, after.LearningRate)
	}

	if before.AttributesSize != after.AttributesSize {
		t.Error("expecting networks AttributesSize to be the same", before.AttributesSize, after.AttributesSize)
	}

	if before.PredictionsSize != after.PredictionsSize {
		t.Error("expecting networks PredictionsSize to be the same", before.PredictionsSize, after.PredictionsSize)
	}

	if (before.ErrorHistory == nil) != (before.ErrorHistory == nil) {
		t.Error("expecting networks ErrorHistory to be the same", before.ErrorHistory, after.ErrorHistory)
	}

	if !mat.Equal(before.ErrorHistory, after.ErrorHistory) {
		t.Error("expecting networks ErrorHistory to be the same", before.ErrorHistory, after.ErrorHistory)
	}

	if len(before.Signals) != len(after.Signals) {
		t.Error("expecting networks Signals length to be the same", len(before.Signals), len(after.Signals))
	}

	if len(before.Outputs) != len(after.Outputs) {
		t.Error("expecting networks Outputs length to be the same", len(before.Outputs), len(after.Outputs))
	}

	if len(before.Weights) != len(after.Weights) {
		t.Error("expecting networks Weights length to be the same", len(before.Weights), len(after.Weights))
	}

	if len(before.Deltas) != len(after.Deltas) {
		t.Error("expecting networks Deltas length to be the same", len(before.Deltas), len(after.Deltas))
	}

	if len(before.Gradients) != len(after.Gradients) {
		t.Error("expecting networks Gradients length to be the same", len(before.Gradients), len(after.Gradients))
	}

	for i := 0; i < len(before.Signals); i++ {
		if (before.Signals[i] == nil) != (after.Signals[i] == nil) {
			t.Error("expecting networks Signals to be the same", "index", i, before.Signals[i], after.Signals[i])
		}
		if !mat.Equal(before.Signals[i], after.Signals[i]) {
			t.Error("expecting networks Signals to be the same", "index", i, before.Signals[i], after.Signals[i])
		}
	}

	for i := 0; i < len(before.Outputs); i++ {
		if (before.Outputs[i] == nil) != (after.Outputs[i] == nil) {
			t.Error("expecting networks Outputs to be the same", "index", i, before.Outputs[i], after.Outputs[i])
		}
		if !mat.Equal(before.Outputs[i], after.Outputs[i]) {
			t.Error("expecting networks Outputs to be the same", "index", i, before.Outputs[i], after.Outputs[i])
		}
	}

	for i := 0; i < len(before.Weights); i++ {
		if (before.Weights[i] == nil) != (after.Weights[i] == nil) {
			t.Error("expecting networks Weights to be the same", "index", i, before.Weights[i], after.Weights[i])
		}
		if !mat.Equal(before.Weights[i], after.Weights[i]) {
			t.Error("expecting networks Weights to be the same", "index", i, before.Weights[i], after.Weights[i])
		}
	}

	for i := 0; i < len(before.Deltas); i++ {
		if (before.Deltas[i] == nil) != (after.Deltas[i] == nil) {
			t.Error("expecting networks Deltas to be the same", "index", i, before.Deltas[i], after.Deltas[i])
		}
		if !mat.Equal(before.Deltas[i], after.Deltas[i]) {
			t.Error("expecting networks Deltas to be the same", "index", i, before.Deltas[i], after.Deltas[i])
		}
	}

	for i := 0; i < len(before.Gradients); i++ {
		if (before.Gradients[i] == nil) != (after.Gradients[i] == nil) {
			t.Error("expecting networks Gradients to be the same", "index", i, before.Gradients[i], after.Gradients[i])
		}
		if !mat.Equal(before.Gradients[i], after.Gradients[i]) {
			t.Error("expecting networks Gradients to be the same", "index", i, before.Gradients[i], after.Gradients[i])
		}
	}
}
