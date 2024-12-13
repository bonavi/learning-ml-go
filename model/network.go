package model

import (
	"log"

	"neuralNetwork/utlis"
)

// Network represents a neural network.
type Network struct {
	learningRate       float64
	layers             []*layer
	activationFunction func(float64) float64
	derivativeFunction func(float64) float64
}

// NewNetwork creates a new neural network with the given layer sizes, learning rate, and activation functions.
func NewNetwork(
	layerSizes []int,
	learningRate float64,
	activationFunction func(float64) float64,
	derivativeFunction func(float64) float64,
) *Network {

	// Initialize the network
	network := &Network{
		learningRate:       learningRate,
		layers:             make([]*layer, len(layerSizes)),
		activationFunction: activationFunction,
		derivativeFunction: derivativeFunction,
	}

	// Create layers
	var previousLayer *layer = nil
	for layerIndex, layerSize := range layerSizes {
		network.layers[layerIndex] = newLayer(layerSize, previousLayer, network)
		previousLayer = network.layers[layerIndex]
	}

	return network
}

// SetInput sets the input values for the network.
func (n *Network) SetInput(val []float64) {
	n.layers[0].setInput(val)
}

// GetPredictions returns the output values of the network.
func (n *Network) GetPredictions() []float64 {

	// Get the last layer
	lastLayer := n.layers[len(n.layers)-1]

	// Create a slice to store the predictions
	prediction := make([]float64, len(lastLayer.neurons))

	// Get the output values
	for i, neuron := range lastLayer.neurons {
		prediction[i] += neuron.getValue()
	}

	return prediction
}

// trainOnce trains the network once on the given inputs and expected outputs.
func (n *Network) trainOnce(inputs [][]float64, labels []int) {

	successes := 0

	// Loop through the inputs
	for i := 0; i < len(inputs); i++ {

		// Get the input and expected output for this iteration
		input := inputs[i]
		label := labels[i]

		// Set the input values
		n.SetInput(input)

		// Get the output values
		predictions := n.GetPredictions()

		// Get the best result for accuracy calculation
		bestResult := utlis.GetBestResult(predictions)

		// If the best result is equal to the best true value, increment the successes
		if bestResult == labels[i] && bestResult != -1 {
			successes++
		}

		// For every 1000th iteration
		if i%100 == 0 && i != 0 {

			// Print the accuracy
			log.Printf("Testcase: %d. Accuracy: %.0f%%\n", i, float64(successes))

			// Reset the successes
			successes = 0
		}

		// Loop through the predictions
		for j := 0; j < len(predictions); j++ {

			// Calculate the error rate
			errorRate := predictions[j] - utlis.ToBinaryFloatArray(float64(label))[j]

			// Set the error rate for each layer
			n.layers[len(n.layers)-1].neurons[j].setError(errorRate)
		}
	}
}

// Train trains the network on the given inputs and expected outputs for the given number of epochs.
func (n *Network) Train(input [][]float64, labels []int, epochs int) {

	// Train the network for the given number of epochs
	for i := 0; i < epochs; i++ {
		log.Printf("Epoch: %d\n", i+1)
		n.trainOnce(input, labels)
	}
}
