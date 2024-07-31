package beginer

import "fmt"

func SimpleProduct() {
	var a, b int

	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)

	prod := a * b

	fmt.Printf("PROD = %d\n", prod)

}
