package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
	"time"
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
    questionsAsked := 0
	correctAnswers := 0

	csvFile := flag.String("f", "problems.csv", "csv file to read")
	duration := flag.Int("d", 30, "time in seconds to run quiz for")
	flag.Parse()

    durationStr := fmt.Sprintf("%ds", *duration)
    timerDuration,err := time.ParseDuration(durationStr);
    if err != nil {
        log.Fatal(err)
    }

	if questionPairs, err := parseCSVFile(*csvFile); err != nil {
		log.Fatal(err)
	} else {
        fmt.Printf("You have %s to answer all questions - press enter to begin", durationStr)
        fmt.Scanln();

        timer := time.NewTimer(timerDuration)
        go func(){ // seperate threaded "goroutine" function that sits and waits for timer channel to fire
            <-timer.C

            fmt.Println("\n\nTime's up!!!")
            outputQuizResults(questionsAsked, correctAnswers)
        }()

        var userInput = ""
        for _, question := range questionPairs {
            questionsAsked++
            fmt.Printf("%s=", question.question)

            if _, err := fmt.Scanln(&userInput); err == nil { // ignore err, blank input
                if userInput == question.answer {
                    correctAnswers += 1
                }
            }
        }

        fmt.Println("\n\nWell done - you answered all questions in the allowed time!!!")
        outputQuizResults(questionsAsked, correctAnswers)
	}
}

func outputQuizResults(numQuestionsAsked, correctAnswers int){
    scoreAsPercentage := math.Floor(float64(correctAnswers)/float64(numQuestionsAsked) * 100)
    fmt.Printf("\n\nYour score was: %d/%d (%.0f%%)\n\n\n", correctAnswers, numQuestionsAsked, scoreAsPercentage)
    os.Exit(1)
}
