package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/joho/godotenv"
)
func main() {

	  err := godotenv.Load(".env")
	  filepath := os.Getenv("PWD")+"\\Day1\\input.txt"
	  f, err := os.Open(filepath)
	  if err != nil {
		  log.Fatalf("Cannot open the file: %v",err)
	  }
	  defer f.Close()

	  calories := make([]int, 0)
	  food := 0

	  scanner := bufio.NewScanner(f)
	  for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			calories = append(calories, food)
			food = 0
		} else {
			tempFood, _ := strconv.Atoi(scanner.Text())
			food = food + tempFood
		}
	  }

	  sort.Ints(calories)

	  //question1:
	  fmt.Println(calories[len(calories)-1])
	  // question2:
	  fmt.Println(calories[len(calories)-1] + calories[len(calories)-2] + calories[len(calories)-3])

	  if err := scanner.Err(); err != nil {
		  log.Fatalf("Cannot scan from file: %v", err)
	  }
}