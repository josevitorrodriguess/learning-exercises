package beginer

import (
	"fmt"
	"math"
)

func TheGreatest() {
	var a, b, c int

	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)

	maiorAB := (a + b + int(math.Abs(float64(a-b)))) / 2

	maiorABC := (maiorAB + c + int(math.Abs(float64(maiorAB-c)))) / 2

	fmt.Printf("%d eh o maior\n", maiorABC)
}
