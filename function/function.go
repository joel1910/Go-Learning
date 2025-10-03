package main

import (
	"fmt"
	"math"
)

func main() {

	greet()
	add(10, 10)
	fmt.Println()
	res := mul(10, 30)
	fmt.Println("Multiplication of number:", res)

	// multiple return
	rem, qua := remainder(10, 5)
	fmt.Printf("The Remainder is :%f and quotient is :%f\n", rem, qua)

	// multiple arguments
	fmt.Println(sumorconcad(1, 2, 4, 5))
	fmt.Println(sumorconcad("Hi", " ", "Developers"))

	//defer
	defer fmt.Println("Joel")
	fmt.Println("Hello")

	// panic ,recover,defer

	dividetwonum(10, 0)
	dividetwonum(10, 2)
	dividetwonum(10, 6)

}

//function

func greet() {
	fmt.Println("Hi Welcome to Functions")
}

//with argument without return

func add(a int, b int) {
	fmt.Print("Sum of number :", a+b)
}

//with argument with return

func mul(x, y int) int {
	return x * y
}

//with multiple return statement

func remainder(x, y float64) (float64, float64) {
	remaind := math.Mod(x, y)
	quotient := x / y

	return remaind, quotient
}

// variadic function
func sumorconcad[T int | string](vals ...T) T {
	var result T
	for _, v := range vals {
		result += v
	}
	return result
}

// panic ,recover,defer

func dividetwonum(num1, num2 int) {
	defer recoverpanic()
	if num2 == 0 {
		panic("Cannot Divide,Denominator is Zero")
	} else {
		result := num1 / num2
		fmt.Println("The Division of Two Number is", result)
	}
}

func recoverpanic() {
	err := recover()
	if err != nil {
		fmt.Println("Recover from :", err)
	}
}
