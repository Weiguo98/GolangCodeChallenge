package day18

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day() {
	filepath := "aoc2024/day18/input.txt"
	if cwd, err := os.Getwd(); err == nil {
		filepath = cwd + "/" + filepath
	}
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Cannot open the file: %v", err)
	}
	defer file.Close()

	memorySpace := make([][]bool, 71)
	for i := range memorySpace {
		memorySpace[i] = make([]bool, 71)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		strInts := strings.Split(line, ",")
		if len(strInts) != 2 {
			continue
		}
		x, _ := strconv.Atoi(strInts[0])
		y, _ := strconv.Atoi(strInts[1])
		if x >= 0 && x < 71 && y >= 0 && y < 71 {
			memorySpace[x][y] = true
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading the file: %v", err)
	}

	steps := bfs(memorySpace, 71, 71)
	fmt.Println("Minimum steps to reach the exit:", steps)
}

func bfs(memorySpace [][]bool, rows, cols int) int {
	type Point struct {
		x, y, steps int
	}
	directions := []Point{{0, 1, 0}, {1, 0, 0}, {0, -1, 0}, {-1, 0, 0}}
	queue := []Point{{0, 0, 0}}
	visited := make(map[Point]bool)
	visited[Point{0, 0, 0}] = true

	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]

		if point.x == rows-1 && point.y == cols-1 {
			return point.steps
		}

		for _, dir := range directions {
			newX, newY := point.x+dir.x, point.y+dir.y
			if newX >= 0 && newX < rows && newY >= 0 && newY < cols && !memorySpace[newX][newY] && !visited[Point{newX, newY, 0}] {
				queue = append(queue, Point{newX, newY, point.steps + 1})
				visited[Point{newX, newY, 0}] = true
			}
		}
	}

	return -1 // Return -1 if no path is found
}
