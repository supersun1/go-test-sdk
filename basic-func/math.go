package basic_func

import (
	"fmt"
)


func Addition( num1 int32,  num2 int32) int32 {
	fmt.Println("Adding....")
	tot := num1 + num2
	fmt.Print("Total: ")
	fmt.Println(tot)
	return num1 + num2
}
