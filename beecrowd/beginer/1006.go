package beginer

import "fmt"

func Avarage2() {
	var a, b, c float64

	fmt.Scanf("%f", &a)
	fmt.Scanf("%f", &b)
	fmt.Scanf("%f", &c)

	media := ((a * 2) + (b * 3) + (c * 5)) / 10

	fmt.Printf("MEDIA = %.1f\n", media)
}
