package main

import (
	"fmt"
	"sknoslo/aoc2024/utils"
	"strings"
)

var _ = fmt.Fprintln

var input string

var dirs = []utils.Vec2{
	utils.NewVec2(0, -1),
	utils.NewVec2(1, -1),
	utils.NewVec2(1, 0),
	utils.NewVec2(1, 1),
	utils.NewVec2(0, 1),
	utils.NewVec2(-1, 1),
	utils.NewVec2(-1, 0),
	utils.NewVec2(-1, -1),
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
						pos := utils.NewVec2(x, y).Add(dir.Mul(i + 1))
						if !pos.InRange(0, w-1, 0, h-1) || rows[pos.Y][pos.X] != byte(l) {
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
