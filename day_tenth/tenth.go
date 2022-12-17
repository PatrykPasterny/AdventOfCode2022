package day_tenth

import (
	datareaders "AdventOfCode2022/data_readers"

	"fmt"
	"strconv"
	"strings"
)

func GetSumOfStrengthsOfGivenCycles(filepath string, cycles []int) (int, error) {
	if len(cycles) == 0 {
		return 0, nil
	}
	cyclesToSumMap := make(map[int]bool)

	for i := range cycles {
		cyclesToSumMap[cycles[i]] = true
	}

	commands, err := datareaders.GetDataFromFile(filepath)
	if err != nil {
		return 0, fmt.Errorf("getting data from file: %w", err)
	}

	commandToLengthMap := map[string]int{
		"noop": 0,
		"addx": 1,
	}

	var result, commandCounter, idx, ticksToWait int

	var commandComponents []string

	cyclesCounter := 1
	xRegister := 1

	for {
		if idx >= len(commands) {
			break
		}

		if _, ok := cyclesToSumMap[cyclesCounter]; ok {
			result += cyclesCounter * xRegister
		}

		if commandCounter == 0 {
			command := commands[idx]
			commandComponents = strings.Split(command, " ")
			ticksToWait = commandToLengthMap[commandComponents[0]]
		}

		if commandCounter == ticksToWait {
			if len(commandComponents) > 1 {
				xRegisterIncrement, err := strconv.Atoi(commandComponents[1])
				if err != nil {
					return 0, fmt.Errorf("converting addx increment to int: %w", err)
				}

				xRegister += xRegisterIncrement
			}

			idx++
			commandCounter = 0
		} else {
			commandCounter++
		}

		cyclesCounter++
	}

	return result, nil
}

func GetScreenOutput(filepath string) (*[6][40]string, error) {
	var screen [6][40]string

	commands, err := datareaders.GetDataFromFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("getting data from file: %w", err)
	}

	commandToLengthMap := map[string]int{
		"noop": 0,
		"addx": 1,
	}

	var commandCounter, idx, ticksToWait int

	var commandComponents []string

	cyclesCounter := 1
	xRegister := 1

	for {
		if idx >= len(commands) {
			break
		}

		if commandCounter == 0 {
			command := commands[idx]
			commandComponents = strings.Split(command, " ")
			ticksToWait = commandToLengthMap[commandComponents[0]]
		}

		screenJdx := (cyclesCounter - 1) % 40
		screenIdx := (cyclesCounter - 1) / 40
		pixelPosition := [3]int{xRegister - 1, xRegister, xRegister + 1}
		if screenJdx == pixelPosition[0] || screenJdx == pixelPosition[1] ||
			screenJdx == pixelPosition[2] {
			screen[screenIdx][screenJdx] = "#"
		} else {
			screen[screenIdx][screenJdx] = "."
		}

		if commandCounter == ticksToWait {
			if len(commandComponents) > 1 {
				xRegisterIncrement, err := strconv.Atoi(commandComponents[1])
				if err != nil {
					return nil, fmt.Errorf("converting addx increment to int: %w", err)
				}

				xRegister += xRegisterIncrement
			}

			idx++
			commandCounter = 0
		} else {
			commandCounter++
		}

		cyclesCounter++
	}

	return &screen, nil
}
