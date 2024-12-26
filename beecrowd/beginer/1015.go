package beginer

import (
	"fmt"
	"math"
)

func distanceBetweenTwoPoints() {
	var x1, x2, y1, y2 float64

	fmt.Scanf("%d", &x1)
	fmt.Scanf("%f", &x2)
	fmt.Scanf("%f", &y1)
	fmt.Scanf("%f", &y2)

	distance := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))

	fmt.Printf("%.4f", &distance)
}
