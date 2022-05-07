package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "CSV file containing the questions and answers")
	timeLimit := flag.Int("time", 20, "Time limit in seconds")

	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Error opening file: %s", *csvFilename))
	}

	csvReader := csv.NewReader(file)

	records, err := csvReader.ReadAll()
	if err != nil {
		exit("Error reading records")
	}

	problems := createProblems(records)
	numCorrect := 0
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

questionLoop:
	for _, v := range problems {
		fmt.Println(v.Question)
		inputChan := make(chan string)
		go getUserInput(inputChan)

		select {
		case <-timer.C:
			break questionLoop
		case userAnswer := <-inputChan:
			if userAnswer == v.Answer {
				numCorrect++
			}
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

func getUserInput(inputChan chan string) {
	inputReader := bufio.NewReader(os.Stdin)
	userInput, err := inputReader.ReadString('\n')
	if err != nil {
		exit("Error reading input")
	}
	inputChan <- strings.Trim(userInput, "\r\n")
}
