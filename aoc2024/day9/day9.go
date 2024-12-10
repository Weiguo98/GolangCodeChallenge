package day9

import (
	utils "adventOfCode/aoc2024/utils"
	"fmt"
	"log"
	"os"
	"strconv"
)

func changeDiskMapToFiles(diskMap string) []string {
	files := []string{}
	ID := 0
	for i, letter := range diskMap {
		num, _ := strconv.Atoi(string(letter))
		if i%2 == 0 {
			for j := 0; j < num; j++ {
				files = append(files, strconv.Itoa(ID))
			}
			ID++
		} else {
			for j := 0; j < num; j++ {
				files = append(files, ".")
			}
		}
	}
	return files
}

func moveWholeFileToEmptyPlace(file string, files []string) []string {
	//fmt.Print(file)
	lastNum := len(files) - 1
	for ; lastNum >= 0; lastNum-- {
		if files[lastNum] == file {
			break
		}
	}
	//print(lastNum)
	count := 0
	for i := lastNum; i >= 0; i-- {
		if files[i] == files[lastNum] {
			count++
		} else {
			break
		}
	}
	//fmt.Println(lastNum, count)
	start := -1
	emptySpaceCount := 0
	for i := 0; i < len(files); i++ {
		if i > lastNum {
			break
		}
		if files[i] == "." && start == -1 {
			start = i
		}
		if files[i] == "." {
			emptySpaceCount++
		} else {
			emptySpaceCount = 0
			start = -1
		}
		if emptySpaceCount == count {
			for j := start; j < start+count; j++ {
				files[j] = files[lastNum]
				files[lastNum] = "."
				lastNum--
			}
			break
		}
	}
	//fmt.Println(files)
	return files
}

func part1(files []string) {
	j := len(files) - 1
	for i, letter := range files {
		if i >= j {
			break
		}
		if letter == "." {
			for files[j] == "." {
				j--
			}
			files[i] = files[j]
			files[j] = "."
			j--
		}
	}
}

// 0.738s
func Day() {
	filepath := os.Getenv("PWD") + "/aoc2024/day9/input.txt"
	inputs, err := utils.InputReader(filepath)
	if err != nil {
		log.Fatalf("Cannot open the file: %v", err)
	}
	diskMap := inputs[0]
	files := changeDiskMapToFiles(diskMap)
	filesDeepCopy := make([]string, len(files))
	copy(filesDeepCopy, files)
	lastNum := ""
	for i := len(filesDeepCopy) - 1; i >= 0; i-- {
		if string(filesDeepCopy[i]) != "." && string(filesDeepCopy[i]) != lastNum {
			files = moveWholeFileToEmptyPlace(string(filesDeepCopy[i]), files)
			lastNum = string(filesDeepCopy[i])
		}
	}

	res := 0
	for i := 0; i < len(files); i++ {
		if files[i] != "." {
			num, _ := strconv.Atoi(string(files[i]))
			res += num * i
		}
	}

	fmt.Println(res)
}
