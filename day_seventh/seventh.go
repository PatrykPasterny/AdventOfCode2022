package day_seventh

import (
	"AdventOfCode2022/common"
	datareaders "AdventOfCode2022/data_readers"
	"AdventOfCode2022/day_seventh/model"
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

var memoryTooSmallForUpdate = errors.New("device memory is too small for update")

func SumAllDirectoriesAboveGivenSize(filepath string, size int) (int, error) {
	filesystem, err := getFilesystem(filepath)
	if err != nil {
		return 0, fmt.Errorf("get filesystem: %w", err)
	}

	result := sumAllDirectoriesLighterThanSize(filesystem, size)

	return result, nil
}

func FindSmallestDirSizeNeededToDeleteForSystemUpdate(filepath string, memorySize, updateSize int) (int, error) {
	filesystem, err := getFilesystem(filepath)
	if err != nil {
		return 0, fmt.Errorf("get filesystem: %w", err)
	}

	spaceTaken := filesystem.TotalSize
	freeSpace := memorySize - spaceTaken

	if freeSpace >= updateSize {
		return 0, nil
	}

	spaceNeeded := updateSize - freeSpace
	result, err := getSmallestDirSizeOverGivenSize(filesystem, spaceNeeded)
	if err != nil {
		return 0, fmt.Errorf("get smallest dir size over the size needed for device update: %w", err)
	}

	return result, nil
}

func getFilesystem(filepath string) (*model.Directory, error) {
	inputData, err := datareaders.GetDataFromFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("get data from file: %w", err)
	}

	filesystem := &model.Directory{
		SubDirectory: make(map[string]*model.Directory),
		Name:         "/",
	}

	err = buildFilesystem(filesystem, inputData)
	if err != nil {
		return nil, fmt.Errorf("build filesystem: %w", err)
	}

	countTotalSizes(filesystem)

	return filesystem, nil
}

func buildFilesystem(filesystem *model.Directory, inputData []string) error {
	const (
		cdRegex   = `^\$.([a-k,m-z]+) (.+)`
		dirRegex  = `^[a-z]+ ([a-z]+)`
		fileRegex = `^(\d.+) (.+)`
	)

	currentDir := filesystem
	regexMap := map[string]*regexp.Regexp{
		"cd":   regexp.MustCompile(cdRegex),
		"dir":  regexp.MustCompile(dirRegex),
		"file": regexp.MustCompile(fileRegex),
	}

	for _, line := range inputData {
		if regexMap["cd"].MatchString(line) {
			groupMatches := regexMap["cd"].FindStringSubmatch(line)
			currentDir = handleCdRegexMatch(groupMatches, currentDir, filesystem)
		} else if regexMap["dir"].MatchString(line) {
			groupMatches := regexMap["dir"].FindStringSubmatch(line)
			handleDirRegexMatch(groupMatches, currentDir)
		} else if regexMap["file"].MatchString(line) {
			groupMatches := regexMap["file"].FindStringSubmatch(line)
			if err := handleFileRegexMatch(groupMatches, currentDir); err != nil {
				return fmt.Errorf("handle 'file' regex match: %w", err)
			}
		}
	}

	return nil
}

func countTotalSizes(node *model.Directory) {
	totalSum := 0
	for idx := range node.SubDirectory {
		countTotalSizes(node.SubDirectory[idx])
		totalSum += node.SubDirectory[idx].TotalSize
	}

	for idx := range node.Files {
		totalSum += node.Files[idx].Size
	}

	node.TotalSize = totalSum
}

func sumAllDirectoriesLighterThanSize(node *model.Directory, size int) int {
	var directoriesSum int
	if node.TotalSize <= size {
		directoriesSum += node.TotalSize
	}

	for idx := range node.SubDirectory {
		currSum := sumAllDirectoriesLighterThanSize(node.SubDirectory[idx], size)
		directoriesSum += currSum
	}

	return directoriesSum
}

func handleCdRegexMatch(groupMatches []string, currentDir *model.Directory, root *model.Directory) *model.Directory {
	result := currentDir
	commandName := groupMatches[1]
	nextDirName := groupMatches[2]

	if commandName == "ls" {
		return result
	}

	switch nextDirName {
	case "..":
		result = currentDir.ParentDirectory
	case "/":
		result = root
	default:
		result = currentDir.SubDirectory[nextDirName]
	}

	return result
}

func handleDirRegexMatch(groupMatches []string, currentDir *model.Directory) {
	dirName := groupMatches[1]
	newSubDir := &model.Directory{
		ParentDirectory: currentDir,
		SubDirectory:    make(map[string]*model.Directory),
		Name:            dirName,
	}

	currentDir.SubDirectory[newSubDir.Name] = newSubDir
}

func handleFileRegexMatch(groupMatches []string, currentDir *model.Directory) error {
	fileName := groupMatches[2]

	fileSize, err := strconv.Atoi(groupMatches[1])
	if err != nil {
		return fmt.Errorf("cast file size to integer: %w", err)
	}

	newSubFile := &model.File{
		Name: fileName,
		Size: fileSize,
	}

	currentDir.Files = append(currentDir.Files, newSubFile)

	return nil
}

func getSmallestDirSizeOverGivenSize(root *model.Directory, size int) (int, error) {
	result := math.MaxInt32

	if size > root.TotalSize {
		return 0, memoryTooSmallForUpdate
	}

	var breadthSearchQueue common.Queue[*model.Directory]

	breadthSearchQueue.Enqueue(root)
	for breadthSearchQueue.Len() > 0 {
		currentNode, err := breadthSearchQueue.Dequeue()
		if err != nil {
			return 0, fmt.Errorf("dequeue during breadth-in search: %w", err)
		}

		if currentNode.TotalSize >= size {
			if currentNode.TotalSize < result {
				result = currentNode.TotalSize
			}
		}

		for key := range currentNode.SubDirectory {
			breadthSearchQueue.Enqueue(currentNode.SubDirectory[key])
		}
	}

	return result, nil
}
