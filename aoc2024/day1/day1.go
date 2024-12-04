package aoc2024

import (
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	utils "adventOfCode/aoc2024/utils"
)

func Day1() {
	filepath := os.Getenv("PWD") + "/aoc2024/day1/input.txt"
	inputs, err := utils.InputReader(filepath)
	if err != nil {
		log.Fatalf("Cannot open the file: %v", err)
	}
	array1 := []int{}
	array2 := []int{}
	for _, input := range inputs {
		strs := strings.Split(input, "   ")
		num1, err := strconv.Atoi(strs[0])
		if err != nil {
			log.Fatalf("Cannot convert the string to int: %v", err)
		}
		num2, err := strconv.Atoi(strs[1])
		if err != nil {
			log.Fatalf("Cannot convert the string to int: %v", err)
		}
		array1 = append(array1, num1)
		array2 = append(array2, num2)
	}
	sort.Ints(array1)
	sort.Ints(array2)
	// part 1
	res :=  0.0
	for i := 0; i < len(array1); i++ {
		res += math.Abs(float64(array1[i] - array2[i]))
	}
	println(int(res))
	// part 2
	resArray := make([]int, len(array1))
	j := 0
	for i := 0; i < len(array1); {
		if array1[i]<array2[j] {
			i++
		} else if array1[i]>array2[j] {
			j++
		} else{
			resArray[i]++
			j++
		}
	}
	result := 0
	for i, m := range resArray {
		result += array1[i]*m
	}
	println(result)
}
