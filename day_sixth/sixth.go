package day_sixth

import (
	"errors"
	"fmt"

	datareaders "AdventOfCode2022/data_readers"
)

var (
	errIncorrectInput = errors.New("input should span only one line")
	errSignalTooShort = errors.New("signal must be at least 4 signs long")
	errNoMarkerFound  = errors.New("no marker found in whole input")
)

func FindFirstStartOfPacketMarker(filename string) (int, error) {
	const markerLength = 4

	result, err := findFirstStartOfMarker(filename, markerLength)
	if err != nil {
		return 0, fmt.Errorf("find first start of packet marker: %w", err)
	}

	return result, nil
}

func FindFirstStartOfMessageMarker(filename string) (int, error) {
	const markerLength = 14

	result, err := findFirstStartOfMarker(filename, markerLength)
	if err != nil {
		return 0, fmt.Errorf("find first start of message marker: %w", err)
	}

	return result, nil
}

func findFirstStartOfMarker(filename string, markerLength int) (int, error) {
	var result int

	dataInput, err := datareaders.GetDataFromFile(filename)
	if err != nil {
		return 0, fmt.Errorf("get data from file: %w", err)
	}

	if len(dataInput) > 1 {
		return 0, errIncorrectInput
	}

	inputSignal := dataInput[0]

	result, err = findMarkerPosition(inputSignal, markerLength)
	if err != nil {
		return 0, fmt.Errorf("find market index: %w", err)
	}

	return result, nil
}

func findMarkerPosition(inputSignal string, markerLength int) (int, error) {
	if len(inputSignal) < markerLength {
		return 0, errSignalTooShort
	}

	signAmountMap := make(map[uint8]int)
	for idx := 0; idx < markerLength; idx++ {
		if _, found := signAmountMap[inputSignal[idx]]; !found {
			signAmountMap[inputSignal[idx]] = 1
		} else {
			signAmountMap[inputSignal[idx]]++
		}
	}

	for idx := markerLength; idx < len(inputSignal); idx++ {
		if len(signAmountMap) == markerLength {
			return idx, nil
		}

		signAmountMap[inputSignal[idx-markerLength]]--
		if value := signAmountMap[inputSignal[idx-markerLength]]; value == 0 {
			delete(signAmountMap, inputSignal[idx-markerLength])
		}

		if _, found := signAmountMap[inputSignal[idx]]; !found {
			signAmountMap[inputSignal[idx]] = 1
		} else {
			signAmountMap[inputSignal[idx]]++
		}
	}

	if len(signAmountMap) != markerLength {
		return 0, errNoMarkerFound
	}

	return len(inputSignal), nil
}
