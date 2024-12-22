package main

import (
	"fmt"
	"sknoslo/aoc2024/deques"
	"sknoslo/aoc2024/grids"
	"sknoslo/aoc2024/utils"
	"sknoslo/aoc2024/vec2"
	"strconv"
	"strings"

	"github.com/hashicorp/go-set/v3"
)

var input string

func init() {
	input = utils.MustReadInput("input.txt")
}

func main() {
	utils.Run(1, partone)
	utils.Run(2, parttwo)
}

func parseInput() []vec2.Vec2 {
	lines := strings.Split(input, "\n")
	out := make([]vec2.Vec2, len(lines))

	for i, line := range lines {
		nums := utils.MustSplitIntsSep(line, ",")
		out[i] = vec2.New(nums[0], nums[1])
	}

	return out
}

type step struct {
	p vec2.Vec2
	d int
}

func partone() string {
	size := 70
	sim := 1024
	end := vec2.New(size, size)
	bytes := parseInput()
	g := grids.FromSize(end.X+1, end.Y+1, '.')

	for i := range sim {
		g.SetCellAt(bytes[i], '#')
	}

	q := deques.New[step](71 * 71)
	s := set.New[vec2.Vec2](71 * 71)

	q.PushFront(step{vec2.New(0, 0), 0})

	for !q.Empty() {
		x := q.PopBack()

		if x.p == end {
			return strconv.Itoa(x.d)
		}

		if !g.InGrid(x.p) || g.CellAt(x.p) == '#' || s.Contains(x.p) {
			continue
		}
		s.Insert(x.p)

		for _, dir := range vec2.CardinalDirs {
			q.PushFront(step{x.p.Add(dir), x.d + 1})
		}
	}

	return "no solution"
}

func hasPath(bytes []vec2.Vec2, sim int) bool {
	size := 70
	end := vec2.New(size, size)
	g := grids.FromSize(end.X+1, end.Y+1, '.')
	for i := range sim {
		g.SetCellAt(bytes[i], '#')
		continue
	}

	q := deques.New[step](71 * 71)
	s := grids.FromSize(end.X+1, end.Y+1, false)

	q.PushFront(step{vec2.New(0, 0), 0})

	for !q.Empty() {
		x := q.PopBack()

		if x.p == end {
			return true
		}

		if s.CellAt(x.p) {
			continue
		}
		s.SetCellAt(x.p, true)

		for _, dir := range vec2.CardinalDirs {
			n := x.p.Add(dir)
			if g.InGrid(n) && g.CellAt(n) != '#' {
				q.PushFront(step{x.p.Add(dir), x.d + 1})
			}
		}
	}

	return false
}

func parttwo() string {
	bytes := parseInput()

	start := 0
	end := len(bytes)
	for end != start+1 {
		middle := (end-start)/2 + start
		if hasPath(bytes, middle) {
			start = middle
		} else {
			end = middle
		}
	}

	return fmt.Sprintf("%d,%d", bytes[start].X, bytes[start].Y)
}
