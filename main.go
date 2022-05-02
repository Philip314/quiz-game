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

	records, err := reader.ReadAll()
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(createProblems(records))
}

type Problem struct {
	Question string
	Answer   string
}

func getFilename() string {
	csvFilename := flag.String("csv", "problems.csv", "CSV file containing the questions and answers")
	flag.Parse()
	return *csvFilename
}

func createProblems(records [][]string) []Problem {
	problems := make([]Problem, len(records))
	for i, row := range records {
		problems[i] = Problem{
			Question: row[0],
			Answer:   row[1],
		}
	}
	return problems
}
