package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type coordinates struct {
	x float64
	y float64
}

func main() {
	var gopher, dog coordinates
	fmt.Scanf("%f %f %f %f", &gopher.x, &gopher.y,
		&dog.x, &dog.y)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := strings.Split(scanner.Text(), " ")
		var hole coordinates
		hole.x, _ = strconv.ParseFloat(str[0], 32)
		hole.y, _ = strconv.ParseFloat(str[1], 32)
		gopherDis := math.Sqrt((gopher.x-hole.x)*(gopher.x-hole.x) + (gopher.y-hole.y)*(gopher.y-hole.y))
		dogDis := math.Sqrt((dog.x-hole.x)*(dog.x-hole.x) + (dog.y-hole.y)*(dog.y-hole.y))
		if dogDis >= gopherDis*2 {
			fmt.Printf("The gopher can escape through the hole at (%.3f,%.3f).\n", hole.x, hole.y)
			return
		}
	}
	fmt.Println("The gopher cannot escape.")
}
