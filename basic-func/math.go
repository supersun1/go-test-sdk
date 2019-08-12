package basic_func

import (
	"fmt"
)


func Addition( num1 int32,  num2 int32)  {
	fmt.Println("Adding....")
	tot := num1 + num2
	fmt.Print("Total: ")
	fmt.Println(tot)
}

func Multiply( num1 int32,  num2 int32) {
	fmt.Println("Multiplying....")
	tot := num1 * num2
	fmt.Print("Total: ")
	fmt.Println(tot)
}
