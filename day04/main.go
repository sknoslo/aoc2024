package main

import (
	"fmt"
	"sknoslo/aoc2024/utils"
	"sknoslo/aoc2024/vec2"
	"strings"
)

var _ = fmt.Fprintln

var input string

var dirs = []vec2.Vec2{
	vec2.New(0, -1),
	vec2.New(1, -1),
	vec2.New(1, 0),
	vec2.New(1, 1),
	vec2.New(0, 1),
	vec2.New(-1, 1),
	vec2.New(-1, 0),
	vec2.New(-1, -1),
}

func init() {
	input = strings.TrimSpace(utils.MustReadInput("input.txt"))
}

func main() {
	utils.Run(1, partone)
	utils.Run(2, parttwo)
}

func partone() string {
	rows := strings.Split(input, "\n")
	w, h := len(rows[0]), len(rows)

	xmases := 0

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if rows[y][x] == 'X' {
			dirloop:
				for _, dir := range dirs {
					for i, l := range "MAS" {
						pos := vec2.New(x, y).Add(dir.Mul(i + 1))
						if !pos.InRange(0, 0, w-1, h-1) || rows[pos.Y][pos.X] != byte(l) {
							continue dirloop
						}
					}
					xmases++
				}
			}
		}
	}

	return fmt.Sprint(xmases)
}

func parttwo() string {
	rows := strings.Split(input, "\n")
	w, h := len(rows[0]), len(rows)

	xmases := 0

	for y := 1; y < h-1; y++ {
		for x := 1; x < w-1; x++ {
			if rows[y][x] == 'A' {
				bs := rows[y-1][x-1] + rows[y+1][x+1]
				fs := rows[y-1][x+1] + rows[y+1][x-1]

				if bs == 'M'+'S' && fs == 'M'+'S' {
					xmases++
				}
			}
		}
	}

	return fmt.Sprint(xmases)
}
