package datareaders

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetDataFromFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			err = fmt.Errorf("could not close file gently: %w", err)
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	result := make([]string, 0)
	for scanner.Scan() {
		value := scanner.Text()

		result = append(result, value)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("run scanner: %w", err)
	}

	return result, nil
}

func GetIntDataFromFile(fileName string) ([]int, error) {
	stringData, err := GetDataFromFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("get data from file: %w", err)
	}

	result := make([]int, 0)
	for _, val := range stringData {
		value, err := strconv.Atoi(val)
		if err != nil {
			return nil, fmt.Errorf("cast value to int: %w", err)
		}

		result = append(result, value)
	}

	return result, nil
}

func GetMatrixDataFromFile(fileName string) (result [][]int, rowLength int, colLength int, err error) {
	stringData, err := GetDataFromFile(fileName)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("get data from file: %w", err)
	}

	rowLength = len(stringData)
	colLength = len(stringData[0])

	result = make([][]int, rowLength)
	for i := range result {
		result[i] = make([]int, colLength)
		for j := range result[i] {
			singleVal, err := strconv.Atoi(string(stringData[i][j]))
			if err != nil {
				log.Fatal(err)
				os.Exit(2)
			}

			result[i][j] = singleVal
		}
	}

	return
}

func GetIntDataWithSeparatorFromFile(filename string, s string) ([]int, error) {
	stringData, err := GetDataFromFile(filename)
	if err != nil {
		return nil, fmt.Errorf("get data from file: %w", err)
	}

	result := make([]int, 0)
	for _, val := range strings.Split(stringData[0], s) {
		value, err := strconv.Atoi(val)
		if err != nil {
			// handle error
			log.Fatal(err)
			os.Exit(2)
		}

		result = append(result, value)
	}

	return result, nil
}
