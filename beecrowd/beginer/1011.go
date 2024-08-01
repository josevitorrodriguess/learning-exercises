package beginer

import (
	"fmt"
	"math"
)

func Sphere() {
	var pi = 3.14159
	var r float64

	fmt.Scanf("%f", &r)
	R := math.Pow(r, 3)

	volume := (4.0 / 3) * pi * R

	fmt.Printf("VOLUME = %.3f\n", volume)
}
