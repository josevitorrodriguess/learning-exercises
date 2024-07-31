package main

import "fmt"

func main() {
	var n, h int
	var sal float64

	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &h)
	fmt.Scanf("%f", &sal)

	salary := float64(h) * sal

	fmt.Printf("NUMBER = %d\nSALARY = U$%.2f\n", n, salary)
}
