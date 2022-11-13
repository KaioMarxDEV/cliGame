package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	csvFile := flag.String("csv", "problems.csv", "a csv file in format of question,answer")
	// TODO: create a time limit for each question
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

	x := parseFields(fields)

	resultInput := ""
	streak := 0

	for i, p := range x {
		fmt.Printf("Problem#%d: %s, Your Guess: ", i+1, p.q)
		fmt.Scan(&resultInput)
		if resultInput == p.a {
			fmt.Print("RIGHT ANSWER\n")
			streak += 1
			continue
		} else {
			fmt.Printf("WRONG ANSWER\nYou LOST! --- SCORE:%d", streak)
			break
		}
	}
}

func parseFields(lines [][]string) []Problem {
	ret := make([]Problem, len(lines))
	for i, lane := range lines {
		ret[i] = Problem{
			q: lane[0],
			a: lane[1],
		}
	}
	return ret
}

type Problem struct {
	q string
	a string
}
