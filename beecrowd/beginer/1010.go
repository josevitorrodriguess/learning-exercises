package beginer

import "fmt"

func SimpleCalculate() {
	var codep1, codep2, nump1, nump2 int
	var pricep1, pricep2 float64

	fmt.Scanf("%d", &codep1)
	fmt.Scanf("%d", &nump1)
	fmt.Scanf("%f", &pricep1)
	fmt.Scanf("%d", &codep2)
	fmt.Scanf("%d", &nump2)
	fmt.Scanf("%f", &pricep2)

	total := float64(nump1)*pricep1 + float64(nump2)*pricep2

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", total)
}
