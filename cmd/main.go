package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	//open file
	file, err := os.Open("../problems.csv")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	// read file

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2

	for {
		record, e := reader.Read()
		if e != nil {
			panic(e)
		}
		// return []string len 2, cap 2
		fmt.Println(record)
	}

	//quiz

	// var question string

	// question = fmt.Sprintf("What")
}
