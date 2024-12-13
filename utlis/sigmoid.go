package utlis

import (
	"math"
)

// Sigmoid is the sigmoid activation function.
func Sigmoid(x float64) float64 {
	result := 1 / (1 + math.Exp(-x))
	return result
}

// SigmoidDerivative is the derivative of the sigmoid activation function.
func SigmoidDerivative(x float64) float64 {
	sigmoidVal := Sigmoid(x)
	return sigmoidVal * (1 - sigmoidVal)
}
