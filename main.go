package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFilename := getFilename()

	file, err := os.Open(csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Error opening file: %s", csvFilename))
	}

	csvReader := csv.NewReader(file)

	records, err := csvReader.ReadAll()
	if err != nil {
		exit("Error reading records")
	}

	problems := createProblems(records)

	numCorrect := 0

	for _, v := range problems {
		fmt.Println(v.Question)
		userAnswer := getUserInput()
		if userAnswer == v.Answer {
			numCorrect++
		}
	}

	fmt.Printf("You've scored %d out of %d", numCorrect, len(problems))
}

type Problem struct {
	Question string
	Answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
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
	userInput, err := inputReader.ReadString('\n')
	if err != nil {
		exit("Error reading input")
	}
	return strings.Trim(userInput, "\r\n")
}
