package day18

import (
	utils "adventOfCode/aoc2024/utils"
	"fmt"
	"log"
	"os"
	"strings"
)

func recursiveMemorySpace(memorySpace [][]bool, visited map[[2]int]bool, x, y, count int) (int, bool) {
	if x < 0 || y < 0 || x > 70 || y > 70 {
		return 0, false
	} else if memorySpace[x][y] || visited[[2]int{x, y}] {
		return 0, false
	} else {
		count++
	}
	if x == 0 && y == 0 {
		fmt.Println(count)
		return count, true
	}
	visited[[2]int{x, y}] = true
	step1, find1 := recursiveMemorySpace(memorySpace, visited, x+1, y, count)
	step2, find2 := recursiveMemorySpace(memorySpace, visited, x-1, y, count)
	step3, find3 := recursiveMemorySpace(memorySpace, visited, x, y+1, count)
	step4, find4 := recursiveMemorySpace(memorySpace, visited, x, y-1, count)
	//visited[[2]int{x, y}] = false
	if find1 || find2 || find3 || find4 {
		minStep := 99999
		if find1 && step1 < minStep {
			minStep = step1
		}
		if find2 && step2 < minStep {
			minStep = step2
		}
		if find3 && step3 < minStep {
			minStep = step3
		}
		if find4 && step4 < minStep {
			minStep = step4
		}
		return minStep, true
	}
	return 0, false
}

func Day() {
	filepath := os.Getenv("PWD") + "/aoc2024/day18/input.txt"
	inputs, err := utils.InputReader(filepath)
	if err != nil {
		log.Fatalf("Cannot open the file: %v", err)
	}
	memorySpace := make([][]bool, 71)
	for i := range memorySpace {
		memorySpace[i] = make([]bool, 71)
	}
	for i, input := range inputs {
		if i == 1024 {
			break
		}
		strInts := strings.Split(input, ",")
		ints, _ := utils.ConvertStrsToInts(strInts)
		memorySpace[ints[1]][ints[0]] = true
	}
	visited := make(map[[2]int]bool)
	fmt.Println(recursiveMemorySpace(memorySpace, visited, 70, 70, 0))
}
