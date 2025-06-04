package main

import "fmt"

func main() {
	p := NewSimplePerceptron()
	for i := range 20 {
		p.Train(0, 1)
		p.Train(1, 0)
		fmt.Printf("%v: %v\n", i, p.String())
	}
	fmt.Printf("%v: %v\n", 0, p.Predict(0))
	fmt.Printf("%v: %v\n", 1, p.Predict(1))
}
