package day_third

import (
	"errors"
	"fmt"

	datareaders "AdventOfCode2022/data_readers"
)

var (
	errRucksackComponentsAreNotEven = errors.New("rucksacks' lengths are not even")
	errElvenGroupsAreNotEven        = errors.New("rucksacks' lengths are not even")
	errNoCommonLetters              = errors.New("no common letters within compared components")
	errInputNotLetter               = errors.New("input rune is not a letter of english alphabet")
)

func SumRucksacksPrioritySupplies(filepath string) (int, error) {
	var result int

	rucksacks, err := getRucksacksWithComponents(filepath)
	if err != nil {
		return 0, fmt.Errorf("get rucksacks from file %w", err)
	}

	for _, rucksack := range rucksacks {
		commonLetters, err := findCommonLetters(rucksack[0], rucksack[1])
		if err != nil {
			return 0, fmt.Errorf("find common letter in components: %w", err)
		}

		partialResult, err := countPartialResult(commonLetters[0])
		if err != nil {
			return 0, fmt.Errorf("count partial result: %w", err)
		}

		result += partialResult
	}

	return result, nil
}

func SumGroupBadgesPriorities(filepath string) (int, error) {
	var result int

	groups, err := getElvenRucksackGroups(filepath)
	if err != nil {
		return 0, fmt.Errorf("get groups from file %w", err)
	}

	for _, group := range groups {
		twoElvesCommonLetters, err := findCommonLetters(group[0], group[1])
		if err != nil {
			return 0, fmt.Errorf("find common letters in groups first two backpacks: %w", err)
		}

		commonLettersForGroup, err := findCommonLetters(string(twoElvesCommonLetters), group[2])
		if err != nil {
			return 0, fmt.Errorf("find common letters in group: %w", err)
		}

		partialResult, err := countPartialResult(commonLettersForGroup[0])
		if err != nil {
			return 0, fmt.Errorf("count partial result: %w", err)
		}

		result += partialResult
	}

	return result, nil
}

func getRucksacksWithComponents(filepath string) ([][2]string, error) {
	var components [][2]string

	rucksacks, err := datareaders.GetDataFromFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("unable to get backpacks from file %w", err)
	}

	for _, rucksack := range rucksacks {
		if len(rucksack)%2 != 0 {
			return nil, errRucksackComponentsAreNotEven
		}

		firstComponent := rucksack[:(len(rucksack) / 2)]
		secComponent := rucksack[(len(rucksack) / 2):]

		components = append(components, [2]string{firstComponent, secComponent})
	}

	return components, nil
}

func findCommonLetters(str1, str2 string) ([]rune, error) {
	var result []rune

	letterMap := make(map[rune]bool)

	for _, ch := range str1 {
		letterMap[ch] = true
	}

	for _, ch := range str2 {
		if _, ok := letterMap[ch]; ok {
			result = append(result, ch)
		}
	}

	if len(result) == 0 {
		return nil, errNoCommonLetters
	}

	return result, nil
}

func countPartialResult(ch rune) (int, error) {
	if ch < 65 || (ch > 90 && ch < 97) || ch > 122 {
		return 0, errInputNotLetter
	}

	var result int
	if ch >= 97 {
		result += int(ch) - 96
	} else {
		result += int(ch) - 64 + 26
	}

	return result, nil
}

func getElvenRucksackGroups(filepath string) ([][3]string, error) {
	var groups [][3]string

	rucksacks, err := datareaders.GetDataFromFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("unable to get rucksacks from file %w", err)
	}

	if len(rucksacks)%3 != 0 {
		return nil, errElvenGroupsAreNotEven
	}

	for idx := 0; idx < len(rucksacks); idx += 3 {
		var group [3]string
		group[0] = rucksacks[idx]
		group[1] = rucksacks[idx+1]
		group[2] = rucksacks[idx+2]

		groups = append(groups, group)
	}

	return groups, nil
}
