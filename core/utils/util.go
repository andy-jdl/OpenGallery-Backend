package core

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

type InternalServerError struct {
	Code    int
	Message string
}

type NotFoundError struct {
	Code    int
	Message string
}

type InvalidMetadata struct {
	Code    int
	Message string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func (e *InternalServerError) Error() string {
	return fmt.Sprintf("Error: %d: %s", e.Code, e.Message)
}

func (e *InvalidMetadata) Error() string {
	return fmt.Sprintf("Error: %d: invalid metadata type for %s", e.Code, e.Message)
}

func Flatten[T any](lists [][]T) []T {
	var res []T
	for _, list := range lists {
		res = append(res, list...)
	}
	return res
}

func GetCSVDataFromFile(inputFile string) [][]string {
	path, _ := filepath.Abs(inputFile)
	file, _ := os.Open(path)

	defer file.Close()
	csvReader := csv.NewReader(file)
	data, _ := csvReader.ReadAll()

	return data
}
