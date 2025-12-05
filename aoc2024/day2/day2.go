package aoc2024

import (
	utils "adventOfCode/utils"
	_"fmt"
	"log"
	"os"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isSafe(levels []int) bool {
	if len(levels) < 2 {
		return true
	}

	isIncreasing := true
	isDecreasing := true

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		absDiff := abs(diff)
		if absDiff < 1 || absDiff > 3 {
			return false
		}
		if diff > 0 {
			isDecreasing = false
		} else if diff < 0 {
			isIncreasing = false
		}
	}

	return isIncreasing || isDecreasing
}

//TODO: 可以根据isSafe的return来优化
func canBeSafeWithRemoval(levels []int) bool {
	for i := 0; i < len(levels); i++ {
		modifiedLevels := append([]int{}, levels[:i]...)
		modifiedLevels = append(modifiedLevels, levels[i+1:]...)
		if isSafe(modifiedLevels) {
			return true
		}
	}
	return false
}

func Day2(){
	filepath := os.Getenv("PWD") + "/aoc2024/day2/input.txt"
	inputs, err := utils.InputReader(filepath)
	if err != nil {
		log.Fatalf("Cannot open the file: %v", err)
	}
	fistSafeReports := 0
	secondSafeReports := 0
	for _, input := range inputs{
		strs := strings.Split(input, " ")
		reports, err := utils.ConvertStrsToInts(strs)
		if err != nil {
			log.Fatalf("Cannot convert the string to int: %v", err)
		}
		if isSafe(reports) {
			fistSafeReports++
		} else if canBeSafeWithRemoval(reports) {
			secondSafeReports++
		}
	}
	//part 1
	println(fistSafeReports)
	//part 2
	println(fistSafeReports+secondSafeReports)
}