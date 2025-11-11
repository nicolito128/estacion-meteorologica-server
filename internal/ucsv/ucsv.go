/*
Package ucsv provides utilities for handling CSV files.
*/
package ucsv

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadAll(filepath string) ([][]string, error) {
	file, err := os.OpenFile(filepath, os.O_RDONLY|os.O_CREATE, 0o644)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	return csv.NewReader(file).ReadAll()
}

func WriteLine(filepath string, record []string) error {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)

	err = writer.Write(record)
	if err != nil {
		return fmt.Errorf("error writing record: %w", err)
	}

	writer.Flush()
	return writer.Error()
}

func WriteLines(filepath string, records [][]string) error {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)

	err = writer.WriteAll(records)
	if err != nil {
		return fmt.Errorf("error writing records: %w", err)
	}

	writer.Flush()
	return writer.Error()
}
