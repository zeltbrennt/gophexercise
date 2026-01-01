package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
	"time"
)

func readCsvFile(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}

type Problem struct {
	Question string
	Answer   string
}

func parseProblems(records [][]string) ([]Problem, error) {
	problems := make([]Problem, len(records))
	for i, line := range records {
		if len(line) != 2 {
			return nil, fmt.Errorf("illegal format")
		}
		problem := Problem{Question: line[0], Answer: strings.ToLower(line[1])}
		problems[i] = problem
	}
	return problems, nil
}

func startClock(limit int, doneChan chan bool) {
	<-time.NewTimer(time.Duration(limit) * time.Second).C
	doneChan <- false
}

func askQuestions(questions []Problem, correct *int, doneChan chan bool) {
	reader := bufio.NewReader(os.Stdin)
	for i, question := range questions {
		fmt.Printf("Problem #%d: %s = ", i+1, question.Question)
		text, err := reader.ReadString('\n')
		if err != nil {
			panic("error while reading input")
		}
		text = strings.ToLower(strings.TrimSpace(text))
		if text == question.Answer {
			*correct++
		}
	}
	doneChan <- true
}

func newQuestionPermutation(questions []Problem) []Problem {
	perm := rand.Perm(len(questions))
	newQuestionPerm := make([]Problem, len(questions))
	for i, p := range questions {
		newQuestionPerm[perm[i]] = p
	}
	return newQuestionPerm
}

func main() {
	// flags
	problems := flag.String("file", "problems.csv", "a file file with problems")
	limit := flag.Int("limit", 30, "time limit in seconds")
	shuffle := flag.Bool("shuffle", true, "shuffle questions each time")
	flag.Parse()
	// parse csv
	records, err := readCsvFile(*problems)
	if err != nil {
		panic("error reading file")
	}
	questions, err := parseProblems(records)
	if err != nil {
		panic("illegal csv")
	}
	if *shuffle {
		questions = newQuestionPermutation(questions)
	}
	// start the game
	reader := bufio.NewReader(os.Stdin)
	correct := 0
	fmt.Print("press Enter to begin...")
	_, _ = reader.ReadString('\n')
	// parallel timeout and question loop
	doneChan := make(chan bool, 1)
	go startClock(*limit, doneChan)
	go askQuestions(questions, &correct, doneChan)
	// wait for one of two goroutines to finish
	done := <-doneChan
	// end
	fmt.Println()
	if done {
		fmt.Println("Congratulations! All done!")
	} else {
		fmt.Println("Time is up!")
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(questions))
}
