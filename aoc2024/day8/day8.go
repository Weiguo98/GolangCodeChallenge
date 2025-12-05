package day8

import (
	utils "adventOfCode/utils"
	"fmt"
	"log"
	"os"
)

func getLetterMap(inputs []string) map[string][]int {
	letterMap := make(map[string][]int)
	for i, input := range inputs {
		for j, letter := range input {
			if string(letter) != "." {
				_, ok := letterMap[string(letter)]
				if ok {
					letterMap[string(letter)] = append(letterMap[string(letter)], i, j)
				} else {
					letterMap[string(letter)] = []int{i, j}
				}
			}
		}
	}
	return letterMap
}

func getEmptyAntiNodes(inputs []string) [][]bool {
	antiNodes := make([][]bool, len(inputs))
	for i := range antiNodes {
		antiNodes[i] = make([]bool, len(inputs[0]))
	}
	return antiNodes
}

// y2>y1
func findAntinodesForOnePair(x1, y1, x2, y2 int) []int {
	return []int{2*x1 - x2, 2*y1 - y2, 2*x2 - x1, 2*y2 - y1}
}


func setAntinodes(antiNodes [][]bool, nodesPair []int) [][]bool {
	for i := 0; i < len(nodesPair); i += 2 {
		if nodesPair[i] >= 0 && nodesPair[i+1] >= 0 && nodesPair[i] < len(antiNodes) && nodesPair[i+1] < len(antiNodes[0]) {
			antiNodes[nodesPair[i]][nodesPair[i+1]] = true
		}
	}
	return antiNodes
}

func findAllAntinodesForOnePair(x1, y1, x2, y2, height, length int) []int {
	res := []int{}
	for n := 0; n<=30 ; n++ {
		if (n+1)*x1-n*x2 >= 0 && (n+1)*y1-n*y2 >= 0 && (n+1)*x1-n*x2 < height && (n+1)*y1-n*y2 < length {
			res = append(res, (n+1)*x1-n*x2, (n+1)*y1-n*y2)
		}
		if (n+1)*x2-n*x1 >= 0 && (n+1)*y2-n*y1 >= 0 && (n+1)*x2-n*x1 < height && (n+1)*y2-n*y1 < length {
			res = append(res, (n+1)*x2-n*x1, (n+1)*y2-n*y1)
		}
	}
	return res
}

func getAntinodes(letterMap map[string][]int, antiNodes [][]bool) [][]bool {
	for _, v := range letterMap {
		for i := 0; i < len(v); i += 2 {
			for j := i + 2; j < len(v); j += 2 {
				// part 1
				// nodesPair := findAntinodesForOnePair(v[i], v[i+1], v[j], v[j+1])
				// part 2
				nodesPair := findAllAntinodesForOnePair(v[i], v[i+1], v[j], v[j+1], len(antiNodes), len(antiNodes[0]))
				antiNodes = setAntinodes(antiNodes, nodesPair)
			}
		}
	}
	return antiNodes
}

func Day() {
	filepath := os.Getenv("PWD") + "/aoc2024/day8/input.txt"
	inputs, err := utils.InputReader(filepath)
	if err != nil {
		log.Fatalf("Cannot open the file: %v", err)
	}
	letterMap := getLetterMap(inputs)
	antiNodes := getAntinodes(letterMap, getEmptyAntiNodes(inputs))
	nrOfAntinodes := 0
	for i := 0; i < len(antiNodes); i++ {
		for j := 0; j < len(antiNodes[0]); j++ {
			if antiNodes[i][j] {
				nrOfAntinodes++
			}
		}
	}
	fmt.Println(nrOfAntinodes)
}
