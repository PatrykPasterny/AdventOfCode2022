package day_eighth

import (
	datareaders "AdventOfCode2022/data_readers"
	"fmt"
)

func CountVisibleTrees(filepath string) (int, error) {
	treeHeightsGrid, rowLen, colLen, err := datareaders.GetMatrixDataFromFile(filepath)
	if err != nil {
		return 0, fmt.Errorf("get matrix data from file: %w", err)
	}

	visible := getVisibleGrid(treeHeightsGrid, rowLen, colLen)
	result := sumGridElements(visible)

	return result, nil
}

func FindHighestScenicScore(filepath string) (int, error) {
	treeHeightsGrid, rowLen, colLen, err := datareaders.GetMatrixDataFromFile(filepath)
	if err != nil {
		return 0, fmt.Errorf("get matrix data from file: %w", err)
	}

	result := getHighestScenicPoints(treeHeightsGrid, rowLen, colLen)

	return result, nil
}

func getVisibleGrid(treeHeightsGrid [][]int, rowLen, colLen int) [][]int {
	visible := createVisitedTreeGrid(rowLen, colLen)

	fillVisibleAndTreeGrids(treeHeightsGrid, visible, rowLen, colLen)

	return visible
}

func createVisitedTreeGrid(rowLen, colLen int) [][]int {
	result := make([][]int, rowLen)
	for idx := range result {
		result[idx] = make([]int, colLen)
		for jdx := range result[idx] {
			if idx == 0 || jdx == 0 || idx == rowLen-1 || jdx == colLen-1 {
				result[idx][jdx] = 1
			}
		}
	}

	return result
}

func fillVisibleAndTreeGrids(treeHeightsGrid [][]int, visible [][]int, rowLen, colLen int) {
	for idx := 1; idx < rowLen-1; idx++ {
		leftHighest := treeHeightsGrid[idx][0]
		for jdx := 1; jdx < colLen-1; jdx++ {
			if treeHeightsGrid[idx][jdx] > leftHighest {
				leftHighest = treeHeightsGrid[idx][jdx]
				visible[idx][jdx] = 1
			}
		}

		rightHighest := treeHeightsGrid[idx][colLen-1]
		for jdx := colLen - 2; jdx >= 0; jdx-- {
			if treeHeightsGrid[idx][jdx] > rightHighest {
				rightHighest = treeHeightsGrid[idx][jdx]
				visible[idx][jdx] = 1
			}
		}
	}

	for jdx := 1; jdx < colLen-1; jdx++ {
		topHighest := treeHeightsGrid[0][jdx]
		for idx := 1; idx < rowLen-1; idx++ {
			if treeHeightsGrid[idx][jdx] > topHighest {
				topHighest = treeHeightsGrid[idx][jdx]
				visible[idx][jdx] = 1
			}
		}

		bottomHighest := treeHeightsGrid[rowLen-1][jdx]
		for idx := rowLen - 2; idx >= 0; idx-- {
			if treeHeightsGrid[idx][jdx] > bottomHighest {
				bottomHighest = treeHeightsGrid[idx][jdx]
				visible[idx][jdx] = 1
			}
		}
	}
}

func getHighestScenicPoints(treeHeightsGrid [][]int, rowLen, colLen int) int {
	var result int

	for idx := 1; idx < rowLen-1; idx++ {
		for jdx := 1; jdx < colLen-1; jdx++ {
			leftDistance, topDistance, rightDistance, bottomDistance := 1, 1, 1, 1

			currentHeight := treeHeightsGrid[idx][jdx-leftDistance]
			for currentHeight < treeHeightsGrid[idx][jdx] {
				if jdx == leftDistance {
					break
				}

				leftDistance++
				currentHeight = treeHeightsGrid[idx][jdx-leftDistance]
			}

			currentHeight = treeHeightsGrid[idx-topDistance][jdx]
			for currentHeight < treeHeightsGrid[idx][jdx] {
				if idx == topDistance {
					break
				}

				topDistance++
				currentHeight = treeHeightsGrid[idx-topDistance][jdx]
			}

			currentHeight = treeHeightsGrid[idx][jdx+rightDistance]
			for currentHeight < treeHeightsGrid[idx][jdx] {
				if jdx+rightDistance == colLen-1 {
					break
				}

				rightDistance++
				currentHeight = treeHeightsGrid[idx][jdx+rightDistance]
			}

			currentHeight = treeHeightsGrid[idx+bottomDistance][jdx]
			for currentHeight < treeHeightsGrid[idx][jdx] {
				if idx+bottomDistance == (rowLen - 1) {
					break
				}

				bottomDistance++
				currentHeight = treeHeightsGrid[idx+bottomDistance][jdx]
			}

			scenicPoints := leftDistance * topDistance * rightDistance * bottomDistance
			if result < scenicPoints {
				result = scenicPoints
			}
		}
	}

	return result
}

func sumGridElements(grid [][]int) int {
	var result int
	for idx := range grid {
		for jdx := range grid[idx] {
			result += grid[idx][jdx]
		}
	}

	return result
}
