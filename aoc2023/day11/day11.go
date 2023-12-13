package aoc2023

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type galaxy struct {
	x int
	y int
}

func Day11() {
	filepath := os.Getenv("PWD") + "/aoc2023/day11/input.txt"
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Cannot open the file: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	rowLen := 0
	columnLen := 0
	galaxyList := []galaxy{}
	rowEmptyList := []int{}
	columnList := make([]bool, 150)

	result := 0
	for scanner.Scan() {
		str := scanner.Text()
		columnLen = len(str)
		flag := false
		for i := 0; i < len(str); i++ {
			if string(str[i]) == "#" {
				columnList[i] = true
				flag = true
				galaxyList = append(galaxyList, galaxy{rowLen, i})
			}
		}
		if !flag {
			rowEmptyList = append(rowEmptyList, rowLen)
		}
		rowLen++
	}
	columnEmptyList := findEmptyColumnList(columnList, columnLen)
	for i := range galaxyList {
		for j := i + 1; j < len(galaxyList); j++ {
			orgLength := math.Abs(float64(galaxyList[j].x-galaxyList[i].x)) + math.Abs(float64(galaxyList[j].y-galaxyList[i].y))
			length := int(orgLength) + findDoubleLength(galaxyList[j].x, galaxyList[i].x, rowEmptyList) +
				findDoubleLength(galaxyList[j].y, galaxyList[i].y, columnEmptyList)
			result = result + length
		}
	}
	fmt.Print(result)
}

func findEmptyColumnList(columnList []bool, columnLen int) []int {
	emptyColumnList := []int{}
	for index := range columnList {
		if !columnList[index] && index < columnLen {
			emptyColumnList = append(emptyColumnList, index)
		}
	}
	return emptyColumnList
}

func findDoubleLength(x, y int, emptyList []int) int {
	count := 0
	var a, b int
	if x > y {
		a = y
		b = x
	} else {
		a = x
		b = y
	}
	for _, i := range emptyList {
		if i > a && i < b {
			count++
		}
	}
	return count*999999
}
