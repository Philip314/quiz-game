# Quiz Game

## About

 CLI quiz game that takes in questions and answers from a CSV file in the form of question,answer.

 I created this from following [Gophercises](https://gophercises.com/ "Gophercises") to learn Go.

## Getting Started

1. Clone repo
2. Go into the project folder with command line interface
3. Run project with `go run main.go`

## Command-line Arguments

```
$ go run main.go -h
  -csv string
        CSV file containing the questions and answers (default "problems.csv")
  -time int
        Time limit in seconds (default 20)
```