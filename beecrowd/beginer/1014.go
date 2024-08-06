package beginer

import "fmt"

func Consuption() {
	var x int
	var y float64

	fmt.Scanf("%d", &x)
	fmt.Scanf("%f", &y)

	avarage := float64(x) / y

	fmt.Printf("%.3f km/l\n", avarage)
}
