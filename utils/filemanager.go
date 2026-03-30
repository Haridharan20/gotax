package utils

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadLines(fileName string) ([]string, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return nil, errors.New("Could not open file!")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		// file.Close()
		return nil, errors.New("Reading the file content failed.")
	}

	// file.Close()
	return lines, nil
}

func WriteJson(path string, data any) error {
	outputDir := "output"

	// Create output directory if it doesn't exist
	err := os.MkdirAll(outputDir, 0755)
	if err != nil {
		return errors.New("Failed to create output directory")
	}

	fullPath := outputDir + "/" + path
	file, err := os.Create(fullPath)

	if err != nil {
		return errors.New("Failed to create a file")
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return errors.New("Failed to convert data to json")
	}

	// file.Close()
	return nil
}
