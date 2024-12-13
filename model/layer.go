package model

// layer represents a layer of neurons in a neural network.
type layer struct {
	network *Network
	neurons []*neuron
}

// newLayer creates a new layer with the given number of neurons, previous layer, and network.
func newLayer(neuronsCount int, previousLayer *layer, network *Network) *layer {

	// Create a new layer
	layer := &layer{
		network: network,
		neurons: make([]*neuron, 0, neuronsCount),
	}

	// Create neurons
	for i := 0; i < neuronsCount; i++ {
		layer.neurons = append(layer.neurons, newNeuron(layer, previousLayer))
	}

	return layer
}

// isFirstLayer returns true if the layer is the first layer in the network.
func (l *layer) isFirstLayer() bool {
	return l.neurons[0].isFirstLayerNeuron()
}

// setInput sets the input values for the layer.
func (l *layer) setInput(vals []float64) {

	// Check if the layer is the first layer
	if !l.isFirstLayer() {
		return
	}

	// Check if the number of values is equal to the number of neurons
	if len(vals) != len(l.neurons) {
		return
	}

	// Set the input values
	for i, val := range vals {
		l.neurons[i].setInput(val)
	}

}
