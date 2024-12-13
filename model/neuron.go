package model

import (
	"math/rand"
)

// neuron represents a neuron in a neural network.
type neuron struct {
	layer  *layer
	inputs []input
}

// newNeuron creates a new neuron with the given layer and previous layer.
func newNeuron(layer *layer, previousLayer *layer) *neuron {

	// Create a new neuron
	n := &neuron{
		layer:  layer,
		inputs: nil,
	}

	// If the previous layer is not nil
	if previousLayer != nil {

		// Create inputs for neurons of the previous layer
		n.inputs = make([]input, 0, len(previousLayer.neurons))
		for _, neuron := range previousLayer.neurons {
			n.inputs = append(n.inputs, newInput(neuron, rand.Float64()-0.5))
		}
	} else {

		// Create an input for the first layer neuron
		n.inputs = []input{newInput(nil, 0)}
	}

	return n
}

// isFirstLayerNeuron returns true if the neuron is the first layer neuron in the network.
func (n *neuron) isFirstLayerNeuron() bool {
	return n.inputs[0].neuron == nil
}

// getInputSum returns the sum of the inputs to the neuron.
func (n *neuron) getInputSum() float64 {

	var sum float64

	// Get the sum of the inputs
	for _, input := range n.inputs {

		// If the input neuron is not nil
		if input.neuron != nil {

			// Add the input value multiplied by the weight
			sum += input.neuron.getValue() * input.weight
		}
	}

	return sum
}

// getValue returns the value of the neuron.
func (n *neuron) getValue() float64 {

	// If the neuron is the first layer neuron
	if n.isFirstLayerNeuron() {

		// Return the input value
		return n.inputs[0].weight

	} else {

		// Return the activation function applied to the sum of the inputs
		return n.layer.network.activationFunction(n.getInputSum())
	}
}

// setInput sets the input value for the neuron.
func (n *neuron) setInput(val float64) {

	// If the neuron is not the first layer neuron, return
	if !n.isFirstLayerNeuron() {
		return
	}

	// Set the input value
	n.inputs[0] = input{
		neuron: nil,
		weight: val,
	}
}

// setError sets the error for the neuron.
func (n *neuron) setError(error float64) {

	// If the neuron is the first layer neuron, return
	if n.isFirstLayerNeuron() {
		return
	}

	// Calculate the error for the neuron
	wDelta := error * n.layer.network.derivativeFunction(n.getInputSum())

	// Loop through the inputs
	for i, input := range n.inputs {

		// If the input neuron is not nil
		if input.neuron != nil {

			// Update the weight of the input
			n.inputs[i].weight -= input.neuron.getValue() * wDelta * n.layer.network.learningRate

			// Update the error of the input neuron
			input.neuron.setError(input.weight * wDelta)
		}
	}
}
