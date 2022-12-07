package day_fifth

import (
	"AdventOfCode2022/common"
	datareaders "AdventOfCode2022/data_readers"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var noCratePositionInfo = errors.New("unable to get crate position info")
var noCraneMovesInfo = errors.New("unable to get crane moves info")

func GetTopCratesNamesForCrateMover9000(filepath string) (string, error) {
	var result string

	crateStacks, err := getCrates(filepath)
	if err != nil {
		return "", fmt.Errorf("get create stacks: %w", err)
	}

	craneMoves, err := getCraneMoves(filepath)
	if err != nil {
		return "", fmt.Errorf("get creane moves: %w", err)
	}

	err = moveCratesForCrateMover9000(crateStacks, craneMoves)
	if err != nil {
		return "", fmt.Errorf("move crates: %w", err)
	}

	result = readTopCrateFromEachStack(crateStacks)

	return result, nil
}

func GetTopCratesNamesForCrateMover9001(filepath string) (string, error) {
	var result string

	crateStacks, err := getCrates(filepath)
	if err != nil {
		return "", fmt.Errorf("get create stacks: %w", err)
	}

	craneMoves, err := getCraneMoves(filepath)
	if err != nil {
		return "", fmt.Errorf("get creane moves: %w", err)
	}

	err = moveCratesForCrateMover9001(crateStacks, craneMoves)
	if err != nil {
		return "", fmt.Errorf("move crates: %w", err)
	}

	result = readTopCrateFromEachStack(crateStacks)

	return result, nil
}

func getCrates(filepath string) ([]common.Stack[string], error) {
	var result []common.Stack[string]

	inputInfo, err := datareaders.GetDataFromFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("get data from file: %w", err)
	}

	cratesPositionLineIdx, err := getCratePositionsLineIdx(inputInfo)
	if err != nil {
		return result, fmt.Errorf("get cranes last index: %w", err)
	}

	result = getCrateStacks(inputInfo, cratesPositionLineIdx)

	return result, nil
}

func getCraneMoves(filepath string) ([][3]int, error) {
	var result [][3]int
	inputInfo, err := datareaders.GetDataFromFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("get data from file: %w", err)
	}

	craneMovesFirstLineIdx, err := getCraneMovesLineIdx(inputInfo)
	if err != nil {
		return nil, fmt.Errorf("get cranes last index: %w", err)
	}

	result = getCraneMovesArray(inputInfo, craneMovesFirstLineIdx)
	if err != nil {
		return nil, fmt.Errorf("fill array of subsequent crane moves: %w", err)
	}

	return result, nil
}

func getCratePositionsLineIdx(infoRows []string) (int, error) {
	for idx, row := range infoRows {
		if row == "" {
			return idx - 1, nil
		}
	}

	return 0, noCratePositionInfo
}

func getCraneMovesLineIdx(infoRows []string) (int, error) {
	for idx, row := range infoRows {
		if row == "" {
			return idx + 1, nil
		}
	}

	return 0, noCraneMovesInfo
}

func getCrateStacks(inputInfo []string, cratesPositionLineIdx int) []common.Stack[string] {
	var stacks []common.Stack[string]
	for idx, ch := range inputInfo[cratesPositionLineIdx] {
		if unicode.IsDigit(ch) {
			var stack common.Stack[string]

			jdx := cratesPositionLineIdx - 1
			crateSign := rune(inputInfo[jdx][idx])

			for unicode.IsLetter(crateSign) && jdx >= 0 {
				stack.Push(string(crateSign))
				if jdx == 0 {
					break
				}

				jdx--
				if len(inputInfo[jdx]) <= idx {
					break
				}

				crateSign = rune(inputInfo[jdx][idx])
			}

			stacks = append(stacks, stack)
		}
	}

	return stacks
}

func getCraneMovesArray(inputInfo []string, craneMovesFirstLineIdx int) [][3]int {
	var craneMoves [][3]int
	for _, line := range inputInfo[craneMovesFirstLineIdx:] {
		var craneMove [3]int
		var idx int

		infoArray := strings.Split(line, " ")
		for _, info := range infoArray {
			if move, err := strconv.Atoi(info); err == nil {
				craneMove[idx] = move
				idx++
			}
		}

		craneMoves = append(craneMoves, craneMove)
	}

	return craneMoves
}

func moveCratesForCrateMover9000(crateStacks []common.Stack[string], craneMoves [][3]int) error {
	for _, move := range craneMoves {
		cratesToMove := move[0]
		for cratesToMove > 0 {
			stackToPop := move[1] - 1
			stackToPush := move[2] - 1

			crateSign, err := crateStacks[stackToPop].Pop()
			if err != nil {
				return fmt.Errorf("pop crate stack: %w", err)
			}

			crateStacks[stackToPush].Push(crateSign)
			cratesToMove--
		}
	}

	return nil
}

func moveCratesForCrateMover9001(crateStacks []common.Stack[string], craneMoves [][3]int) error {
	for _, move := range craneMoves {
		cratesToMove := move[0]
		var tempStack common.Stack[string]
		for cratesToMove > 0 {
			stackToPop := move[1] - 1

			crateSign, err := crateStacks[stackToPop].Pop()
			if err != nil {
				return fmt.Errorf("pop crate stack: %w", err)
			}

			tempStack.Push(crateSign)
			cratesToMove--
		}

		for {
			stackToPush := move[2] - 1
			value, err := tempStack.Pop()
			if err != nil {
				break
			}

			crateStacks[stackToPush].Push(value)
		}
	}

	return nil
}

func readTopCrateFromEachStack(crateStacks []common.Stack[string]) string {
	sb := strings.Builder{}

	for idx := range crateStacks {
		stackTopCrate, err := crateStacks[idx].Peek()
		if err != nil {
			continue
		}

		sb.WriteString(stackTopCrate)
	}

	return sb.String()
}
