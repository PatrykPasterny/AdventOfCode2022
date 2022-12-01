package main

import (
	"AdventOfCode2022/day_first"
	"fmt"
	"log"
)

const (
	firstDayFilepath = "/Users/patrykpasterny/Desktop/MyStuff/AdventOfCode2022/day_first/data/data.txt"
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

	fmt.Println()

	fmt.Println()
	fmt.Println("####################### TASK TWO  ########################")

	fmt.Println()

	fmt.Println()
	fmt.Println("####################### DAY THIRD ########################")
	fmt.Println("##########################################################")
	fmt.Println("####################### TASK ONE  ########################")

}
