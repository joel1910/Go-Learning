package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("C:/Learning/data.csv")
	if err != nil {
		fmt.Println("Error in Create File")
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	//	defer writer.Flush()

	// why comment means - if want dont want show  data in output terminal. but store in file *
	// bcoz,defer execute end of main func so the data not store in file before end of program *

	writer.Write([]string{"Name", "Email"})
	writer.Write([]string{"Joel", "joel@example.com"})
	writer.Write([]string{"shabana", "shabana@example.com"})

	// writer.Write() ena pannum na write panna file ah temp ah memory la store pannanum *
	// writer.Flush() flush enna pannum na data va file ku push pannum *

	writer.Flush()
	fmt.Println("File Created Successfully")

	data, err := os.Open("C:/Learning/data.csv")
	if err != nil {
		fmt.Println("Error In Open File")
		return
	}

	defer data.Close()

	reader := csv.NewReader(data)

	readdata, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error In Read File")
		return
	}

	fmt.Println("The CSV File Content")

	for _, records := range readdata {
		fmt.Println(records)
	}
}
