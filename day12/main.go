package main

import (
	"fmt"
	"sknoslo/aoc2024/grids"
	"sknoslo/aoc2024/stacks"
	"sknoslo/aoc2024/utils"
	"sknoslo/aoc2024/vec2"

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

func partone() string {
	grid := grids.FromRunes(input)
	sum := 0
	seen := set.New[vec2.Vec2](64)
	toVisit := stacks.New[vec2.Vec2](64)

	for start, plant := range grid.Cells() {
		if seen.Contains(start) {
			continue
		}

		toVisit.Clear()
		toVisit.Push(start)

		area := 0
		edges := 0

		for !toVisit.Empty() {
			plot := toVisit.Pop()
			if seen.Contains(plot) {
				continue
			}
			seen.Insert(plot)

			area++

			for _, dir := range vec2.CardinalDirs {
				n := plot.Add(dir)

				if grid.InGrid(n) && grid.CellAt(n) == plant {
					toVisit.Push(n)
				} else {
					edges++
				}
			}
		}

		sum += area * edges
	}

	return fmt.Sprint(sum)
}

type Step struct {
	plot, dir vec2.Vec2
}

func countSides(region *set.Set[vec2.Vec2]) int {
	sides := 0

	var mn, mx vec2.Vec2

	for plot := range region.Items() {
		if plot.X < mn.X {
			mn.X = plot.X
		} else if plot.X > mx.X {
			mx.X = plot.X
		}
		if plot.Y < mn.Y {
			mn.Y = plot.Y
		} else if plot.Y > mx.Y {
			mx.Y = plot.Y
		}
	}

	for y := mn.Y; y <= mx.Y; y++ {
		on, top, bottom := false, false, false

		for x := mn.X; x <= mx.X; x++ {
			p := vec2.New(x, y)
			nextOn, nextTop, nextBottom := region.Contains(p), region.Contains(p.Add(vec2.North)), region.Contains(p.Add(vec2.South))
			if nextOn && !nextTop && (!on || top) {
				sides++
			}
			if nextOn && !nextBottom && (!on || bottom) {
				sides++
			}
			on, top, bottom = nextOn, nextTop, nextBottom

		}
	}

	for x := mn.X; x <= mx.X; x++ {
		on, left, right := false, false, false

		for y := mn.Y; y <= mx.Y; y++ {
			p := vec2.New(x, y)
			nextOn, nextLeft, nextRight := region.Contains(p), region.Contains(p.Add(vec2.West)), region.Contains(p.Add(vec2.East))
			if nextOn && !nextLeft && (!on || left) {
				sides++
			}
			if nextOn && !nextRight && (!on || right) {
				sides++
			}
			on, left, right = nextOn, nextLeft, nextRight
		}
	}

	return sides
}

func parttwo() string {
	grid := grids.FromRunes(input)
	sum := 0
	seen := set.New[vec2.Vec2](64)
	toVisit := stacks.New[vec2.Vec2](64)

	for start, plant := range grid.Cells() {
		if seen.Contains(start) {
			continue
		}

		toVisit.Clear()
		toVisit.Push(start)

		region := set.New[vec2.Vec2](64)

		for !toVisit.Empty() {
			plot := toVisit.Pop()
			if seen.Contains(plot) {
				continue
			}
			seen.Insert(plot)
			region.Insert(plot)

			for _, dir := range vec2.CardinalDirs {
				n := plot.Add(dir)

				if grid.InGrid(n) && grid.CellAt(n) == plant {
					toVisit.Push(n)
				}
			}
		}

		area := region.Size()
		sides := countSides(region)
		sum += area * sides
	}

	return fmt.Sprint(sum)
}
