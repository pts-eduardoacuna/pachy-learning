package learning

// EncodeAttributes applies a hard threshold on the [0,255] range of values of an array such that
// any number other than 0 is transformed into a 1.
func EncodeAttributes(xs []float64) []float64 {
	encoding := make([]float64, len(xs))
	for i, x := range xs {
		if x > 0 {
			encoding[i] = 1.0
		} else {
			encoding[i] = 0.0
		}
	}
	return encoding
}

// EncodeTarget takes a number between [0,9] and computes an array of zeros with a 1 in the position
// of the input.
func EncodeTarget(target float64) []float64 {
	index := int(target)
	encoding := make([]float64, 10)
	encoding[index] = 1.0

	return encoding
}

// DecodeTargets takes an array of numbers and returns the index of the maximum value.
func DecodeTargets(targets []float64) float64 {
	return float64(ArgMax(targets))
}

// ArgMax takes an array of numbers and returns the index of the maximum value.
func ArgMax(xs []float64) int {
	idx := 0
	max := 0.0
	for i, x := range xs {
		if x > max {
			max = x
			idx = i
		}
	}
	return idx
}
