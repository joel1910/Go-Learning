package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type user struct {
	Name, Email string
}

func main() {

	users := []user{
		{Name: "test", Email: "test@example"},
		{Name: "test1", Email: "test1@example"},
	}

	file, _ := os.Create("C:/Learning/dataxml.xml")
	defer file.Close()

	xmldata, _ := xml.MarshalIndent(users, "", " ")
	file.Write([]byte(xml.Header))
	file.Write(xmldata)

	fmt.Println("XML is Created")

	data, _ := os.ReadFile("C:/Learning/dataxml.xml")

	var decoded []user

	_ = xml.Unmarshal(data, &decoded)

	for _, records := range decoded {
		fmt.Println(records.Name, records.Email)
	}

}
