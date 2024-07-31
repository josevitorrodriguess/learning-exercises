package beginer

import "fmt"

func Salary() {
	var n, h int
	var sal float64

	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &h)
	fmt.Scanf("%f", &sal)

	salary := float64(h) * sal

	fmt.Printf("NUMBER = %d\n", n)
	fmt.Printf("SALARY = U$ %.2f\n", salary)
}
