# quiz-game

## About
 CLI quiz game that takes in questions and answers from a CSV file in the form of question,answer.

 I created this from following [Gophercises](https://gophercises.com/ "Gophercises") to learn Go.

## Getting Started
1. Pull repo
2. Go into the project folder with command line interface
3. Run project with `go run main.go`

Time limit and filename can be set with the command line options `csv` and `time` respectively.
E.g. `go run main.go -time 10 -csv problems.csv` starts the quiz with questions and answers from the file problems.csv with a time limit of 10 seconds.
More details can be found by running `go run main.go -help`.