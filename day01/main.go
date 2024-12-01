package main

import (
	"log"
	"sknoslo/aoc2024/utils"
	"slices"
	"strconv"
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

func atoi(in string) int {
	i, err := strconv.Atoi(in)
	if err != nil {
		log.Fatal(err)
	}

	return i
}

func parseList() ([]int, []int) {
	lines := strings.Split(input, "\n")
	l, r := make([]int, len(lines)), make([]int, len(lines))

	for i, line := range lines {
		nums := strings.Fields(line)

		l[i] = atoi(nums[0])
		r[i] = atoi(nums[1])
	}

	return l, r
}

func absDiff(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}

func partone() string {
	l, r := parseList()

	slices.Sort(l)
	slices.Sort(r)

	sum := 0

	for i, a := range l {
		b := r[i]

		sum += absDiff(a, b)
	}

	return strconv.Itoa(sum)
}

func parttwo() string {
	l, r := parseList()

	rc := make(map[int]int, len(r))

	for _, v := range r {
		rc[v]++
	}

	sum := 0
	for _, v := range l {
		if c, ok := rc[v]; ok {
			sum += v * c
		}
	}

	return strconv.Itoa(sum)
}
