package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	csvFile := flag.String("csv", "problems.csv", "a csv file in format of question,answer")
	timeLimit := flag.Int("limit", 30, "a time limit in seconds with 30s as default")
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		log.Fatalf("failed to open the file: %v", *csvFile)
	}

	defer file.Close()
	r := csv.NewReader(file)
	fields, err := r.ReadAll()
	if err != nil {
		log.Fatalf("failed while reading the csv file passed")
	}

	problems := parseFields(fields)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	streak := 0

outer:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s, Your Guess: ", i+1, p.q)
		resultCh := make(chan string)
		go func() {
			var resultInput string
			fmt.Scan(&resultInput)
			resultCh <- resultInput
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nRun out of time --- Scored %d out of %d", streak, len(problems))
			break outer
		case answer := <-resultCh:
			if answer == p.a {
				fmt.Print("RIGHT ANSWER\n")
				streak += 1
			}
		}
	}
}

func parseFields(lines [][]string) []Problem {
	ret := make([]Problem, len(lines))
	for i, lane := range lines {
		ret[i] = Problem{
			q: lane[0],
			a: strings.TrimSpace(lane[1]),
		}
	}
	return ret
}

type Problem struct {
	q string
	a string
}
