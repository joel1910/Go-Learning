package main

import "fmt"

func main() {

	x := 20
	p := &x // storing memory address of value x
	fmt.Println("Actual Value of x:", x)
	fmt.Println("Memory Address of x:", p)
	//x = 30  // only inside the function change the value - for outside the function use pointer
	//changeval(x)   // In this function we pass the value but the value not change - for outside the function use pointer
	changevalue(&x) //  here value change because we pass the memory address of x
	fmt.Println("After Change the actual value:", x)
}

// func changeval(value int){
// 	value = 200
// }

func changevalue(value *int) {
	*value = 200
}
