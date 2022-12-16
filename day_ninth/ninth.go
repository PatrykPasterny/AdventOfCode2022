package day_ninth

import (
	datareaders "AdventOfCode2022/data_readers"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	keyFormat = "[%d, %d]"
	epsilon   = 0.00001
)

var (
	errUndefinedMoveDirection = errors.New("given move direction is not defined, choose between U R D or L")
	errRopeTooShort           = errors.New("rope needs to have at least 2 knots")
)

type HeadMove struct {
	Direction string
	Steps     int
}

type Position struct {
	x int
	y int
}

type Velocity struct {
	vx int
	vy int
}

func CountAllTailPositions(filepath string, ropeLength int) (int, error) {
	if ropeLength < 2 {
		return 0, errRopeTooShort
	}
	data, err := datareaders.GetDataFromFile(filepath)
	if err != nil {
		return 0, fmt.Errorf("get data from file: %w", err)
	}

	headMovesDetails, err := decodeRopeMoves(data)
	if err != nil {
		return 0, fmt.Errorf("decoding rope moves: %w", err)
	}

	tailVisited := make(map[string]bool)

	var nodesPositions []*Position
	var nodeVelocities []*Velocity

	for i := 0; i < ropeLength; i++ {
		nodesPositions = append(nodesPositions, &Position{x: 0, y: 0})
		nodeVelocities = append(nodeVelocities, &Velocity{vx: 0, vy: 0})
	}

	tailVisited[fmt.Sprintf(keyFormat, nodesPositions[ropeLength-1].x, nodesPositions[ropeLength-1].y)] = true

	for i := range headMovesDetails {
		nodeVelocities[0].vx, nodeVelocities[0].vy, err = determineHeadVelocity(headMovesDetails[i].Direction)
		if err != nil {
			return 0, fmt.Errorf("determining head velocity: %w", err)
		}

		steps := headMovesDetails[i].Steps
		simulateRopeMoves(tailVisited, nodesPositions, nodeVelocities, steps)
	}

	result := sumAppearancesOfTrue(tailVisited)

	return result, nil
}

func simulateRopeMoves(tailVisited map[string]bool,
	nodePositions []*Position,
	nodeVelocities []*Velocity,
	steps int,
) {
	for steps > 0 {
		nodePositions[0].x += nodeVelocities[0].vx
		nodePositions[0].y += nodeVelocities[0].vy

		for i := 1; i < len(nodePositions); i++ {
			nodesDistance := countDistance(nodePositions[i-1], nodePositions[i])
			if nodesDistance+epsilon > 2 {
				if nodesDistance-epsilon > 2 {
					nodeVelocities[i] = findDiagonalMove(nodePositions[i-1], nodePositions[i], nodeVelocities[i-1])
				} else {
					nodeVelocities[i] = findStraightMove(nodePositions[i-1], nodePositions[i])
				}

				nodePositions[i].x += nodeVelocities[i].vx
				nodePositions[i].y += nodeVelocities[i].vy

				if i == len(nodePositions)-1 {
					tailVisited[fmt.Sprintf(keyFormat, nodePositions[i].x, nodePositions[i].y)] = true
				}
			}
		}

		steps--
	}
}

func determineHeadVelocity(direction string) (int, int, error) {
	switch direction {
	case "U":
		return 0, 1, nil
	case "R":
		return 1, 0, nil
	case "D":
		return 0, -1, nil
	case "L":
		return -1, 0, nil
	}

	return 0, 0, errUndefinedMoveDirection
}

func findDiagonalMove(successorPos, currentPos *Position, successorVelocity *Velocity) *Velocity {
	var result Velocity

	if successorVelocity.vx != 0 && successorVelocity.vy != 0 {
		result.vx = successorVelocity.vx
		result.vy = successorVelocity.vy
	}
	if successorVelocity.vx != 0 {
		firstPotentialPosition := &Position{
			x: currentPos.x + successorVelocity.vx,
			y: currentPos.y + 1,
		}

		secondPotentialPosition := &Position{
			x: currentPos.x + successorVelocity.vx,
			y: currentPos.y - 1,
		}

		if countDistance(successorPos, firstPotentialPosition) < countDistance(successorPos, secondPotentialPosition) {
			result.vx = successorVelocity.vx
			result.vy = 1
		} else {
			result.vx = successorVelocity.vx
			result.vy = -1
		}
	} else {
		firstPotentialPosition := &Position{
			x: currentPos.x + 1,
			y: currentPos.y + successorVelocity.vy,
		}

		secondPotentialPosition := &Position{
			x: currentPos.x - 1,
			y: currentPos.y + successorVelocity.vy,
		}

		if countDistance(successorPos, firstPotentialPosition) < countDistance(successorPos, secondPotentialPosition) {
			result.vx = 1
			result.vy = successorVelocity.vy
		} else {
			result.vx = -1
			result.vy = successorVelocity.vy

		}
	}

	return &result
}

func findStraightMove(successorPos, currentPos *Position) *Velocity {
	var result Velocity

	if successorPos.y == currentPos.y {
		firstPotentialPosition := &Position{
			x: currentPos.x + 1,
			y: currentPos.y,
		}

		secondPotentialPosition := &Position{
			x: currentPos.x - 1,
			y: currentPos.y,
		}

		if countDistance(successorPos, firstPotentialPosition) < countDistance(successorPos, secondPotentialPosition) {
			result.vx = 1
			result.vy = 0
		} else {
			result.vx = -1
			result.vy = 0
		}
	}

	if successorPos.x == currentPos.x {
		firstPotentialPosition := &Position{
			x: currentPos.x,
			y: currentPos.y + 1,
		}

		secondPotentialPosition := &Position{
			x: currentPos.x,
			y: currentPos.y - 1,
		}

		if countDistance(successorPos, firstPotentialPosition) < countDistance(successorPos, secondPotentialPosition) {
			result.vx = 0
			result.vy = 1
		} else {
			result.vx = 0
			result.vy = -1
		}
	}

	return &result
}

func countDistance(p1, p2 *Position) float64 {
	return math.Sqrt(math.Pow(math.Abs(float64(p2.y-p1.y)), 2) + math.Pow(math.Abs(float64(p2.x-p1.x)), 2))
}

func decodeRopeMoves(data []string) ([]*HeadMove, error) {
	var headMovesDetails []*HeadMove

	for _, line := range data {
		headMoveDetails := strings.Split(line, " ")
		steps, err := strconv.Atoi(headMoveDetails[1])
		if err != nil {
			return nil, fmt.Errorf("unable to cast head steps number to int: %w", err)
		}

		headMove := &HeadMove{
			Direction: headMoveDetails[0],
			Steps:     steps,
		}

		headMovesDetails = append(headMovesDetails, headMove)
	}

	return headMovesDetails, nil
}

func sumAppearancesOfTrue(matrix map[string]bool) int {
	return len(matrix)
}
