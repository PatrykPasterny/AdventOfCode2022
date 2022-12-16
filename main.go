package main

import (
	"AdventOfCode2022/day_eighth"
	"AdventOfCode2022/day_fifth"
	"AdventOfCode2022/day_first"
	"AdventOfCode2022/day_fourth"
	"AdventOfCode2022/day_ninth"
	"AdventOfCode2022/day_second"
	"AdventOfCode2022/day_seventh"
	"AdventOfCode2022/day_sixth"
	"AdventOfCode2022/day_third"
	"fmt"
	"log"
)

const (
	firstDayFilepath   = "/Users/patrykpasterny/Desktop/MyStuff/AdventOfCode2022/day_first/data/data.txt"
	secondDayFilepath  = "/Users/patrykpasterny/Desktop/MyStuff/AdventOfCode2022/day_second/data/data.txt"
	thirdDayFilepath   = "/Users/patrykpasterny/Desktop/MyStuff/AdventOfCode2022/day_third/data/data.txt"
	fourthDayFilepath  = "/Users/patrykpasterny/Desktop/MyStuff/AdventOfCode2022/day_fourth/data/data.txt"
	fifthDayFilepath   = "/Users/patrykpasterny/Desktop/MyStuff/AdventOfCode2022/day_fifth/data/data.txt"
	sixthDayFilepath   = "/Users/patrykpasterny/Desktop/MyStuff/AdventOfCode2022/day_sixth/data/data.txt"
	seventhDayFilepath = "/Users/patrykpasterny/Desktop/MyStuff/AdventOfCode2022/day_seventh/data/data.txt"
	eighthDayFilepath  = "/Users/patrykpasterny/Desktop/MyStuff/AdventOfCode2022/day_eighth/data/data.txt"
	ninthDayFilepath   = "/Users/patrykpasterny/Desktop/MyStuff/AdventOfCode2022/day_ninth/data/data.txt"
)

const (
	deviceMemory = 70000000
	updateWeight = 30000000
)

func main() {
	fmt.Println("####################### DAY FIRST ########################")
	fmt.Println("##########################################################")
	fmt.Println("####################### TASK ONE  ########################")
	dayFirstTaskOneResult, err := day_first.GetBackpackWithMostCalories(firstDayFilepath)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println()
	fmt.Println(dayFirstTaskOneResult)
	fmt.Println()
	fmt.Println("####################### TASK TWO  ########################")
	dayFirstSecondTaskResult, err := day_first.GetThreeBackpacksWithMostCalories(firstDayFilepath)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println()
	fmt.Println(dayFirstSecondTaskResult)
	fmt.Println()
	fmt.Println("###################### DAY SECOND ########################")
	fmt.Println("##########################################################")
	fmt.Println("####################### TASK ONE  ########################")
	daySecondTaskFirstResult, err := day_second.CountFirstStrategyGuideResult(secondDayFilepath)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println()
	fmt.Println(daySecondTaskFirstResult)
	fmt.Println()
	fmt.Println("####################### TASK TWO  ########################")
	daySecondTaskSecondResult, err := day_second.CountSecondStrategyGuideResult(secondDayFilepath)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println()
	fmt.Println(daySecondTaskSecondResult)
	fmt.Println()
	fmt.Println("####################### DAY THIRD ########################")
	fmt.Println("##########################################################")
	fmt.Println("####################### TASK ONE  ########################")
	dayThirdTaskFirstResult, err := day_third.SumRucksacksPrioritySupplies(thirdDayFilepath)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println()
	fmt.Println(dayThirdTaskFirstResult)
	fmt.Println()
	fmt.Println("####################### TASK TWO  ########################")
	dayThirdTaskSecondResult, err := day_third.SumGroupBadgesPriorities(thirdDayFilepath)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println()
	fmt.Println(dayThirdTaskSecondResult)
	fmt.Println()
	fmt.Println("####################### DAY FOURTH #######################")
	fmt.Println("##########################################################")
	fmt.Println("####################### TASK ONE  ########################")
	fmt.Println()
	dayFourthTaskFirstResult, err := day_fourth.SumFullyOverlappingAssignments(fourthDayFilepath)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println()
	fmt.Println(dayFourthTaskFirstResult)
	fmt.Println()
	fmt.Println("####################### TASK TWO  ########################")
	fmt.Println()
	dayFourthTaskSecondResult, err := day_fourth.SumEveryOverlappingAssignments(fourthDayFilepath)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println()
	fmt.Println(dayFourthTaskSecondResult)
	fmt.Println()
	fmt.Println("####################### DAY FIFTH #######################")
	fmt.Println("##########################################################")
	fmt.Println("####################### TASK ONE  ########################")
	fmt.Println()
	dayFifthTaskFirstResult, err := day_fifth.GetTopCratesNamesForCrateMover9000(fifthDayFilepath)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println()
	fmt.Println(dayFifthTaskFirstResult)
	fmt.Println()
	fmt.Println("####################### TASK TWO  ########################")
	fmt.Println()
	dayFifthTaskSecondResult, err := day_fifth.GetTopCratesNamesForCrateMover9001(fifthDayFilepath)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println()
	fmt.Println(dayFifthTaskSecondResult)
	fmt.Println()
	fmt.Println("####################### DAY SIXTH #######################")
	fmt.Println("##########################################################")
	fmt.Println("####################### TASK ONE  ########################")
	fmt.Println()
	daySixthTaskFirstResult, err := day_sixth.FindFirstStartOfPacketMarker(sixthDayFilepath)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println()
	fmt.Println(daySixthTaskFirstResult)
	fmt.Println()
	fmt.Println("####################### TASK TWO  ########################")
	daySixthTaskSecondResult, err := day_sixth.FindFirstStartOfMessageMarker(sixthDayFilepath)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println()
	fmt.Println(daySixthTaskSecondResult)
	fmt.Println()
	fmt.Println("###################### DAY SEVENTH #######################")
	fmt.Println("##########################################################")
	fmt.Println("####################### TASK ONE  ########################")
	daySeventhTaskFirstResult, err := day_seventh.SumAllDirectoriesAboveGivenSize(seventhDayFilepath, 100000)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println()
	fmt.Println(daySeventhTaskFirstResult)
	fmt.Println()
	fmt.Println("####################### TASK TWO  ########################")
	fmt.Println()
	daySeventhTaskSecondResult, err := day_seventh.
		FindSmallestDirSizeNeededToDeleteForSystemUpdate(seventhDayFilepath, deviceMemory, updateWeight)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println()
	fmt.Println(daySeventhTaskSecondResult)
	fmt.Println()
	fmt.Println("####################### DAY EIGHT #######################")
	fmt.Println("##########################################################")
	fmt.Println("####################### TASK ONE  ########################")
	dayEightTaskFirstResult, err := day_eighth.CountVisibleTrees(eighthDayFilepath)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println()
	fmt.Println(dayEightTaskFirstResult)
	fmt.Println()
	fmt.Println("####################### TASK TWO  ########################")
	fmt.Println()
	dayEighthTaskSecondResult, err := day_eighth.FindHighestScenicScore(eighthDayFilepath)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println()
	fmt.Println(dayEighthTaskSecondResult)
	fmt.Println()
	fmt.Println("####################### DAY NINTH #######################")
	fmt.Println("##########################################################")
	fmt.Println("####################### TASK ONE  ########################")
	dayNineTaskFirstResult, err := day_ninth.CountAllTailPositions(ninthDayFilepath, 2)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println()
	fmt.Println(dayNineTaskFirstResult)
	fmt.Println()
	fmt.Println("####################### TASK TWO  ########################")
	fmt.Println()
	dayNineTaskSecondResult, err := day_ninth.CountAllTailPositions(ninthDayFilepath, 10)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println()
	fmt.Println(dayNineTaskSecondResult)
	fmt.Println()
}
