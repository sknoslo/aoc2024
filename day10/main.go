package main

import (
	"fmt"
	"sknoslo/aoc2024/utils"
	"sknoslo/aoc2024/vec2"
	"strings"

	"github.com/hashicorp/go-set/v3"
)

var input string

type Grid struct {
	cells []int
	w int
	h int
}

func init() {
	input = utils.MustReadInput("input.txt")
	input = strings.TrimSpace(input)
}

func main() {
	utils.Run(1, partone)
	utils.Run(2, parttwo)
}

func parseInput() Grid {
	var grid Grid

	lines := strings.Split(input, "\n")
	grid.h = len(lines)
	grid.w = len(lines[0])

	grid.cells = make([]int, grid.h * grid.w)

	for y, line := range lines {
		for x, v := range line {
			grid.cells[y * grid.w + x] = int(v - '0')
		}
	}

	return grid
}

func trailScore1(grid Grid, curr, i int, found *set.Set[int]) int {
	if grid.cells[i] - curr != 1 {
		return 0
	}

	if grid.cells[i] == 9 {
		if found.Contains(i) {
			return 0
		}
		found.Insert(i)
		return 1
	}

	score := 0
	pos := vec2.New(i % grid.w, i / grid.w)
	for _, dir := range vec2.CardinalDirs {
		n := pos.Add(dir)
		if n.InRange(0, 0, grid.w-1, grid.h-1) {
			ni := n.Y * grid.w + n.X
			score += trailScore1(grid, grid.cells[i], ni, found)
		}
	}

	return score
}

func partone() string {
	grid := parseInput()

	sum := 0

	for i, v := range grid.cells {
		if v == 0 {
			sum += trailScore1(grid, -1, i, set.New[int](4))
		}
	}

	return fmt.Sprint(sum)
}

func trailScore2(grid Grid, curr, i int) int {
	if grid.cells[i] - curr != 1 {
		return 0
	}

	if grid.cells[i] == 9 {
		return 1
	}


	score := 0
	pos := vec2.New(i % grid.w, i / grid.w)
	for _, dir := range vec2.CardinalDirs {
		n := pos.Add(dir)
		if n.InRange(0, 0, grid.w-1, grid.h-1) {
			ni := n.Y * grid.w + n.X
			score += trailScore2(grid, grid.cells[i], ni)
		}
	}

	return score
}

func parttwo() string {
	grid := parseInput()

	sum := 0

	for i, v := range grid.cells {
		if v == 0 {
			sum += trailScore2(grid, -1, i)
		}
	}

	return fmt.Sprint(sum)
}
