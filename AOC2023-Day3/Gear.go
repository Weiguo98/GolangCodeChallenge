package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	filepath := os.Getenv("PWD") + "/input.txt"
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Cannot open the file: %v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	result := 0
	for scanner.Scan() {
		game := strings.Split(scanner.Text(), ":")
		id := strings.Split(game[0], " ")[1]
		idInt, _ := strconv.Atoi(id)
		fmt.Print(idInt)
		sets := strings.Split(game[1], ";")
		colorMap := make(map[string]int, 3)
		for _, set := range sets {
			colorSets := strings.Split(set, ",")
			for _, colorSet := range colorSets {
				colorSet := strings.TrimSpace(colorSet)
				str := strings.Split(colorSet, " ")
				num, err := strconv.Atoi(str[0])
				if err != nil {
					continue
				}
				color := str[1]
				fmt.Println(color)
				if value, ok := colorMap[color]; ok && num > value {
					colorMap[color] = num
				} else if !ok {
					colorMap[color] = num
				}
			}
		}
		colorRes := 1
		for _, num := range colorMap {
			colorRes = colorRes * num
		}
		result = result + colorRes
	}
	fmt.Print(result)
}
