package utlis

// ToBinaryFloatArray converts a number to a binary float array
func ToBinaryFloatArray(num float64) []float64 {

	// Create a float array of 10 elements
	binaryFloat := make([]float64, 10)

	// Set the element at the index of the number to 1
	binaryFloat[int(num)] = 1

	return binaryFloat
}
