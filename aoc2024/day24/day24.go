package day24

import (
	utils "adventOfCode/aoc2024/utils"
	"fmt"
	"log"
	"os"
	"strings"
)

func isAllVisited(visited []bool) bool {
	for _, v := range visited {
		if !v {
			return false
		}
	}
	return true
}

func Day() {
	filepath := os.Getenv("PWD") + "/aoc2024/day24/input.txt"
	inputs, err := utils.InputReader(filepath)
	if err != nil {
		log.Fatalf("Cannot open the file: %v", err)
	}
	gates := make(map[string]bool, 0)
	formulas := []string{}
	for i, input := range inputs {
		if i < 90 {
			gates[input[:3]] = input[len(input)-1] == '1'
		} else if i == 90 {
			continue
		} else {
			formulas = append(formulas, input)
		}
	}
	// visited := make([]bool, len(formulas))
	// day 1
	// for !isAllVisited(visited) {
	// 	for i, formula := range formulas {
	// 		wires := strings.Split(formula, " -> ")
	// 		words := strings.Split(wires[0], " ")
	// 		word1, ok1 := gates[words[0]]
	// 		word2, ok2 := gates[words[2]]
	// 		if ok1 && ok2 {
	// 			visited[i] = true
	// 			if words[1] == "AND" {
	// 				gates[wires[1]] = word1 && word2
	// 			} else if words[1] == "OR" {
	// 				gates[wires[1]] = word1 || word2
	// 			} else if words[1] == "XOR" {
	// 				gates[wires[1]] = word1 != word2
	// 			}
	// 		}
	// 	}
	// }
	// keys := []string{}
	// for k := range gates {
	// 	if string(k[0]) == "z" {
	// 		keys = append(keys, k)
	// 	}
	// }
	// sort.Strings(keys)
	// res := 0
	// for i := len(keys) - 1; i >= 0; i-- {
	// 	if gates[keys[i]] {
	// 		res += 1 << i
	// 	}
	// }
	// fmt.Println(res)

	//day2
	for _, formula := range formulas {
		wires := strings.Split(formula, " -> ")
		words := strings.Split(wires[0], " ")
		if string(words[0][0]) != "x" || string(words[2][0]) != "y" || words[1] != "AND" {
			continue
		}

	}
	fmt.Println(gates)
}
