package main

import (
	"sknoslo/aoc2024/grids"
	"sknoslo/aoc2024/utils"
	"sknoslo/aoc2024/vec2"
	"slices"
	"strconv"
)

var input string

const target = 100

func init() {
	input = utils.MustReadInput("input.txt")
}

func main() {
	utils.Run(1, partone)
	utils.Run(2, parttwo)
}

func partone() string {
	track := grids.FromRunes(input)
	size := track.Size()
	dists := grids.FromSize(size.X, size.Y, -1)
	e := track.Find('E')

	pos := e
	prev := vec2.New(-1, -1)
	dist := 0
	path := make([]vec2.Vec2, 0, 10_000)

	for track.CellAt(pos) != 'S' {
		path = append(path, pos)
		dists.SetCellAt(pos, dist)
		for _, dir := range vec2.CardinalDirs {
			npos := pos.Add(dir)
			if track.CellAt(npos) != '#' && npos != prev {
				prev = pos
				pos = npos
				dist++
				break
			}
		}
	}
	dists.SetCellAt(pos, dist)
	path = append(path, pos)
	slices.Reverse(path)

	cheats := 0

	for _, p0 := range path {
		d0 := dists.CellAt(p0)
		for _, dir := range vec2.CardinalDirs {
			p1 := p0.Add(dir.Mul(2))
			if dists.InGrid(p1) {
				d1 := dists.CellAt(p1)

				if d1 != -1 && d0-d1-2 >= target {
					cheats++
				}
			}
		}
	}
	return strconv.Itoa(cheats)
}

type step struct {
	p vec2.Vec2
	d int
}

func findCheats(dists *grids.Grid[int], p0 vec2.Vec2) int {
	cheats := 0

	d0 := dists.CellAt(p0)
	
	for y := p0.Y - 20; y <= p0.Y + 20; y++ {
		m := utils.AbsDiff(y, p0.Y)
		for x := p0.X - 20 + m; x <= p0.X + 20 - m; x++ {
			p1 := vec2.New(x, y)
			if dists.InGrid(p1) {
				d1 := dists.CellAt(p1)
				d := vec2.Distance(p0, p1)
				if d1 != -1 && d0 - d1 - d >= target {
					cheats++
				}
			}
		}
	}

	return cheats
}

func parttwo() string {
	track := grids.FromRunes(input)
	size := track.Size()
	dists := grids.FromSize(size.X, size.Y, -1)
	e := track.Find('E')

	pos := e
	prev := vec2.New(-1, -1)
	dist := 0
	path := make([]vec2.Vec2, 0, 10_000)

	for track.CellAt(pos) != 'S' {
		path = append(path, pos)
		dists.SetCellAt(pos, dist)
		for _, dir := range vec2.CardinalDirs {
			npos := pos.Add(dir)
			if track.CellAt(npos) != '#' && npos != prev {
				prev = pos
				pos = npos
				dist++
				break
			}
		}
	}
	dists.SetCellAt(pos, dist)
	path = append(path, pos)
	slices.Reverse(path)

	cheats := 0

	for _, p0 := range path {
		cheats += findCheats(dists, p0)
	}
	return strconv.Itoa(cheats)
}
