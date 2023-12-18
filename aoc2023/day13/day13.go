package aoc2023

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Day13() {
	filepath := os.Getenv("PWD") + "/aoc2023/day13/input.txt"
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Cannot open the file: %v", err)
	}
	defer f.Close()
	result := 0
	scanner := bufio.NewScanner(f)
	found := false
	pattern := []string{}
	for scanner.Scan() {
		str := scanner.Text()
		if str != "" {
			pattern = append(pattern, str)
		} else {
			found = false
			for row := 1; row < len(pattern); row++ {
				if isMirrorSmudge(row-1, row, pattern) {
					result += row * 100
					found = true
					fmt.Printf("rows: %v\n", row)
					break
				}
			}
			if !found {
				columnPattern := changePatternInColumn(pattern)
				for column := 1; column < len(columnPattern); column++ {
					if isMirrorSmudge(column-1, column, columnPattern) {
						result += column
						found = true
						fmt.Printf("column: %v\n", column)
						break
					}
				}

			}
			pattern = []string{}
		}
	}
	fmt.Println(result)
}

func changePatternInColumn(pattern []string) []string {
	columnLen := len(pattern[0])
	columnPattern := []string{}
	for i := 0; i < columnLen; i++ {
		column := ""
		for j := 0; j < len(pattern); j++ {
			column += string(pattern[j][i])
		}
		columnPattern = append(columnPattern, column)
	}
	return columnPattern
}

func findSimilarity(pattern []string) []int {
	samePlace := []int{}
	for i := 1; i < len(pattern); i++ {
		if pattern[i] == pattern[i-1] {
			samePlace = append(samePlace, i)
		}
	}
	return samePlace
}

func isMirror(x, y int, pattern []string) bool {
	for x >= 0 && y >= 0 && x < len(pattern) && y < len(pattern) {
		if pattern[x] != pattern[y] {
			return false
		}
		x--
		y++
	}
	return true
}

func isMirrorSmudge(x, y int, pattern []string) bool {
	count := 0
	for x >= 0 && y >= 0 && x < len(pattern) && y < len(pattern) {
		if pattern[x] != pattern[y] {
			if checkDifference(pattern[x], pattern[y]) {
				count++
			} else{
				return false
			}
			// fmt.Println(pattern[x], pattern[y])
			// return false
		}
		x--
		y++
	}
	return count == 1
}

func checkDifference(s1, s2 string) bool {
	count := 0
	for i := 0; i < len(s1) && i < len(s2); i++ {
		if string(s1[i]) != string(s2[i]) {
			count++
			if count > 1 {
				break
			}
		}
	}
	return count == 1
}
