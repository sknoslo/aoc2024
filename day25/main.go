package main

import (
	"sknoslo/aoc2024/grids"
	"sknoslo/aoc2024/utils"
	"strconv"
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

func partone() string {
	schematics := strings.Split(input, "\n\n")

	keys, locks := make([][5]int, 0, len(schematics)/2), make([][5]int, 0, len(schematics)/2)

	for _, schematic := range schematics {
		grid := grids.FromRunes(schematic)

		iskey := grid.CellAtXY(0, 0) == '.'

		if iskey {
			keys = append(keys, [5]int{})
		} else {
			locks = append(locks, [5]int{})
		}

		for col := range 5 {
			var row int
			for row = range 7 {
				cell := grid.CellAtXY(col, row)
				if iskey && cell == '#' || !iskey && cell == '.' {
					break
				}
			}
			if iskey {
				keys[len(keys)-1][col] = 7 - row
			} else {
				locks[len(locks)-1][col] = row
			}
		}
	}

	sum := 0
	for _, key := range keys {
	lockloop:
		for _, lock := range locks {
			for pin := range 5 {
				if lock[pin]+key[pin] > 7 {
					continue lockloop
				}
			}
			sum++
		}
	}
	return strconv.Itoa(sum)
}

func parttwo() string {
	return "Merry Christmas!"
}
