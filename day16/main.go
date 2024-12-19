package main

import (
	"fmt"
	"sknoslo/aoc2024/pqueues"
	"sknoslo/aoc2024/grids"
	"sknoslo/aoc2024/utils"
	"sknoslo/aoc2024/vec2"
	"slices"

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

type step struct {
	pos, dir vec2.Vec2
	cost int
}

type step2 struct {
	pos, dir vec2.Vec2
	cost int
	path []vec2.Vec2
}

type memory struct {
	pos, dir vec2.Vec2
}

func partone() string {
	g := grids.FromRunes(input)
	q := pqueues.New[step](1024)
	s := set.New[memory](1024)
	q.Push(step{ g.Find('S'), vec2.East, 0 }, 0)

	for !q.Empty() {
		st := q.Pop()
		key := memory{ st.pos, st.dir }

		if c := g.CellAt(st.pos); c == 'E' {
			return fmt.Sprint(st.cost)
		} else if s.Contains(key) {
			continue
		}
		s.Insert(key)

		// TODO: perf - no sense in pushing every step along a straight path, could just add the length and skip ahead
		if g.CellAt(st.pos.Add(st.dir)) != '#' {
			q.Push(step{ st.pos.Add(st.dir), st.dir, st.cost + 1 }, st.cost + 1)
		}
		if ccw := st.dir.RotateCardinalCCW(); g.CellAt(st.pos.Add(ccw)) != '#' {
			q.Push(step{ st.pos, st.dir.RotateCardinalCCW(), st.cost + 1000 }, st.cost + 1000)
		}
		if cw := st.dir.RotateCardinalCW(); g.CellAt(st.pos.Add(cw)) != '#' {
			q.Push(step{ st.pos, st.dir.RotateCardinalCW(), st.cost + 1000 }, st.cost + 1000)
		}
	}

	return "no solution"
}

func parttwo() string {
	g := grids.FromRunes(input)
	q := pqueues.New[step2](1024)
	s := make(map[memory]int, 1024)
	q.Push(step2{ g.Find('S'), vec2.East, 0, make([]vec2.Vec2, 0, 1024) }, 0)
	b := -1
	p := make([][]vec2.Vec2, 0, 64)

	for !q.Empty() {
		st := q.Pop()
		key := memory{ st.pos, st.dir }

		if c := g.CellAt(st.pos); c == 'E' {
			if b == -1 {
				b = st.cost
			} else if b < st.cost {
				break
			}
			p = append(p, append(st.path, st.pos))
			continue
		} else if v, ok := s[key]; ok && st.cost > v {
			continue
		}
		s[key] = st.cost

		path := append(slices.Clone(st.path), st.pos)
		if g.CellAt(st.pos.Add(st.dir)) != '#' {
			q.Push(step2{ st.pos.Add(st.dir), st.dir, st.cost + 1, path }, st.cost + 1)
		}
		if ccw := st.dir.RotateCardinalCCW(); g.CellAt(st.pos.Add(ccw)) != '#' {
			q.Push(step2{ st.pos, st.dir.RotateCardinalCCW(), st.cost + 1000, path }, st.cost + 1000)
		}
		if cw := st.dir.RotateCardinalCW(); g.CellAt(st.pos.Add(cw)) != '#' {
			q.Push(step2{ st.pos, st.dir.RotateCardinalCW(), st.cost + 1000, path }, st.cost + 1000)
		}
	}

	cm := set.New[vec2.Vec2](1024)

	for _, path := range p {
		for _, t := range path {
			cm.Insert(t)
		}
	}

	return fmt.Sprint(cm.Size())
}
