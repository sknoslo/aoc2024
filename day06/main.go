package main

import (
	"fmt"
	"slices"
	"sknoslo/aoc2024/utils"
	"sknoslo/aoc2024/vec2"
	"strings"

	"github.com/hashicorp/go-set/v3"
)

var input string

var w, h int

func init() {
	input = utils.MustReadInput("input.txt")
}

func main() {
	utils.Run(1, partone)
	utils.Run(2, parttwo)
}

func parseInput() (map[vec2.Vec2]rune, vec2.Vec2) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	w, h = len(lines[0]), len(lines)
	grid := make(map[vec2.Vec2]rune, w*h)

	var start vec2.Vec2

	for y, line := range lines {
		for x, r := range line {
			if r == '^' {
				start.X = x
				start.Y = y
				r = '.'
			}
			grid[vec2.New(x, y)] = r
		}
	}

	return grid, start
}

func partone() string {
	grid, pos := parseInput()
	dir := vec2.North
	seen := set.New[vec2.Vec2](100)

	for _, ok := grid[pos]; ok; {
		seen.Insert(pos)
		n := pos.Add(dir)
		if tile, ok := grid[n]; ok {
			if tile != '.' {
				dir = dir.RotateCardinalCW()
			} else {
				pos = pos.Add(dir)
			}
		} else {
			break
		}
	}

	return fmt.Sprint(seen.Size())
}

type Memory struct {
	pos vec2.Vec2
	dir vec2.Vec2
}

func detectLoop(pos, dir vec2.Vec2, grid map[vec2.Vec2]rune) bool {
	seen := make([][4]bool, w*h)

	key := pos.Add(dir)
	grid[key] = 'O'
	defer func() {
		grid[key] = '.'
	}()

	for _, ok := grid[pos]; ok; {
		di := slices.Index(vec2.CardinalDirs, dir)
		if seen[pos.Y * w + pos.X][di] {
			return true
		}
		seen[pos.Y * w + pos.X][di] = true
		n := pos.Add(dir)
		if tile, ok := grid[n]; ok {
			if tile != '.' {
				dir = dir.RotateCardinalCW()
			} else {
				pos = pos.Add(dir)
			}
		} else {
			break
		}
	}

	return false
}

func parttwo() string {
	grid, pos := parseInput()
	dir := vec2.North
	seen := make([]bool, w * h)

	loops := 0

	for _, ok := grid[pos]; ok; {
		seen[pos.Y * w + pos.X] = true
		n := pos.Add(dir)
		if tile, ok := grid[n]; ok {
			if tile != '.' {
				dir = dir.RotateCardinalCW()
			} else {
				next := pos.Add(dir)
				if !seen[next.Y * w + next.X] && detectLoop(pos, dir, grid) {
					loops++
				}
				pos = pos.Add(dir)
			}
		} else {
			break
		}
	}

	return fmt.Sprint(loops)
}
