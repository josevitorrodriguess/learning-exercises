package beginer

import "fmt"

const (
	pi = 3.14159
)

func Area() {
	var a, b, c float64
	var areaRec, areaTri, areaTra, areaQua, areaCir float64

	fmt.Scanf("%f", &a)
	fmt.Scanf("%f", &b)
	fmt.Scanf("%f", &c)

	areaTri = (a * c) / 2
	areaCir = pi * c * c
	areaTra = ((a + b) * c) / 2
	areaQua = b * b
	areaRec = a * b

	fmt.Printf("TRIANGULO: %.3f\n", areaTri)
	fmt.Printf("CIRCULO: %.3f\n", areaCir)
	fmt.Printf("TRAPEZIO: %.3f\n", areaTra)
	fmt.Printf("QUADRADO: %.3f\n", areaQua)
	fmt.Printf("RETANGULO: %.3f\n", areaRec)
}
