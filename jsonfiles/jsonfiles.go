package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	Name, Email string
}

func main() {
	users := []user{
		{"joel", "Joel@example.com"},
		{"shabana", "shabana@example.com"},
	}

	file, err := os.Create("C:/Learning/datajson.json")
	if err != nil {
		fmt.Println("Error in Create File")
		return
	}

	defer file.Close()

	jsondata, _ := json.MarshalIndent(users, "", " ")

	file.Write(jsondata)

	fmt.Println("File Created Successfully")

	data, _ := os.ReadFile("C:/Learning/datajson.json")

	var decoded []user

	_ = json.Unmarshal(data, &decoded)

	fmt.Println("Json Content")

	for _, records := range decoded {
		fmt.Println("Name :", records.Name, "Email:", records.Email)
	}

}
