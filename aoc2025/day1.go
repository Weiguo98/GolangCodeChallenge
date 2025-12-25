package aoc2025

import (
	utils "adventOfCode/utils"
	"log"
	"math"
	"os"
	"strconv"
)

// part1
func countLeft(start, num int) int {
	if start >= num {
		return start - num
	}
	if (num-start)%100 == 0 {
		return 0
	}
	return 100 - (num-start)%100
}

func countRight(start, num int) int {
	if 99-start >= num {
		return start + num
	}
	return (num + start - 100) % 100
}

// part2
func countLeft2(start, num int) (int, int) {
	if start == 0 {
		return (num - start) / 100, 100 - (num-start)%100
	}
	if start > num {
		return 0, start - num
	}
	if (num-start)%100 == 0 {
		return (num-start)/100 + 1, 0
	}
	return (num-start)/100 + 1, 100 - (num-start)%100
}

func countRight2(start, num int) (int, int) {
	if 99-start >= num {
		return 0, start + num
	}
	return (num + start) / 100, (num + start - 100) % 100
}

func Day1() {
	filepath := os.Getenv("PWD") + "/aoc2025/day1/input.txt"
	inputs, err := utils.InputReader(filepath)
	if err != nil {
		log.Fatalf("Cannot open the file: %v", err)
	}
	// part 1
	dial := 50
	res := 0
	for _, line := range inputs {
		direction := line[0]
		nums, err := strconv.Atoi(string(line[1:]))
		if err != nil {
			log.Fatalf("Cannot convert the string to int: %v", err)
		}
		switch direction {
		case 'L':
			dial = countLeft(dial, nums)
		case 'R':
			dial = countRight(dial, nums)
		default:
			log.Fatalf("Invalid direction: %v", direction)
		}
		if dial == 0 {
			res++
		}
	}
	println(res)
	res = 0
	dial = 50
	// part 2
	for _, line := range inputs {
		direction := line[0]
		nums, err := strconv.Atoi(string(line[1:]))
		if err != nil {
			log.Fatalf("Cannot convert the string to int: %v", err)
		}
		times := 0
		switch direction {
		case 'L':
			times, dial = countLeft2(dial, nums)
		case 'R':
			times, dial = countRight2(dial, nums)
		default:
			log.Fatalf("Invalid direction: %v", direction)
		}
		res += int(math.Abs(float64(times)))
	}
	println(res)
}
