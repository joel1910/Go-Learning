package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	//write
	textContent := "Hello!This is Joel\nBackend Developer\n"

	err := os.WriteFile("test.txt", []byte(textContent), 0644)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("File Created Sucessfully")
	}

	//read

	readfileData := "C:/Learning/report.txt"
	data, err := os.ReadFile(readfileData)

	if err != nil {
		fmt.Println("Unable to Read File:", err)
		return
	} else {
		fmt.Println("Data :", string(data))
	}

	// first - find src and destination -> then copy destination from source

	//source
	src, err := os.Open("C:/Learning/report.txt")

	if err != nil {
		fmt.Println("Error in Open", err)
		return
	}

	defer src.Close()

	//destination
	dest, err := os.Create("C:/Learning/Rough/copy.txt")

	if err != nil {
		fmt.Println("Error in Create", err)
	}

	defer dest.Close()

	//copy

	_, err = io.Copy(dest, src)
	if err != nil {
		fmt.Println("Error in Copy :", err)
		return
	}
	fmt.Println("File Copied Successfully!")

	//delete

	err = os.Remove("C:/Learning/copyreport.txt")
	if err != nil {
		fmt.Println("Error In delete :", err)
		return
	}
	fmt.Println("File Deleted Successfully")

}
