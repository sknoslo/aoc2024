package main

import (
	"fmt"
	"sknoslo/aoc2024/utils"
	"sknoslo/aoc2024/vec2"
	"slices"
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

func parseInput() (map[vec2.Vec2]rune, vec2.Vec2, string) {
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

	strgrid := strings.Join(lines, "")
	strgrid = strings.Replace(strgrid, "^", ".", 1)

	return grid, start, strgrid
}

func partone() string {
	grid, pos, _ := parseInput()
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

var dlseen [][4]bool

func detectLoop(pos, dir vec2.Vec2, grid string) bool {
	if len(dlseen) == 0 {
		dlseen = make([][4]bool, w*h)
	} else {
		clear(dlseen)
	}

	barrier := pos.Add(dir)

	for pos.InRange(0, 0, w-1, h-1) {
		di := slices.Index(vec2.CardinalDirs, dir)
		if dlseen[pos.Y*w+pos.X][di] {
			return true
		}
		dlseen[pos.Y*w+pos.X][di] = true
		n := pos.Add(dir)
		if n.InRange(0, 0, w-1, h-1) {
			tile := grid[n.Y*w+n.X]
			if tile != '.' || n == barrier {
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
	_, pos, grid := parseInput()
	dir := vec2.North
	seen := make([]bool, w*h)

	loops := 0

	for pos.InRange(0, 0, w-1, h-1) {
		seen[pos.Y*w+pos.X] = true
		n := pos.Add(dir)
		if n.InRange(0, 0, w-1, h-1) {
			tile := grid[n.Y*w+n.X]
			if tile != '.' {
				dir = dir.RotateCardinalCW()
			} else {
				next := pos.Add(dir)
				if !seen[next.Y*w+next.X] && detectLoop(pos, dir, grid) {
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
