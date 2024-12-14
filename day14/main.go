package main

import (
	"fmt"
	"sknoslo/aoc2024/utils"
	"sknoslo/aoc2024/vec2"
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

type robot struct {
	p, v vec2.Vec2
}

func parseInput() []robot {
	lines := strings.Split(input, "\n")
	robots := make([]robot, 0, len(lines))

	for _, line := range lines {
		var px, py, vx, vy int
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)
		robots = append(robots, robot{vec2.New(px, py), vec2.New(vx, vy)})
	}

	return robots
}

func partone() string {
	robots := parseInput()
	w, h := 101, 103
	e := 100

	var nw, ne, sw, se int

	for _, r := range robots {
		x := (r.p.X + r.v.X*e + e*w) % w
		y := (r.p.Y + r.v.Y*e + e*h) % h
		switch {
		case x < w/2 && y < h/2:
			nw++
		case x > w/2 && y < h/2:
			ne++
		case x < w/2 && y > h/2:
			sw++
		case x > w/2 && y > h/2:
			se++
		}
	}

	return fmt.Sprint(nw * ne * sw * se)
}

func parttwo() string {
	robots := parseInput()
	w, h := 101, 103
	e := 0

	tiles := make([][]rune, h)

mainloop:
	for true {
		e++
		for i := range h {
			if len(tiles[i]) == 0 {
				tiles[i] = make([]rune, w)
			}
			for j := range w {
				tiles[i][j] = ' '
			}
		}

		for _, r := range robots {
			x := (r.p.X + r.v.X*e + e*w) % w
			y := (r.p.Y + r.v.Y*e + e*h) % h

			tiles[y][x] = '#'
		}

		mustmatch := 20 // cross fingers and hope a tree has this many consecutive #
		for _, row := range tiles {
			c := 0
			for _, r := range row {
				if r == '#' {
					c++
				} else {
					c = 0
				}
				if c >= mustmatch {
					break mainloop
				}
			}
		}
	}

	for _, row := range tiles {
		fmt.Println(string(row))
	}

	return fmt.Sprint(e)
}
