package main

import "fmt"

func main() {

	// if ,else if ladder

	var hour int
	fmt.Print("Enter the hour, (i.e - 15) :")
	fmt.Scan(&hour)

	if hour < 12 {
		fmt.Println("Good Morning !")
	} else if hour >= 12 && hour < 16 {
		fmt.Println("Good Afternoon !")
	} else if hour >= 16 && hour < 19 {
		fmt.Println("Good Evening !")
	} else if hour >= 19 && hour <= 24 {
		fmt.Println("Good Night !")
	} else {
		fmt.Println("Invalid Hours - Upto 24 Hours")
	}

	//switch Arithematic operator

	var num1, num2 float64
	var operator string

	fmt.Print("Enter first Number :")
	fmt.Scan(&num1)
	fmt.Print("Enter Second Number :")
	fmt.Scan(&num2)
	fmt.Print("Enter the Operator- (+,-,*,/) :")
	fmt.Scan(&operator)

	switch operator {

	case "+":
		fmt.Println("Sum of Two Number is :", num1+num2)
	case "-":
		fmt.Println("Substraction of Two Number is :", num1-num2)
	case "/":
		fmt.Println("Division of Two Number is :", num1/num2)
	case "*":
		fmt.Println("Multiplication of Two Number is :", num1*num2)
	default:
		fmt.Println("Invalid Operator")
	}

	// for loop

	for i := 1; i <= 5; i++ {
		fmt.Print(i)
	}
	fmt.Println()

	// while-type for loop

	sum := 0
	var num int

	fmt.Print("Enter the Number for Sum of Digits :")
	fmt.Scan(&num)

	for num > 0 {
		digit := num % 10
		sum = sum + digit
		num = num / 10
	}
	fmt.Println(sum)

	// infinite Loop

	num3 := 0
	for {
		num3++
		if num3 == 8 {
			continue
		}
		if num3 > 15 {
			break
		}
		fmt.Print(num3)
	}

}
