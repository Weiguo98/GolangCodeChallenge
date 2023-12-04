package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	filepath := os.Getenv("PWD") + "/AOC2023-Day1/input.txt"
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Cannot open the file: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	result := 0
	for scanner.Scan() {
		nums := make([]int, 0)
		str := scanner.Text()
		first := 0
		for first < len(str) {
			if num, err := strconv.Atoi(string(str[first])); err == nil {
				nums = append(nums, num)
				first++
				continue
			}
			if num, found := foundNumbers(str, first, 3); found {
				nums = append(nums, num)
				first++
				continue
			} else if num, found := foundNumbers(str, first, 4); found {
				nums = append(nums, num)
				first++
				continue
			} else if num, found := foundNumbers(str, first, 5); found {
				nums = append(nums, num)
				first++
				continue
			}
			first++
		}
		fmt.Println(str)
		fmt.Println(nums[0]*10 + nums[len(nums)-1])
		result = result + nums[0]*10 + nums[len(nums)-1]
	}
	fmt.Println(result)
}

func foundNumbers(str string, index int, length int) (int, bool) {
	if index+length > len(str) {
		return 0, false
	}
	switch length {
	case 3:
		return foundThreeLetterNumbers(str[index : index+3])
	case 4:
		return foundFourLetterNumbers(str[index : index+4])
	case 5:
		return foundFiveLetterNumbers(str[index : index+5])
	}
	return 0, false
}

func foundThreeLetterNumbers(str string) (int, bool) {
	switch str {
	case "one":
		return 1, true
	case "two":
		return 2, true
	case "six":
		return 6, true
	}
	return 0, false
}

func foundFourLetterNumbers(str string) (int, bool) {
	switch str {
	case "four":
		return 4, true
	case "five":
		return 5, true
	case "nine":
		return 9, true
	}
	return 0, false
}

func foundFiveLetterNumbers(str string) (int, bool) {
	switch str {
	case "seven":
		return 7, true
	case "three":
		return 3, true
	case "eight":
		return 8, true
	}
	return 0, false
}
