package main

import (
	"fmt"
	"github.com/hashicorp/go-set/v3"
	"sknoslo/aoc2024/utils"
	"slices"
	"strings"
)

var input string

func init() {
	input = utils.MustReadInput("input.txt")
}

func main() {
	utils.Run(1, partone)
	utils.Run(2, parttwo)
}

func parseInput() ([][]int, [][]int) {
	var rules [][]int
	var updates [][]int

	parts := strings.Split(strings.TrimSpace(input), "\n\n")

	for _, line := range strings.Split(parts[0], "\n") {
		rule := utils.MustSplitIntsSep(line, "|")
		rules = append(rules, rule)
	}

	for _, line := range strings.Split(parts[1], "\n") {
		update := utils.MustSplitIntsSep(line, ",")
		updates = append(updates, update)
	}

	return rules, updates
}

func partone() string {
	rules, updates := parseInput()
	ruleMap := make(map[int][]int, len(rules))

	for _, rule := range rules {
		ruleMap[rule[0]] = append(ruleMap[rule[0]], rule[1])
	}

	sum := 0

updateLoop:
	for _, update := range updates {
		seen := set.New[int](len(update))
		for _, page := range update {
			seen.Insert(page)
			for _, rule := range ruleMap[page] {
				if seen.Contains(rule) {
					continue updateLoop
				}
			}
		}
		sum += update[len(update)/2]
	}

	return fmt.Sprint(sum)
}

func parttwo() string {
	rules, updates := parseInput()
	ruleMap := make(map[int][]int, len(rules))

	for _, rule := range rules {
		ruleMap[rule[0]] = append(ruleMap[rule[0]], rule[1])
	}

	sum := 0

updateLoop:
	for _, update := range updates {
		seen := set.New[int](len(update))
		for _, page := range update {
			seen.Insert(page)
			for _, rule := range ruleMap[page] {
				if seen.Contains(rule) {
					slices.SortFunc(update, func(a, b int) int {
						if slices.Contains(ruleMap[b], a) {
							return -1
						}
						return 1
					})
					sum += update[len(update)/2]
					continue updateLoop
				}
			}
		}
	}

	return fmt.Sprint(sum)
}
