package main

import "fmt"

type SimplePerceptron struct {
	Weight       float64
	Bias         float64
	LearningRate float64
}

func NewSimplePerceptron() *SimplePerceptron {
	return &SimplePerceptron{
		Weight:       0.0,
		Bias:         0.0,
		LearningRate: 0.1,
	}
}

func Activate(result float64) int {
	if result >= 1.0 {
		return 1
	}
	return 0
}

func (p *SimplePerceptron) Predict(input int) int {
	result := p.Weight*float64(input) + p.Bias
	return Activate(result)
}

func (p *SimplePerceptron) Train(input int, expected int) {
	err := expected - p.Predict(input)
	p.Weight += float64(err) * float64(input) * p.LearningRate
	p.Bias += float64(err) * p.LearningRate
}

func (p *SimplePerceptron) String() string {
	result := fmt.Sprintf("Weight: %v ", p.Weight)
	result += fmt.Sprintf("Bias: %v ", p.Bias)
	return result
}
