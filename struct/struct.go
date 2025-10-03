package main

import "fmt"

func main() {

	// basic struct and  map combination

	//struct as a value
	type userdetail struct {
		id         string
		age        int
		designaton string
	}

	// userdetails := map[string]userdetail{
	// 	"Joel":    {"EMP0001", 23, "Backend Developer"},
	// 	"Shabana": {"EMP0002", 22, "Backend Developer"},
	// }

	userdetails := make(map[string]userdetail)

	userdetails["Joel"] = userdetail{"EMP00001", 23, "Backend Developer"}
	userdetails["Shabana"] = userdetail{"EMP00002", 23, "CTP"}

	fmt.Println(userdetails)

	//struct as key - using string

	type name struct {
		firstName string
		lastName  string
	}

	person := make(map[name]string)

	person[name{"Weslin Joel", "raj"}] = "Backend Developer"
	person[name{"Shabana", "Fathima"}] = "CTP"
	person[name{"Ram", "Kumar"}] = "Technical Consultant"

	for persons, value := range person {
		fmt.Println(persons.firstName, persons.lastName, "-", value)
	}

	// using Int

	type position struct {
		x, y int
	}

	points := make(map[position]string)

	points[position{0, 0}] = "Place A"
	points[position{0, 1}] = "Place B"
	points[position{1, 1}] = "Place c"

	for place, value := range points {
		fmt.Println(place.x, ",", place.y, "and the place is ", value)
	}

	// struct basic

	type Employee struct {
		name string
		role string
	}

	employee1 := Employee{"Joel", "BE Developer"}
	employee2 := Employee{"Ram", "Technical Consultant"}

	fmt.Println("Employee 1:", employee1)
	fmt.Println("Employee 2", employee2)

}
