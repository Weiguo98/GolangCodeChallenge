package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	filepath := os.Getenv("PWD") + "\\Day3\\input.txt"
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Cannot open the file: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	compMap := make(map[rune]int)
	for scanner.Scan() {
		leftMap := make(map[rune]bool)
		str := scanner.Text()
		strLeft := str[:len(str)/2]
		strRight := str[len(str)/2:]
		for _, letter := range strLeft {
			leftMap[letter] = true
		}
		for _, letter := range strRight {
			if _, prs := leftMap[letter]; prs {
				compMap[letter] += 1
				delete(leftMap, letter)
			}
		}
	}

	sum := 0
	for key, value := range compMap {
		fmt.Printf("key is %s, value is %v \n", string(key), value)
		if int(key) > 96 {
			sum = (int(key)-96)*value + sum
		} else {
			sum = (int(key)-38)*value + sum
		}
	}
	fmt.Println(sum)
}
