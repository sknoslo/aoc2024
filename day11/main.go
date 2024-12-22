package main

import (
	"fmt"
	"sknoslo/aoc2024/utils"
	"strconv"
	"strings"
)

var input string

func init() {
	input = utils.MustReadInput("input.txt")
	input = strings.TrimSpace(input)
}

func main() {
	utils.Run(1, partone)
	utils.Run(2, parttwo)
}

type Memory struct {
	s     string
	steps int
}

var memo map[Memory]int

func step(s string, steps int) int {
	if memo == nil {
		memo = make(map[Memory]int, 256)
	}

	if steps == 0 {
		return 1
	}

	key := Memory{s, steps}
	if v, ok := memo[key]; ok {
		return v
	}

	if s == "0" {
		memo[key] = step("1", steps-1)
	} else if len(s)%2 == 0 {
		h := len(s) / 2

		left := s[:h]
		right := strconv.Itoa(utils.MustAtoi(s[h:])) // strip leading zeros
		memo[key] = step(left, steps-1) + step(right, steps-1)
	} else {
		memo[key] = step(strconv.Itoa(utils.MustAtoi(s)*2024), steps-1)
	}

	return memo[key]
}

func partone() string {
	stones := strings.Fields(input)

	count := 0

	for _, s := range stones {
		count += step(s, 25)
	}

	return fmt.Sprint(count)
}

func parttwo() string {
	stones := strings.Fields(input)

	count := 0

	for _, s := range stones {
		count += step(s, 75)
	}

	return fmt.Sprint(count)
}
