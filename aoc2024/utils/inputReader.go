package aoc2024

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func InputReader(inputPath string) ([]string, error) {
	filepath := inputPath
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Cannot open the file: %v", err)
	}
	defer f.Close()
	strs := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		strs = append(strs, scanner.Text())
	}
	return strs, err
}

func ConvertStrsToInts(strs []string) ([]int, error) {
	ints := []int{}
	for _, str := range strs {
		num, err := strconv.Atoi(str)
		if err != nil {
			log.Fatalf("Cannot convert the string to float: %v", err)
			return nil, err
		}
		ints = append(ints, num)
	}
	return ints, nil
}