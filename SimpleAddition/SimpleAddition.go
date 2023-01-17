package main

import (
	"fmt"
	// "strings"
	"bufio"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s1 := scanner.Text()
	scanner.Scan()
	s2 := scanner.Text()

	var length int
	lengthS1 := len([]rune(s1))
	lengthS2 := len([]rune(s2))
	if lengthS1 > lengthS2 {
		length = lengthS1
		for i := lengthS2; i < length; i++ {
			s2 = "0" + s2
		}
	} else {
		length = lengthS2
		for i := lengthS1; i < length; i++ {
			s1 = "0" + s1
		}
	}

	res := ""
	flag := 0
	for i := length - 1; i >= 0; i-- {
		numS1 := int(s1[i] - 48)
		numS2 := int(s2[i] - 48)
		count := numS1 + numS2 + flag
		current := 0
		if count > 9 {
			flag = count / 10
			current = count % 10
		} else {
			flag = 0
			current = count
		}
		str := strconv.Itoa(current)
		res = str + res
	}

	if flag > 0 {
		str := strconv.Itoa(flag)
		res = str + res
	}
	fmt.Print(res)
}
