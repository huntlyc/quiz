package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
    "math"
	"log"
	"os"
	"strings"
)

type questionPair struct {
    question, answer string
}

func parseCSVFile(filename string) ([]questionPair, error) {
	var questionPairs []questionPair

	csvFile, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	csvReader := csv.NewReader(strings.NewReader(string(csvFile)))

	for {
		record, err := csvReader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

        questionPairs = append(questionPairs, questionPair{question: record[0], answer: record[1]})
	}

	return questionPairs, nil
}

func main() {
	score := 0

	csvFile := flag.String("f", "problems.csv", "a string")
	flag.Parse()

	if questionPairs, err := parseCSVFile(*csvFile); err != nil {
		log.Fatal(err)
	} else {
		for _, question := range questionPairs {
			fmt.Printf("%s=", question.question)
			var userInput = ""
			if _, err := fmt.Scanln(&userInput); err != nil {
				log.Fatal(err)
			} else {
				if userInput == question.answer {
					score += 1
				}
			}
		}
        numQs := len(questionPairs)
        percent := math.Floor(float64(score)/float64(numQs) * 100)
		fmt.Printf("\n\n Your score was: %d/%d (%.0f%%)\n", score, numQs, percent)
	}
}
