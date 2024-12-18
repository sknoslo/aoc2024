package main

import (
	"fmt"
	"sknoslo/aoc2024/grids"
	"sknoslo/aoc2024/utils"
	"sknoslo/aoc2024/vec2"
	"strings"
	// "time"
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
	parts := strings.Split(input, "\n\n")
	grid := grids.FromRunes(parts[0])
	moves := strings.Join(strings.Split(parts[1], "\n"), "")
	var pos vec2.Vec2

	for s, v := range grid.Cells() {
		if v == '@' {
			pos = s
			grid.SetCellAt(pos, '.')
			break
		}
	}

	for _, move := range moves {
		var dir vec2.Vec2
		switch move {
		case '^':
			dir = vec2.North
		case '>':
			dir = vec2.East
		case 'v':
			dir = vec2.South
		case '<':
			dir = vec2.West
		}

		n := pos.Add(dir)

		if c := grid.CellAt(n); c == '.' {
			pos = n
		} else if c == 'O' {
			bn := n
			bc := grid.CellAt(bn)
			for bc != '#' {
				if bc == '.' {
					pos = n
					grid.SetCellAt(bn, 'O')
					grid.SetCellAt(n, '.')
					break
				}
				bn = bn.Add(dir)
				bc = grid.CellAt(bn)
			}
		}
	}

	sum := 0
	for s, v := range grid.Cells() {
		if v == 'O' {
			sum += 100*s.Y + s.X
		}
	}

	return fmt.Sprint(sum)
}

func doubleGrid(in string) (*grids.Grid[rune], vec2.Vec2) {
	lines := strings.Split(in, "\n")
	w, h := len(lines[0])*2, len(lines)
	cells := make([]rune, w*h)
	var pos vec2.Vec2

	for i, v := range strings.Join(lines, "") {
		if v == 'O' {
			cells[i*2] = '['
			cells[i*2+1] = ']'
			continue
		} else if v == '@' {
			v = '.'
			pos.X = i * 2 % w
			pos.Y = i * 2 / w
		}
		cells[i*2] = v
		cells[i*2+1] = v
	}

	return grids.New(w, h, cells), pos
}

func canPush(g *grids.Grid[rune], pos, dir vec2.Vec2) bool {
	n := pos.Add(dir)

	switch g.CellAt(n) {
	case '.':
		return true
	case '[':
		return canPush(g, n, dir) && canPush(g, n.Add(vec2.East), dir)
	case ']':
		return canPush(g, n.Add(vec2.West), dir) && canPush(g, n, dir)
	}

	return false
}

func pushUpDown(g *grids.Grid[rune], pos, dir vec2.Vec2) {
	n := pos.Add(dir)

	switch g.CellAt(n) {
	case '[':
		r := n.Add(vec2.East)
		pushUpDown(g, n, dir)
		pushUpDown(g, r, dir)
		g.SetCellAt(n, '.')
		g.SetCellAt(r, '.')
		g.SetCellAt(n.Add(dir), '[')
		g.SetCellAt(r.Add(dir), ']')
	case ']':
		l := n.Add(vec2.West)
		pushUpDown(g, l, dir)
		pushUpDown(g, n, dir)
		g.SetCellAt(l, '.')
		g.SetCellAt(n, '.')
		g.SetCellAt(l.Add(dir), '[')
		g.SetCellAt(n.Add(dir), ']')
	}
}

func pushLeftRight(g *grids.Grid[rune], pos, dir vec2.Vec2) bool {
	n := pos.Add(dir)
	t := g.CellAt(n)
	if t == '#' {
		return false
	} else if t == '.' || pushLeftRight(g, n, dir) {
		g.SetCellAt(n, g.CellAt(pos))
		g.SetCellAt(pos, '.')
		return true
	}

	return false
}

func parttwo() string {
	parts := strings.Split(input, "\n\n")
	g, pos := doubleGrid(parts[0])
	moves := strings.Join(strings.Split(parts[1], "\n"), "")

	for _, move := range moves {
		var dir vec2.Vec2
		switch move {
		case '^':
			dir = vec2.North
		case '>':
			dir = vec2.East
		case 'v':
			dir = vec2.South
		case '<':
			dir = vec2.West
		}

		n := pos.Add(dir)

		if c := g.CellAt(n); c == '.' {
			pos = n
		} else if c == '[' || c == ']' {
			if dir == vec2.North || dir == vec2.South {
				if canPush(g, pos, dir) {
					pushUpDown(g, pos, dir)
					pos = n
				}
			} else {
				if pushLeftRight(g, n, dir) {
					pos = n
				}
			}
		}

		// time.Sleep(16 * time.Millisecond)
		// fmt.Println(g.StringOverlayf("%c", '@', pos))
	}

	sum := 0
	for s, v := range g.Cells() {
		if v == '[' {
			sum += 100*s.Y + s.X
		}
	}

	return fmt.Sprint(sum)
}
