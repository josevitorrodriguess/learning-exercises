package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	// Function to calculate the greater of two numbers
	maiorAB := (a + b + int(math.Abs(float64(a-b)))) / 2

	// Calculate the greatest of the three numbers
	maiorABC := (maiorAB + c + int(math.Abs(float64(maiorAB-c)))) / 2

	// Print the result with the required message
	fmt.Printf("%d eh o maior\n", maiorABC)
}
