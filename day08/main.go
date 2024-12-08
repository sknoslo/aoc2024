package main

import (
	"fmt"
	"sknoslo/aoc2024/utils"
	"sknoslo/aoc2024/vec2"
	"strings"

	"github.com/hashicorp/go-set/v3"
)

var input string

func init() {
	input = utils.MustReadInput("input.txt")
	input = strings.TrimSpace(input)
}

func main() {
	utils.Run(1, partone)
	utils.Run(2, parttwo)
}

func partone() string {
	antennas := make(map[rune][]vec2.Vec2)
	lines := strings.Split(input, "\n")
	w, h := len(lines[0]), len(lines)

	for y, line := range lines {
		for x, c := range line {
			if c != '.' {
				antennas[c] = append(antennas[c], vec2.New(x, y))
			}
		}
	}

	antinodes := set.New[vec2.Vec2](len(antennas))

	for _, pairs := range antennas {
		for i, a := range pairs[:len(pairs)-1] {
			for _, b := range pairs[i+1:] {
				diff := a.Sub(b)
				da := a.Add(diff)
				db := b.Sub(diff)
				if da.InRange(0,0,w-1,h-1) {
					antinodes.Insert(da)
				}
				if db.InRange(0,0,w-1,h-1) {
					antinodes.Insert(db)
				}
			}
		}
	}

	return fmt.Sprint(antinodes.Size())
}

func parttwo() string {
	antennas := make(map[rune][]vec2.Vec2)
	lines := strings.Split(input, "\n")
	w, h := len(lines[0]), len(lines)

	for y, line := range lines {
		for x, c := range line {
			if c != '.' {
				antennas[c] = append(antennas[c], vec2.New(x, y))
			}
		}
	}

	antinodes := set.New[vec2.Vec2](len(antennas))

	for _, pairs := range antennas {
		for i, a := range pairs[:len(pairs)-1] {
			for _, b := range pairs[i+1:] {
				diff := a.Sub(b)

				// turns out this wasn't necessary given the input, oh well, leaving it.
				gcd := utils.Gcd(diff.X, diff.Y)
				diff = diff.Div(gcd)

				antinodes.Insert(a)
				ds := a.Sub(diff)
				da := a.Add(diff)
				for ds.InRange(0,0,w-1,h-1) {
					antinodes.Insert(ds)
					ds = ds.Sub(diff)
				}
				for da.InRange(0,0,w-1,h-1) {
					antinodes.Insert(da)
					da = da.Add(diff)
				}
			}
		}
	}

	return fmt.Sprint(antinodes.Size())
}
