package datareaders

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetDataFromFile(fileName string) []string {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := make([]string, 0)
	for scanner.Scan() {
		value := scanner.Text()

		result = append(result, value)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func GetIntDataFromFile(fileName string) []int {
	stringData := GetDataFromFile(fileName)

	result := make([]int, 0)
	for _, val := range stringData {
		value, err := strconv.Atoi(val)
		if err != nil {
			// handle error
			log.Fatal(err)
			os.Exit(2)
		}

		result = append(result, value)
	}

	return result
}

func GetMatrixDataFromFile(fileName string) (result [][]int, rowLength int, colLength int) {
	stringData := GetDataFromFile(fileName)
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

func GetIntDataWithSeparatorFromFile(filename string, s string) []int {
	stringData := GetDataFromFile(filename)

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

	return result
}
