package main

import (
	"fmt"
	"log"
	"sknoslo/aoc2024/utils"
	"strings"
)

var input string

func init() {
	in, err := utils.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(in)
}

func main() {
	utils.Run(1, partone)
	utils.Run(2, parttwo)
}

func parseInput() [][]int {
	lines := strings.Split(input, "\n")
	out := make([][]int, len(lines))
	for i, line := range lines {
		levels, err := utils.SplitInts(line)
		if err != nil {
			log.Fatal(err)
		}
		out[i] = levels
	}
	return out
}

func remove(report []int, index int) []int {
	out := make([]int, 0, len(report)-1)
	out = append(out, report[:index]...)
	return append(out, report[index+1:]...)
}

func isSafe(report []int) bool {
	inc := report[0] < report[1]

	for i := 1; i < len(report); i++ {
		diff := utils.AbsDiff(report[i-1], report[i])
		if diff == 0 || diff > 3 || (report[i-1] < report[i]) != inc {
			return false
		}
	}

	return true
}

func partone() string {
	reports := parseInput()

	safe := 0

	for _, report := range reports {
		if isSafe(report) {
			safe++
		}
	}

	return fmt.Sprint(safe)
}

func parttwo() string {
	reports := parseInput()

	safe := 0

	for _, report := range reports {
		if isSafe(report) {
			safe++
		} else {
			for i := 0; i < len(report); i++ {
				newReport := remove(report, i)
				if isSafe(newReport) {
					safe++
					break
				}
			}
		}
	}

	return fmt.Sprint(safe)
}
