package day_first

import (
	"fmt"
	"strconv"

	datareaders "AdventOfCode2022/data_readers"
)

func GetBackpackWithMostCalories(filepath string) (int, error) {
	backpacks, err := getBackpacksFromFile(filepath)
	if err != nil {
		return 0, fmt.Errorf("get backpacks from file %w", err)
	}

	return sumMostCaloriesBackpack(backpacks), nil
}

func GetThreeBackpacksWithMostCalories(filepath string) (int, error) {
	backpacks, err := getBackpacksFromFile(filepath)
	if err != nil {
		return 0, fmt.Errorf("get backpacks from file %w", err)
	}

	return sumThreeMostCaloriesBackpacks(backpacks), nil
}

func sumMostCaloriesBackpack(backpacks [][]int) int {
	var result int
	for _, backpack := range backpacks {
		var currentResult int
		for _, cal := range backpack {
			currentResult += cal
		}

		if currentResult > result {
			result = currentResult
		}
	}

	return result
}

func sumThreeMostCaloriesBackpacks(backpacks [][]int) int {
	var results [3]int
	for _, backpack := range backpacks {
		var currentResult int
		for _, cal := range backpack {
			currentResult += cal
		}

		updateResults(currentResult, &results)
	}

	result := results[0] + results[1] + results[2]

	return result
}

func updateResults(currentResult int, results *[3]int) {
	if currentResult < results[0] {
		return
	}

	results[0] = currentResult
	for i := 1; i < len(results); i++ {
		if currentResult < results[i] {
			break
		}

		results[i-1] = results[i]
		results[i] = currentResult
	}
}

func getBackpacksFromFile(filepath string) ([][]int, error) {
	backpackInfo, err := datareaders.GetDataFromFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("get data from file: %w", err)
	}
	var backpacks [][]int
	var backpack []int
	for _, str := range backpackInfo {
		if str == "" || str == " " {
			backpacks = append(backpacks, backpack)
			backpack = make([]int, 0)
			continue
		}

		calories, err := strconv.Atoi(str)
		if err != nil {
			return nil, fmt.Errorf("converting calories information from string to int: %w", err)
		}

		backpack = append(backpack, calories)
	}

	backpacks = append(backpacks, backpack)

	return backpacks, nil
}
