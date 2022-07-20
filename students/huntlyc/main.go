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

func parseCSVFile(filename string) ([][]string, error) {
	var records [][]string

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

		records = append(records, record)
	}

	return records, nil
}

func main() {
	score := 0

	csvFile := flag.String("f", "problems.csv", "a string")
	flag.Parse()

	if records, err := parseCSVFile(*csvFile); err != nil {
		log.Fatal(err)
	} else {
		for _, row := range records {
			fmt.Printf("%s=", row[0])
			var userInput = ""
			if _, err := fmt.Scanln(&userInput); err != nil {
				log.Fatal(err)
			} else {
				if userInput == row[1] {
					score += 1
				}
			}
		}
        numQs := len(records)
        percent := math.Floor(float64(score)/float64(numQs) * 100)
		fmt.Printf("\n\n Your score was: %d/%d (%.0f%%)\n", score, numQs, percent)
	}
}
