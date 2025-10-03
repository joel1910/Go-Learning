package main

import "fmt"

type animals interface {
	speak() string
}

type dog struct{}
type cat struct{}

func (dog) speak() string {
	return "Wow Wow!"
}

func (cat) speak() string {
	return "meow meow"
}

func main() {
	var temp interface{}
	var catVar cat
	var dogVar dog

	temp = catVar
	fmt.Println(temp.(cat).speak())

	temp = dogVar
	fmt.Println(temp.(dog).speak())

	// animal := []animals{dog{}, cat{}}

	// for _, animalsname := range animal {
	// 	fmt.Println(animalsname.speak())
	// }
}
