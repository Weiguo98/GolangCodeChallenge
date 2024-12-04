package aoc2024

import (
	utils "adventOfCode/aoc2024/utils"
	"log"
	"math"
	"os"
	"strings"
)

func Day2(){
	filepath := os.Getenv("PWD") + "/aoc2024/day2/input.txt"
	inputs, err := utils.InputReader(filepath)
	if err != nil {
		log.Fatalf("Cannot open the file: %v", err)
	}
	safeCount := 0
	for _, input := range inputs{
		strs := strings.Split(input, " ")
		reports, err := utils.ConvertStrsToInts(strs)
		if err != nil {
			log.Fatalf("Cannot convert the string to int: %v", err)
		}
		decreasing := false
		if reports[0] > reports[1]{
			decreasing = true
		}
		flag := 0
		// for i := 0; i < len(reports) - 1; i++ {
		// 	if decreasing && reports[i] < reports[i+1] {
		// 		flag = true
		// 		break
		// 	} else if !decreasing && reports[i] > reports[i+1]{
		// 		flag = true
		// 	 	break
		// 	}
		// 	if math.Abs(float64(reports[i] - reports[i+1])) > 3 || math.Abs(float64(reports[i] - reports[i+1])) < 1{
		// 		flag = true
		// 		break
		// 	}
		// }
		for i := 0; i < len(reports) - 1; i++ {
			if decreasing && reports[i] < reports[i+1] {
				flag ++
				continue
			} else if !decreasing && reports[i] > reports[i+1]{
				flag++
			 	continue
			}
			if math.Abs(float64(reports[i] - reports[i+1])) > 3 || math.Abs(float64(reports[i] - reports[i+1])) < 1 {
				if (i + 2) == len(reports){
					flag++
					continue
				} else if math.Abs(float64(reports[i] - reports[i+2])) <= 3 || math.Abs(float64(reports[i] - reports[i+2])) >= 1 {
					flag++
					continue
				}
			}
		}
		if flag <= 1 {
			safeCount++
		}
	}
	println(safeCount)
}