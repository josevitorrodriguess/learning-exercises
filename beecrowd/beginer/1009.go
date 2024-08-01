package beginer

import "fmt"

func SalaryWithBonus() {
	var name string
	var salary, valueSold float64

	fmt.Scanf("%s", &name)
	fmt.Scanf("%f", &salary)
	fmt.Scanf("%f", &valueSold)

	finalSalary := salary + valueSold*0.15

	fmt.Printf("TOTAL = R$ %.2f\n", finalSalary)
}
