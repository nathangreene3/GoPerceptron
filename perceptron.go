package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

// perceptron.go
// Nathan Greene
// Fall 2018

// perceptron is a set of weights and a bias.
type perceptron struct {
	// weights is an ordered set of real values
	weights []float64
	// bias is the default weight applied
	bias float64
}

// Stringer does something in using perceptron.String.
var _ = fmt.Stringer(&perceptron{})

// String returns a formatted string representation of a perceptron. A
// perceptron is represented as
// 	[weights], [bias]: [0.0, ..., 0.0], 0.0.
func (p *perceptron) String() string {
	a := make([]string, 0, len(p.weights))
	for i := range p.weights {
		a = append(a, fmt.Sprintf("%0.2f", p.weights[i]))
	}
	return fmt.Sprintf("[%s], %0.2f", strings.Join(a, ", "), p.bias)
}

// newPerceptron initiates an empty perceptron with a specified number of
// dimensions. All weights and the bias are set to zero.
func newPerceptron(dimensions int) *perceptron {
	// -1 < weights, bias < 1
	p := &perceptron{
		weights: make([]float64, 0, dimensions),
		bias:    1 - 2*rand.Float64(),
	}
	for i := 0; i < dimensions; i++ {
		p.weights = append(p.weights, 1-2*rand.Float64())
	}
	return p
}

// feedForward computes the perceptron decision (result) given an input
// value. A descision function must return a value on the range [0,1].
func (p *perceptron) feedForward(input []float64, descision func(float64) float64) float64 {
	result := p.bias
	for i := range input {
		result += input[i] * p.weights[i]
	}
	return descision(result) // threshold, sigmoid, etc.
}

// backPropagate adjusts the weights by rate x delta given an input.
func (p *perceptron) backPropagate(input []float64, delta, rate float64) {
	p.bias += rate * delta
	for i := range input {
		p.weights[i] += rate * delta * input[i]
	}
}

// learn trains the perceptron given a set of training data (inputs), a
// function accepting training data (trainer), and the learning rate.
func (p *perceptron) learn(descision func(float64) float64, inputs [][]float64, class []float64, rate float64) {
	for i := range inputs {
		p.backPropagate(inputs[i], class[i]-p.feedForward(inputs[i], descision), rate)
	}
}

// verify returns the ratio of the number of correct classifications to
// the total number of inputs to classify.
func (p *perceptron) verify(inputs [][]float64, class []float64, descision func(float64) float64) float64 {
	count := float64(len(inputs)) // Number of inputs
	correct := count              // Number of correct results
	for i := range inputs {
		if p.feedForward(inputs[i], descision) != class[i] {
			correct--
		}
	}
	return correct / count
}

// threshold is a simple decision function alternative to the logistic
// function (sigmoid) or other decision functions. It returns 1 if x is
// positive and 0 otherwise.
func threshold(x float64) float64 {
	if 0 < x {
		return 1
	}
	return 0
}

// sigmoid returns a value on the range (0,1) for any real x.
func sigmoid(x float64) float64 {
	return 1 / (1 + 1/math.Exp(x))
}
