package main

import "fmt"

func main() {

	//create map
	// initial map empid := make(map[string]string)

	empid := map[string]string{
		"Joel":    "EMP000001",
		"Shabana": "EMP000002",
	}

	//fetch data in map using key
	fmt.Println("Employee ID of Joel :", empid["Joel"])

	//insert key and value
	empid["Ram"] = "EMP000003"
	fmt.Println("The Map is :", empid)

	//delete key and value
	delete(empid, "Joel")
	fmt.Println("After Delete Joel :", empid)

	//check the key already exists

	emp, exists := empid["Ram"] // exists return true/false ,emp return - value
	if exists {
		fmt.Println("Employee Already Exists", emp, exists)
	} else {
		fmt.Println("Employee not found")
		empid["Ram"] = "EMP000003"
	}

	emp1, exists1 := empid["Karthi"]
	if exists1 {
		fmt.Println("Employee Already Exists", emp1, exists1)
	} else {
		fmt.Println("Employee not found")
		empid["Karthi"] = "EMP000004"
	}

	fmt.Println("Full Map :", empid)

	//Iteration over map

	for key, value := range empid {
		fmt.Printf("%s : %s\n", key, value)
	}

	//word count

	language := []string{"Go", "Python", "Java", "Go", "Java", "Go"}

	count := make(map[string]int)

	for _, v := range language {
		count[v]++
	}
	fmt.Println(count)

}
