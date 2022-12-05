package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func getScoreFromRes(response string) (score int) {
	switch response {
	case "Y":
		score = 2
	case "X":
		score = 1
	case "Z":
		score = 3
	default:
		score = 0
	}
	return score
}

func getScoreFromGame(opponent, response string) (score int) {

	score = getScoreFromRes(response)

	switch game := opponent + response; game {
	case "AY":
		score += 6
	case "AX":
		score += 3
	case "BZ":
		score += 6
	case "BY":
		score += 3
	case "CX":
		score += 6
	case "CZ":
		score += 3
	}
	return score
}

func main() {

	err := godotenv.Load(".env")
	filepath := os.Getenv("PWD") + "\\Day2\\input.txt"
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Cannot open the file: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	score1 := 0
	score2 := 0

	scoreMap := [3][3]int{{3, 4, 8}, {1, 5, 9}, {2, 6, 7}}
	opponentMap := map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
	}
	endMap := map[string]int{
		"X": 0, //lose + 0
		"Y": 1, // draw + 3
		"Z": 2, // win + 6
	}

	for scanner.Scan() {
		str := strings.Split(scanner.Text(), " ")
		score1 = score1 + getScoreFromGame(str[0], str[1])
		score2 = score2 + scoreMap[opponentMap[str[0]]][endMap[str[1]]]

	}

	//question1
	fmt.Println(score1)

	//question2
	fmt.Println(score2)

}
