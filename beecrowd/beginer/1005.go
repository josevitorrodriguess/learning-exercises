package beginer
 
import (
    "fmt"
)
 
func Avarage1() {

	var a, b float64

	fmt.Scanf("%f", &a)
	fmt.Scanf("%f", &b)
	
	media := (a*3.5 + b*7.5)/11
	
	fmt.Printf("MEDIA = %.5f\n", media)

}
