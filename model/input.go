package model

// input represents an input to a neuron.
type input struct {
	neuron *neuron
	weight float64
}

// newInput creates a new input with the given neuron and weight.
func newInput(neuron *neuron, weight float64) input {
	return input{
		neuron: neuron,
		weight: weight,
	}
}
