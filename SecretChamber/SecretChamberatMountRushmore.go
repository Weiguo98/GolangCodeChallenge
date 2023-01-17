package main

import (
	"fmt"
)

type pair struct {
	left  string
	right string
}

type transPairRes struct {
	left  string
	right []string
}

func getTranslateRes(trans []pair, visited []bool, target string, res *[]string) {
	for index, t := range trans {
		if t.left == target && visited[index] == false {
			visited[index] = true
			*res = append(*res, t.right)
			getTranslateRes(trans, visited, t.right, res)
		}
	}
	return
}

func foundInTranslation(right byte, target []string) bool {
	for _, letter := range target {
		if string(right) == letter {
			return true
		}
	}
	// fmt.Printf("Cannot found target %s in the trans array %v", string(right), target)
	return false
}

func compareWords(leftWord, rightWord string, transMap map[string][]string) bool {
	for i := 0; i < len(leftWord); i++ {
		if leftWord[i] == rightWord[i] {
			continue
		}
		if transForLeft, ok := transMap[string(leftWord[i])]; ok {
			if !foundInTranslation(rightWord[i], transForLeft) {
				return false
			}
		} else {
			// fmt.Printf("Cannot get %s in the transmap\n", string(leftWord[i]))
			return false
		}
	}
	return true
}
func main() {

	var transNum, pairNum int
	fmt.Scanf("%d %d", &transNum, &pairNum)
	transPairs := make([]pair, transNum)
	wordPairs := make([]pair, pairNum)
	visited := make([]bool, transNum)
	transRes := make([]transPairRes, transNum)

	for i := 0; i < transNum; i++ {
		fmt.Scanf("%s %s", &transPairs[i].left, &transPairs[i].right)
	}

	for i := 0; i < pairNum; i++ {
		fmt.Scanf("%s %s", &wordPairs[i].left, &wordPairs[i].right)
	}

	for i := 0; i < transNum; i++ {
		transRes[i].left = transPairs[i].left
		visited[i] = true
		transRes[i].right = append(transRes[i].right, transPairs[i].right)
		getTranslateRes(transPairs, visited, transPairs[i].right, &transRes[i].right)
		visited = make([]bool, transNum)
	}
	fmt.Println(transRes)
	transMap := make(map[string][]string)
	for j := 0; j < len(transRes); j++ {
		item := transRes[j]
		for i := 0; i < len(item.right); i++ {
			transMap[item.left] = append(transMap[item.left], item.right[i])
		}
	}
	fmt.Println(transMap)
	res := make([]string, 0)
	for _, word := range wordPairs {
		if len(word.left) != len(word.right) {
			res = append(res, "no")
			continue
		}
		if compareWords(word.left, word.right, transMap) {
			res = append(res, "yes")
		} else {
			res = append(res, "no")
		}
	}

	for j := 0; j < len(res); j++ {
		fmt.Println(res[j])
	}
}
