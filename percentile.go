package percentile

// Calculate returns the percentile value of the given numbers.
func Calculate(numbers []float64, percent int) float64 {
	if percent == 0 {
		return numbers[0]
	}

	l := len(numbers)
	if percent == 100 {
		return numbers[l-1]
	}

	index := float64(l*percent) / 100
	return (numbers[int(index)] + numbers[int(index)-1]) / 2
}
