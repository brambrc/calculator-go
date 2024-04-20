package main

import (
	"fmt"
)

type Operation interface {
	Compute(a, b float64) float64
}

type Addition struct{}

func (Addition) Compute(a, b float64) float64 { return a + b }

// fill the code here

type OperationFactory struct{}

func (OperationFactory) CreateOperation(operator string) Operation {
	switch operator {
	case "+":
		return Addition{}
	// fill the code here
	default:
		return nil
	}

}

func main() {
	var a, b float64 = 10, 5
	factory := OperationFactory{}

	// Pertambahan
	operation := factory.CreateOperation("+")
	fmt.Println("Addition: ", operation.Compute(a, b)) // Output: 15

	// Pengurangan

	// Perkalian

	// Pembagian

}
