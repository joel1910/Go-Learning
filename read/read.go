package main

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

func readCsv() {
	// csv file
	file, err := os.Open("C:/Learning/files/csvsample.csv")
	if err != nil {
		fmt.Println("Error in open File")
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error in Read")
		return
	}

	for i, records := range data {
		if i == 0 {
			continue // remove header
		}
		fmt.Println(records)
	}
}

func readJson() {
	//json

	data, err := os.ReadFile("C:/Learning/files/jsonsample.json")
	if err != nil {
		fmt.Println("Error reading JSON:", err)
		return
	}
	fmt.Println("Data :", string(data))
	var result interface{} // can hold any JSON
	err = json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Println("Raw JSON content:")
	fmt.Printf("%+v\n", result) // print map/array - map[note:map[body:Don't forget our meeting tomorrow! from:Bob heading:Reminder to:Alice]]

	// If it's a map, print keys
	if m, ok := result.(map[string]interface{}); ok {
		for k, v := range m {
			fmt.Println(k, ":", v) // note : map[body:Don't forget our meeting tomorrow! from:Bob heading:Reminder to:Alice]
		}
	}
}

func readxml() {

	file, err := os.Open("C:/Learning/files/xmlsample.xml")
	if err != nil {
		fmt.Println("Error opening XML:", err)
		return
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)

	for {
		tok, _ := decoder.Token()
		if tok == nil {
			break
		}

		switch t := tok.(type) {
		case xml.StartElement:
			fmt.Println("Start Tag:", t.Name.Local)
		case xml.EndElement:
			fmt.Println("End Tag:", t.Name.Local)
		case xml.CharData:
			fmt.Println("Text:", string(t))
		}
	}
}

func main() {
	var num int
	fmt.Println("1.CSV Read File\n2.Json Read File\n3.XML Read File")
	fmt.Scan(&num)

	switch num {
	case 1:
		fmt.Println("CSV File")
		readCsv()
	case 2:
		fmt.Println("JSON File")
		readJson()
	case 3:
		fmt.Println("XML File")
		readxml()
	default:
		fmt.Println("Invalid Input")
	}
}
