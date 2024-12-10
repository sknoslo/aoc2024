package main

import (
	"fmt"
	"sknoslo/aoc2024/utils"
	"sknoslo/aoc2024/vec2"
	"sknoslo/aoc2024/grid"
	"strings"

	"github.com/hashicorp/go-set/v3"
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

func trailScore1(grid *grid.Grid[int], curr int, i vec2.Vec2, found *set.Set[vec2.Vec2]) int {
	v := grid.CellAt(i)
	if v - curr != 1 {
		return 0
	}

	if v == 9 {
		if found.Contains(i) {
			return 0
		}
		found.Insert(i)
		return 1
	}

	score := 0
	for _, dir := range vec2.CardinalDirs {
		n := i.Add(dir)
		if grid.InGrid(n) {
			score += trailScore1(grid, v, n, found)
		}
	}

	return score
}

func partone() string {
	grid := grid.MustFromDigits(input)

	sum := 0

	for i, v := range grid.Cells() {
		if v == 0 {
			sum += trailScore1(grid, -1, i, set.New[vec2.Vec2](4))
		}
	}

	return fmt.Sprint(sum)
}

func trailScore2(grid *grid.Grid[int], curr int, i vec2.Vec2) int {
	v := grid.CellAt(i)
	if v - curr != 1 {
		return 0
	}

	if v == 9 {
		return 1
	}


	score := 0
	for _, dir := range vec2.CardinalDirs {
		n := i.Add(dir)
		if grid.InGrid(n) {
			score += trailScore2(grid, v, n)
		}
	}

	return score
}

func parttwo() string {
	grid := grid.MustFromDigits(input)

	sum := 0

	for i, v := range grid.Cells() {
		if v == 0 {
			sum += trailScore2(grid, -1, i)
		}
	}

	return fmt.Sprint(sum)
}
