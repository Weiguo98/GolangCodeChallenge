package main

import (
	// "bufio"
	"fmt"
	// "os"
	// "sort"
	"math"
	"strconv"
)

func isLetter(l string) bool {
	if "A" <= l && l <= "Z" {
		return true
	} else {
		return false
	}
}

func isNumber(l string) bool {
	if "0" <= l && l <= "9" {
		return true
	} else {
		return false
	}
}

func getMoleculesMap(m map[string]int, inputMolecules string) map[string]int {
	preIsLetter := false
	preIsNumber := false
	tempNumber := ""
	currentLetter := ""
	for _, l := range inputMolecules {
		item := string(l)
		// fmt.Println(item)
		if preIsLetter && isNumber(item) {
			tempNumber = tempNumber + item
			preIsNumber = true
			preIsLetter = false
		} else if preIsNumber && isNumber(item) {
			// fmt.Println("here")
			tempNumber = tempNumber + item
		} else if preIsNumber && isLetter(item) {
			currentNum, ok := m[currentLetter]
			num, _ := strconv.Atoi(tempNumber)
			if ok {
				m[currentLetter] = currentNum + num
			} else {
				m[currentLetter] = num
			}
			tempNumber = ""
			currentLetter = item
			preIsLetter = true
			preIsNumber = false
		} else if preIsLetter && isLetter(item) {
			m[currentLetter] += 1
			currentLetter = item
		} else if !preIsLetter && !preIsNumber && isLetter(item) {
			currentLetter = item
			preIsLetter = true
		}
		// else{
		//     fmt.Println(preIsNumber)
		//     fmt.Println(preIsLetter)
		//     fmt.Println(tempNumber)
		//     fmt.Println(currentLetter)
		// }
	}
	if preIsLetter {
		m[currentLetter] += 1
	} else if preIsNumber {
		currentNum, ok := m[currentLetter]
		num, _ := strconv.Atoi(tempNumber)
		if ok {
			m[currentLetter] = currentNum + num
		} else {
			m[currentLetter] = num
		}
	}

	return m
}

func main() {
	var inputMolecules, outputMolecules string
	var num int
	fmt.Scanf("%s %v\n", &inputMolecules, &num)
	fmt.Scanf("%s", &outputMolecules)

	m := make(map[string]int)
	n := make(map[string]int)

	m = getMoleculesMap(m, inputMolecules)
	n = getMoleculesMap(n, outputMolecules)
	// fmt.Println(m)
	// fmt.Println(n)
	if num > 1 {
		for key, value := range m {
			m[key] = value * num
		}
	}

	res := math.MaxInt
	for key, value := range n {
		mValue, ok := m[key]
		if ok {
			if mValue < value {
				res = 0
				break
			} else {
				temp := mValue / value
				if temp < res {
					res = temp
				}
			}
		} else {
			res = 0
			break
		}
	}

	fmt.Println(res)

}
