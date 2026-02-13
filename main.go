package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Parse command-line flag for CSV filename
	filename := flag.String("file", "time.csv", "CSV file containing time entries")
	flag.Parse()

	// Open the CSV file
	file, err := os.Open(*filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Create CSV reader
	reader := csv.NewReader(file)

	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	// Sum up all minutes and seconds
	totalMinutes := 0
	totalSeconds := 0

	for i, record := range records {
		if len(record) < 2 {
			fmt.Fprintf(os.Stderr, "Warning: Skipping row %d - insufficient columns\n", i+1)
			continue
		}

		// Parse minutes
		min, err := strconv.Atoi(record[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Skipping row %d - invalid minutes: %v\n", i+1, err)
			continue
		}

		// Parse seconds
		sec, err := strconv.Atoi(record[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Skipping row %d - invalid seconds: %v\n", i+1, err)
			continue
		}

		totalMinutes += min
		totalSeconds += sec
	}

	// Convert total seconds to minutes
	totalMinutes += totalSeconds / 60
	totalSeconds = totalSeconds % 60

	// Convert total minutes to hours
	hours := totalMinutes / 60
	minutes := totalMinutes % 60

	// Output the result
	fmt.Printf("%d:%d %d\n", hours, minutes, totalSeconds)
}
