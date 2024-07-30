package beginer
 
import (
    "fmt"
)
 
func ex1001() {

 var a, b int

	fmt.Println("Enter two numsber separated for this form(num1,num2)")
	fmt.Scanf("%d,%d",&a,&b)
	X := a+b 
	fmt.Printf("\nx = %d\n", X)
}
