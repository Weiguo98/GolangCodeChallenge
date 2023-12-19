package aoc2023

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day18() {
	filepath := os.Getenv("PWD") + "/aoc2023/day18/input.txt"
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Cannot open the file: %v", err)
	}
	defer f.Close()

	digMap := make(map[int]int, 0)

	scanner := bufio.NewScanner(f)
	count := 0
	rowLen := 1
	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), " ")
		direction := strs[0]
		num, _ := strconv.Atoi(strs[1])

		fmt.Printf("direction: %s, num: %d\n", direction, num) // Print direction and num for debugging

		if direction == "R" {
			if val, ok := digMap[count]; ok {
				digMap[count] = val + num
			} else {
				digMap[count] = num + rowLen
			}
			rowLen = digMap[count]
			fmt.Printf("R: count: %d, rowLen: %d\n", count, rowLen) // Print count and rowLen for debugging
		} else if direction == "L" {
			if val, ok := digMap[count]; ok {
				digMap[count] = val - num + num 
				rowLen = val - num
				fmt.Printf("%v", digMap[count])
			} else {
				fmt.Print("wrong")
			}
			fmt.Printf("L: count: %d, rowLen: %d\n", count, rowLen) // Print count and rowLen for debugging
		} else if direction == "D" {
			count++
			for i := count; i < count+num; i++ {
				digMap[i] = rowLen
			}
			count = count + num - 1
			fmt.Printf("D: count: %d, digmap:%v\n", count, digMap) // Print count for debugging
		} else if direction == "U" {
			for i := count; i > count-num; i-- {
				if val, ok := digMap[i]; ok {
					digMap[i] = val - rowLen + 1
				} else {
					fmt.Print("wrong")
				}
			}
			count = count - num + 1
			fmt.Printf("U: count: %d, digmap:%v\n", count, digMap) // Print count for debugging
		}
	}
	result := 0
	for index, count := range digMap {
		fmt.Printf("index: %d, count: %d\n", index, count)
		result += count + 1
	}
	fmt.Println(result)
}
