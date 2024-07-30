package main

import (
	"fmt"
)

func ex1002() {

	n := 3.14159

	var r float64
	fmt.Scanf("%f",&r)

	A := n * (r*r)

	fmt.Printf("A = %.4f", A)
	
}
