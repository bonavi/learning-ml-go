package utlis

// GetBestResult returns the index of the highest value in the predictions array.
func GetBestResult(predictions []float64) int {

	// Set the initial best result and value
	bestResult := 0
	bestValue := 0.0

	// Loop through the predictions
	for i, prediction := range predictions {

		// If the prediction is greater than the best value
		if prediction > bestValue {

			// Set the best result and value
			bestResult = i
			bestValue = prediction
		}
	}

	//// If the best value is less than 0.5
	//if bestValue < 0.5 {
	//
	//	// It means the prediction is not confident
	//	return -1
	//}

	// Return the best result
	return bestResult
}
