package beginer

import "fmt"

func Diference() {

	var a, b, c, d int

	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)
	fmt.Scanf("%d", &d)

	diff := (a*b - c*d)

	fmt.Printf("DIFERENCA = %d\n", diff)
}
