package main

import (
	"bufio"
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

	csvReader := csv.NewReader(file)

	records, err := csvReader.ReadAll()
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

func getUserInput() string {
	inputReader := bufio.NewReader(os.Stdin)
	answer, err := inputReader.ReadString('\n')
	if err != nil {
		os.Exit(1)
	}
	return answer
}
