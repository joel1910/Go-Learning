package main

//basic

import (
	"fmt"
	"strconv"
)

func main() {

	// Data Types

	var age int = 23
	var Employee string = "Joel"
	var salary float64 = 50000.58
	var isMarried bool = false

	fmt.Println(age, Employee, salary, isMarried)

	// Variables
	var c int = 20
	d := 40
	name := "joel"

	fmt.Println(c, d, name)

	//constant

	const pi = 3.14
	fmt.Println("Pi :", pi)

	// operators

	a, b := 5, 10
	fmt.Println("Sum is :", a+b)
	fmt.Println("Difference is :", a-b)
	fmt.Println("Multiplication is :", a*b)
	fmt.Println("Division is :", a/b)
	fmt.Println("Mod is :", a%b)
	fmt.Println("Compare is :", a > b)

	// type conversion

	var x int = 10
	var y float64 = float64(x) // convert int → float
	var z int = int(y)         // convert float → int

	fmt.Println("Int:", x, "Float:", y, "Back to Int:", z)

	//str -> Int

	str := "123"
	num, _ := strconv.Atoi(str) // return two values,so here _ used to Ignore Error,
	fmt.Println("String to Int:", num)

	// Int to string

	number := 456
	strng := strconv.Itoa(number)
	fmt.Println("Int to String:", strng)

	//ASCII Value

	ch := 'A'
	fmt.Println("ASCII Value of ch: ", ch)
	fmt.Println("String of ch:", string(ch))

	//function syntax

	result := add(10, 30)
	fmt.Println("Sum is : ", result)

}

func add(a int, b int) int {
	return a + b
}
