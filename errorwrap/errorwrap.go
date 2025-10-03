package main

import (
	"errors"
	"fmt"
	"os"
)

var errnotavailable = errors.New("tea not available")
var errboil = errors.New("can't boiled water")

var errMsg = errors.New("not found")

//var errmsg = "not found"

func maketea(num int) error {
	switch num {
	case 2:
		return errboil
	case 4:
		return fmt.Errorf("sorry !%w", errnotavailable)
	}
	return nil
}

func findUser() error {
	return fmt.Errorf("look up failed %w", errMsg)
}

func main() {
	//error IS
	err := findUser()

	fmt.Println("err == errmsg ?", err == errMsg)
	fmt.Println(errors.Is(err, errMsg))

	for i := range 5 {

		if err := maketea(i); err != nil {
			if errors.Is(err, errnotavailable) {
				fmt.Println("Sorry Friend Tea Not Available Right Now")
			} else if errors.Is(err, errboil) {
				fmt.Println("Sorry for Inconvinience")
			} else {
				fmt.Println("Unknown Error ", err)
			}
			continue
		}
		fmt.Println("Tea Ready")
	}

	//ErrorAS

	_, err = os.Open("missing-file.txt")

	var errPath *os.PathError

	if errors.As(err, &errPath) {
		fmt.Println("File operation Failed on the path", errPath.Path)
		fmt.Println("op:", errPath.Op)
	} else {
		fmt.Println("No Error on Path")
	}
}
