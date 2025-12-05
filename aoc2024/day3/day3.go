package day3

import (
	utils "adventOfCode/utils"
	_"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func returnMulsFromStr(match string) int {
	str := strings.Split(match, "(")[1]
	str = strings.Split(str, ")")[0]
	strs := strings.Split(str, ",")
	num1, _ := strconv.Atoi(strs[0])
	num2, _ := strconv.Atoi(strs[1])
	return num1 * num2
}

func part1(inputs []string) int {
	res := 0
	for _, input := range inputs {
		re := regexp.MustCompile(`mul\(\d+,\s*\d+\)`)
		matches := re.FindAllString(input, -1)
		for _, match := range matches {
			res += returnMulsFromStr(match)
		}
	}
	return res
}

func part2(inputs []string) int{
	flag := false
	res := 0
	for _, input := range inputs {
		re := regexp.MustCompile(`mul\(\d+,\s*\d+\)|don't\(\)|do\(\)`)
		matches := re.FindAllString(input, -1)
		for _, match := range matches {
			if match == "don't()" {
				flag = true
				continue
			} else if match == "do()" {
				flag = false
				continue
			}
			if !flag{
				res += returnMulsFromStr(match)
			}
		}
	}
	return res
}


func Day3() {
	filepath := os.Getenv("PWD") + "/aoc2024/day3/input.txt"
	inputs, err := utils.InputReader(filepath)
	if err != nil {
		log.Fatalf("Cannot open the file: %v", err)
	}
	println(part1(inputs))
	println(part2(inputs))
}
