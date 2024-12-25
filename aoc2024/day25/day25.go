package day25

import (
	utils "adventOfCode/aoc2024/utils"
	"fmt"
	"log"
	"os"
)

func Day() {
	filepath := os.Getenv("PWD") + "/aoc2024/day25/input.txt"
	lock := [][5]int{{0, 0, 0, 0, 0}}
	key := [][5]int{{0, 0, 0, 0, 0}}
	inputs, err := utils.InputReader(filepath)
	if err != nil {
		log.Fatalf("Cannot open the file: %v", err)
	}
	keyCount := 0
	lockCount := 0
	isLock := false
	for i, input := range inputs {
		if input == "" {
			isLock = false
			continue
		} else if i%8 == 0 && input == "#####" {
			isLock = true
			lock = append(lock, [5]int{0, 0, 0, 0, 0})
			lockCount++
			continue
		} else if i%8 == 6 && input == "#####" {
			continue
		} else if i%8 == 0 && input == "....." {
			isLock = false
			key = append(key, [5]int{0, 0, 0, 0, 0})
			keyCount++
			continue
		} else if i%8 == 6 && input == "....." {
			continue
		}


		if isLock {
			for j, c := range input {
				if string(c) == "#" {
					lock[lockCount][j] += 1
				}
			}
		} else {
			for j, c := range input {
				if string(c) == "#" {
					key[keyCount][j] += 1
				}
			}
		}
	}

	res := 0
	for _, lock := range lock[1:] {
		for _, key := range key[1:] {
			flag := true
			for i := 0; i < 5; i++ {
				if lock[i]+key[i] > 5 {
					flag = false
				}
			}
			if flag {
				res += 1
			}
		}
	}
	fmt.Println(res)
}
