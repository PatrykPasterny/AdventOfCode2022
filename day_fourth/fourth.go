package day_fourth

import (
	datareaders "AdventOfCode2022/data_readers"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var (
	errTooLongGroup      = errors.New("elven group should have only two elves")
	errTooLongAssignment = errors.New("assignment should have only two boundaries")
)

func SumFullyOverlappingAssignments(filepath string) (int, error) {
	var result int

	groups, err := getElvenGroupsWithAssignments(filepath)
	if err != nil {
		return 0, fmt.Errorf("get groups' assignment info from file: %w", err)
	}

	for _, group := range groups {
		if areAssignmentsFullyOverlapping(group) {
			result++
		}
	}

	return result, nil
}

func SumEveryOverlappingAssignments(filepath string) (int, error) {
	var result int

	groups, err := getElvenGroupsWithAssignments(filepath)
	if err != nil {
		return 0, fmt.Errorf("get groups' assignment info from file: %w", err)
	}

	for _, group := range groups {
		if areAssignmentsOverlapping(group) {
			result++
		}
	}

	return result, nil
}

func getElvenGroupsWithAssignments(filename string) ([][2][2]int, error) {
	var result [][2][2]int

	groups, err := datareaders.GetDataFromFile(filename)
	if err != nil {
		return nil, fmt.Errorf("get data from file: %w", err)
	}

	for _, group := range groups {
		var resultGroup [2][2]int
		assignments := strings.Split(group, ",")
		if len(assignments) != 2 {
			return nil, errTooLongGroup
		}

		for idx, assignment := range assignments {
			elvenAssignment := strings.Split(assignment, "-")
			if len(elvenAssignment) != 2 {
				return nil, errTooLongAssignment
			}
			assignmentStart, err := strconv.Atoi(elvenAssignment[0])
			if err != nil {
				return nil, fmt.Errorf("parse assignment left bound: %w", err)
			}

			assignmentEnd, err := strconv.Atoi(elvenAssignment[1])
			if err != nil {
				return nil, fmt.Errorf("parse assignment right bound: %w", err)
			}

			resultGroup[idx] = [2]int{assignmentStart, assignmentEnd}
		}

		result = append(result, resultGroup)
	}

	return result, nil
}

func areAssignmentsFullyOverlapping(group [2][2]int) bool {
	if group[0][0] == group[1][0] || group[0][1] == group[1][1] {
		return true
	}

	combinedAssignments := []int{group[0][0], group[0][1], group[1][0], group[1][1]}
	sort.Ints(combinedAssignments)

	if group[0][0] == combinedAssignments[0] {
		return group[0][1] == combinedAssignments[3]
	} else {
		return group[1][1] == combinedAssignments[3]
	}
}

func areAssignmentsOverlapping(group [2][2]int) bool {
	if group[0][0] == group[1][0] || group[0][1] == group[1][1] {
		return true
	}

	if group[0][0] < group[1][0] {
		return group[0][1] >= group[1][0]
	} else {
		return group[1][1] >= group[0][0]
	}
}
