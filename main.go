package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	Narrative1 = 01
	Debit      = 8
	Currency   = 9
)

// this method will addup all the currencies
func addUp(reader *csv.Reader) map[string]float64 {
	result := make(map[string]float64)

	for {
		// read ever record
		records, err := reader.Read()

		// breakes the loop once find the end of line
		if err == io.EOF {
			break
		}

		// compare if it has the PAY keyword which is the requirment
		if strings.Contains(records[Narrative1], "PAY") {
			// convert the debit amount to float to bitsize of 64

			// check if debit length is greater than zero means debit is there
			if len(records[Debit]) > 0 {
				debit, _ := strconv.ParseFloat(records[Debit], 64)

				val, ok := result[records[Currency]]

				if !ok {
					result[records[Currency]] = debit
				} else {
					result[records[Currency]] = val + debit
				}
			}
		}

	}

	return result
}

func main() {
	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.File
	file, err := os.Open("statement.csv")

	// Checks for the error
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	// Closes the file
	defer file.Close()

	// The csv.NewReader() function is called in
	// which the object os.File passed as its parameter
	// and this creates a new csv.Reader that reads
	// from the file
	reader := csv.NewReader(file)

	// ReadAll reads all the records from the CSV file
	// and Returns them as slice of slices of string
	// and an error if any

	// Loop to iterate through
	// and print each of the string slice

	// this will addup all the currencies and return the result
	result := addUp(reader)
	fmt.Println("Total: ")
	for val, key := range result {
		fmt.Println(val, key)
	}

}
