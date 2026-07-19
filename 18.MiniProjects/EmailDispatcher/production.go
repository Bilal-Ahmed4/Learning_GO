package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func loadRecipent(filePath string, ch chan Recipent) error {
	defer close(ch)
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	csvReader := csv.NewReader(f)

	result, err := csvReader.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range result[1:] {
		fmt.Println(record)
		ch <- Recipent{
			Name:  record[0],
			Email: record[1],
		}
	}

	return nil
}
