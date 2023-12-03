package main

import (
	"bufio"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"strings"
)

type assignment struct {
	left  int
	right int
}

func main() {

	err := godotenv.Load(".env")
	filepath := os.Getenv("PWD") + "/AOC2022-Day4/input.txt"
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Cannot open the file: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var first, second assignment
	firstCount := 0
	secondCount := 0

	for scanner.Scan() {
		str := strings.Split(scanner.Text(), ",")
		temp := strings.Split(str[0], "-")
		first.left, _ = strconv.Atoi(temp[0])
		first.right, _ = strconv.Atoi(temp[1])
		temp = strings.Split(str[1], "-")
		second.left, _ = strconv.Atoi(temp[0])
		second.right, _ = strconv.Atoi(temp[1])

		if first.left <= second.left && first.right >= second.right {
			firstCount++
		} else if second.left <= first.left && second.right >= first.right {
			firstCount++
		}

		if first.left > second.right || first.right < second.left {
			continue
		} else {
			secondCount++
		}
	}
	fmt.Println(firstCount)
	fmt.Println(secondCount)
}
