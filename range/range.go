package main

import "fmt"

func main() {

	// Range in array
	num := []int{10, 20, 30, 40}

	for Index, Value := range num {
		fmt.Printf("Index : %d Value : %d\n", Index, Value)
	}

	// Range in Map

	mymap := map[string]int{"A": 1, "B": 2, "C": 3}

	for key, pair := range mymap {
		fmt.Printf("key :%s pair:%d\n", key, pair)
	}

	// range in string

	str := "GO"

	for Index1, char := range str {
		fmt.Printf("The Index is %d The Char is %c\n", Index1, char)
	}

}
