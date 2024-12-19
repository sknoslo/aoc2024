package main

import (
	"sknoslo/aoc2024/stacks"
	"sknoslo/aoc2024/utils"
	"strconv"
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

type step struct {
	str string
	i int
}

func partone() string {
	towelstr, patternstr, _ := strings.Cut(input, "\n\n")

	towels := strings.Split(towelstr, ", ")
	patterns := strings.Split(patternstr, "\n")
	towelmap := make(map[byte][]string, 5)

	matches := 0

	for _, towel := range towels {
		towelmap[towel[0]] = append(towelmap[towel[0]], towel)
	}

	for _, pattern := range patterns {
		stack := stacks.New[step](len(towels))

		for _, t := range towelmap[pattern[0]] {
			stack.Push(step{ t, 0 })
		}

	stackloop:
		for !stack.Empty() {
			s := stack.Pop()

			for i := 0; i < len(s.str); i++ {
				if i + s.i >= len(pattern) || s.str[i] != pattern[i+s.i] {
					continue stackloop
				}
			}

			if s.i + len(s.str) == len(pattern) {
				matches++
				break stackloop
			}

			ni := s.i + len(s.str)

			for _, t := range towelmap[pattern[ni]] {
				stack.Push(step{ t, ni })
			}
		}
	}

	return strconv.Itoa(matches)
}

func parttwo() string {
	towelstr, patternstr, _ := strings.Cut(input, "\n\n")

	towels := strings.Split(towelstr, ", ")
	patterns := strings.Split(patternstr, "\n")
	towelmap := make(map[byte][]string, 5)

	matches := 0

	for _, towel := range towels {
		towelmap[towel[0]] = append(towelmap[towel[0]], towel)
	}

	for _, pattern := range patterns {
		match := func(i int, str string) bool {
			if i + len(str) > len(pattern) {
				return false
			}

			for j := 0; j < len(str); j++ {
				if str[j] != pattern[i+j] {
					return false
				}
			}

			return true
		}

		c := make(map[int]int, 1024)
		var rec func(int) int
		rec = func (i int) int {
			if v, ok := c[i]; ok {
				return v
			}

			if i == len(pattern) {
				return 1
			}

			sum := 0

			for _, t := range towelmap[pattern[i]] {
				if match(i, t) {
					sum += rec(i + len(t))
				}
			}

			c[i] = sum
			return sum
		}

		matches += rec(0)
	}

	return strconv.Itoa(matches)
}
