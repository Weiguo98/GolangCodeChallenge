package aoc2024

import (
	utils "adventOfCode/aoc2024/utils"
	_"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day3() {
	filepath := os.Getenv("PWD") + "/aoc2024/day3/input.txt"
	inputs, err := utils.InputReader(filepath)
	if err != nil {
		log.Fatalf("Cannot open the file: %v", err)
	}
	oneResult := 0
	for _,input := range inputs{
		re := regexp.MustCompile(`mul\((\d+),\s*(\d+)\)`)
		matches := re.FindAllString(input, -1)
		for _, match := range matches{
			str := strings.Split(match, "(")[1]
			str = strings.Split(str, ")")[0]
			strs := strings.Split(str, ",")
			num1, _ := strconv.Atoi(strs[0])
            num2, _ := strconv.Atoi(strs[1])
            oneResult += num1 * num2
		}
	}
	println(oneResult)
}
