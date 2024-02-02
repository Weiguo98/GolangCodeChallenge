package aoc2023

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type rules struct {
	preRules []string
	end      string
}

func Day19() {
	filepath := os.Getenv("PWD") + "/aoc2023/day19/input.txt"
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Cannot open the file: %v", err)
	}
	defer f.Close()
	strRuleSets := []string{}
	strSystem := []string{}

	emptyLine := false

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() == "" {
			emptyLine = true
			continue
		}
		if !emptyLine {
			strRuleSets = append(strRuleSets, scanner.Text())
		} else {
			strSystem = append(strSystem, scanner.Text())
		}
	}
	systems := parseSystem(strSystem)
	ruleSets := parseRuleSet(strRuleSets)

	// part1(systems, ruleSets)
	part2(systems, ruleSets)
}

func part1(systems []map[string]int, ruleSets map[string]rules) {
	result := 0
	for _, system := range systems {
		key := "in"
		fmt.Println(system)
		for {
			// fmt.Println(ruleSets[key])
			key = applyRuleSet(ruleSets[key], system)
			fmt.Println(key)
			if key == "R" {
				break
			} else if key == "A" {
				result += system["x"] + system["m"] + system["a"] + system["s"]
				break
			}
		}
	}
	fmt.Println(result)
}

func part2(systems []map[string]int, ruleSets map[string]rules) {
	key := "in"
	resultRuleSet := []string{}
	combinations := [][]string{}
	applyRuleSetRecursive(key, ruleSets, resultRuleSet, &combinations)
	fmt.Println(combinations)
}

func parseSystem(systems []string) []map[string]int {
	systemsMap := []map[string]int{}
	for _, strSystem := range systems {
		strSystem = strings.Trim(strSystem, "}")
		strSystem = strings.Trim(strSystem, "{")
		strs := strings.Split(strSystem, ",")
		system := make(map[string]int, 4)
		for _, str := range strs {
			num, _ := strconv.Atoi(str[2:])
			system[string(str[0])] = num
		}
		systemsMap = append(systemsMap, system)
	}
	return systemsMap
}

func parseRuleSet(ruleSets []string) map[string]rules {
	ruleSetMap := map[string]rules{}
	for _, strRuleSet := range ruleSets {
		strs := strings.Split(strRuleSet, "{")
		ruleName := string(strs[0])
		rules := rules{}
		strRules := strings.Split(strs[1], ",")
		for i := 0; i < len(strRules)-1; i++ {
			rules.preRules = append(rules.preRules, strRules[i])
		}
		rules.end = strings.Trim(strRules[len(strRules)-1], "}")
		ruleSetMap[ruleName] = rules
	}
	return ruleSetMap
}
var memo = make(map[string][][]string)
func applyRuleSetRecursive(rule string, ruleSets map[string]rules, resultRuleSet []string, combinations *[][]string) {
	// If the result is in the memoization map, return it
    if _, ok := memo[rule]; ok {
        return 
    }
	if rule == "R" {
		return
	} else if rule == "A" {
		*combinations = append(*combinations, resultRuleSet)
		return
	}

	ruleSet := ruleSets[rule]

	for _, preRule := range ruleSet.preRules {
		key, condition:= applyPreRule(preRule)
		applyRuleSetRecursive(key, ruleSets, append(resultRuleSet, condition), combinations)
	}
	applyRuleSetRecursive(ruleSet.end, ruleSets, append(resultRuleSet, ruleSet.preRules...), combinations)
	memo[rule]= *combinations
}

func applyPreRule(rule string) (string, string) {
	preRules := strings.Split(rule, ":")
	return preRules[1], preRules[0]
}

func applyRuleSet(rules rules, system map[string]int) string {
	for _, preRule := range rules.preRules {
		preRules := strings.Split(preRule, ":")
		rule := preRules[0]
		end := preRules[1]
		if applyRule(rule, system) {
			return end
		}
	}
	return rules.end
}

func applyRule(rule string, system map[string]int) bool {
	key := string(rule[0])
	operations := string(rule[1])
	num, _ := strconv.Atoi(string(rule[2:]))
	if operations == "<" {
		return system[key] < num
	} else if operations == "=" {
		return system[key] == num
	} else if operations == ">" {
		return system[key] > num
	}
	return false
}

// import re

// # Parse the input
// workflows = {}
// with open('input.txt', 'r') as file:
//     for line in file:
//         name, conditions, next_workflows = re.match(r'(\w+)\{(.+):(.+)\}', line.strip()).groups()
//         conditions = conditions.split(',')
//         next_workflows = next_workflows.split(',')
//         workflows[name] = {'conditions': conditions, 'next': next_workflows}

// # Initialize the memoization table
// memo = {}

// def dfs(workflow, ratings):
//     # If the result is in the memoization table, return it
//     if (workflow, ratings) in memo:
//         return memo[(workflow, ratings)]

//     # If the current workflow is 'A', return 1
//     if workflow == 'A':
//         return 1

//     # If the current workflow is 'R', return 0
//     if workflow == 'R':
//         return 0

//     # Check the conditions for the current workflow
//     conditions = workflows[workflow]['conditions']
//     for condition in conditions:
//         rating, operator, value = re.match(r'(\w+)([<>])(\d+)', condition).groups()
//         value = int(value)
//         if operator == '<' and ratings[rating] >= value or operator == '>' and ratings[rating] <= value:
//             return 0

//     # Recursively traverse the next workflows
//     next_workflows = workflows[workflow]['next']
//     count = 0
//     for next_workflow in next_workflows:
//         for rating in ['x', 'm', 'a', 's']:
//             for value in range(1, 4000):
//                 new_ratings = ratings.copy()
//                 new_ratings[rating] = value
//                 count += dfs(next_workflow, new_ratings)

//     # Store the result in the memoization table
//     memo[(workflow, ratings)] = count

//     return count

// # Perform the DFS from the 'in' node
// ratings = {'x': 0, 'm': 0, 'a': 0, 's': 0}
// result = dfs('in', ratings)

// print(result)