package day_second

import (
	datareaders "AdventOfCode2022/data_readers"
	"fmt"
	"strings"
)

const (
	OpponentRock                = "A"
	FirstStrategyPlayerRock     = "X"
	SecondStrategyPlayerLoses   = "X"
	OpponentPaper               = "B"
	FirstStrategyPlayerPaper    = "Y"
	SecondStrategyPlayerDraws   = "Y"
	OpponentScissors            = "C"
	FirstStrategyPlayerScissors = "Z"
	SecondStrategyPlayerWins    = "Z"
	RockPlayed                  = 1
	PaperPlayed                 = 2
	ScissorsPlayed              = 3
	Defeat                      = 0
	Draw                        = 3
	Win                         = 6
)

var (
	firstStrategyGameRulesMap = map[string]int{
		OpponentRock + FirstStrategyPlayerRock:         Draw,
		OpponentPaper + FirstStrategyPlayerRock:        Defeat,
		OpponentScissors + FirstStrategyPlayerRock:     Win,
		OpponentRock + FirstStrategyPlayerPaper:        Win,
		OpponentPaper + FirstStrategyPlayerPaper:       Draw,
		OpponentScissors + FirstStrategyPlayerPaper:    Defeat,
		OpponentRock + FirstStrategyPlayerScissors:     Defeat,
		OpponentPaper + FirstStrategyPlayerScissors:    Win,
		OpponentScissors + FirstStrategyPlayerScissors: Draw,
	}
	secondStrategyGameRulesMap = map[string]string{
		OpponentRock + SecondStrategyPlayerLoses:     FirstStrategyPlayerScissors,
		OpponentRock + SecondStrategyPlayerDraws:     FirstStrategyPlayerRock,
		OpponentRock + SecondStrategyPlayerWins:      FirstStrategyPlayerPaper,
		OpponentPaper + SecondStrategyPlayerLoses:    FirstStrategyPlayerRock,
		OpponentPaper + SecondStrategyPlayerDraws:    FirstStrategyPlayerPaper,
		OpponentPaper + SecondStrategyPlayerWins:     FirstStrategyPlayerScissors,
		OpponentScissors + SecondStrategyPlayerLoses: FirstStrategyPlayerPaper,
		OpponentScissors + SecondStrategyPlayerDraws: FirstStrategyPlayerScissors,
		OpponentScissors + SecondStrategyPlayerWins:  FirstStrategyPlayerRock,
	}
)

func CountFirstStrategyGuideResult(filepath string) (int, error) {
	roundResults, err := getRoundsResultFromFile(filepath)
	if err != nil {
		return 0, fmt.Errorf("unable to get results from file %w", err)
	}

	var points int
	for _, result := range roundResults {
		points += countFirstStrategyRoundPoint(result)
	}

	return points, nil
}

func CountSecondStrategyGuideResult(filepath string) (int, error) {
	roundResults, err := getRoundsResultFromFile(filepath)
	if err != nil {
		return 0, fmt.Errorf("unable to get results from file %w", err)
	}

	var points int
	for _, result := range roundResults {
		points += countSecondStrategyRoundPoints(result)
	}

	return points, nil
}

func countFirstStrategyRoundPoint(round []string) int {
	var result int
	switch round[1] {
	case FirstStrategyPlayerRock:
		result += RockPlayed
	case FirstStrategyPlayerPaper:
		result += PaperPlayed
	case FirstStrategyPlayerScissors:
		result += ScissorsPlayed
	}

	result += countRockPaperScissorsOutcome(round)

	return result
}

func countSecondStrategyRoundPoints(round []string) int {
	playerPlayed := round[0]
	opponentPlayed := secondStrategyGameRulesMap[round[0]+round[1]]

	return countFirstStrategyRoundPoint([]string{playerPlayed, opponentPlayed})
}

func countRockPaperScissorsOutcome(round []string) int {
	return firstStrategyGameRulesMap[round[0]+round[1]]
}

func getRoundsResultFromFile(filepath string) ([][]string, error) {
	var rounds [][]string
	roundsInfo, err := datareaders.GetDataFromFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("unable to get backpacks from file %w", err)
	}

	for _, roundResult := range roundsInfo {
		round := strings.Split(roundResult, " ")
		rounds = append(rounds, round)
	}

	return rounds, nil
}
