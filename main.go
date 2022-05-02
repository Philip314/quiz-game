package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := getFilename()

	file, err := os.Open(csvFilename)
	if err != nil {
		os.Exit(1)
	}

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(lines)
}

func getFilename() string {
	csvFilename := flag.String("csv", "problems.csv", "CSV file containing the questions and answers")
	flag.Parse()
	return *csvFilename
}
