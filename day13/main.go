package main

import (
	"fmt"
	"sknoslo/aoc2024/utils"
	"sknoslo/aoc2024/vec2"
	"strings"
)

const inputf = "Button A: X+%d, Y+%d\nButton B: X+%d, Y+%d\nPrize: X=%d, Y=%d\n"

var input string

func init() {
	input = utils.MustReadInput("input.txt")
}

func main() {
	utils.Run(1, partone)
	utils.Run(2, parttwo)
}

type machine struct {
	a, b, p vec2.Vec2
}

func parseInput() []machine {
	machinestrs := strings.Split(input, "\n\n")

	out := make([]machine, 0, len(machinestrs))

	for _, m := range machinestrs {
		var ax, ay, bx, by, px, py int
		fmt.Sscanf(m, inputf, &ax, &ay, &bx, &by, &px, &py)
		fmt.Println(m)
		out = append(out, machine{vec2.New(ax, ay), vec2.New(bx, by), vec2.New(px, py)})
	}

	return out
}

func partone() string {
	tokens := 0

	for _, m := range strings.Split(input, "\n\n") {
		var ax, ay, bx, by, px, py int
		fmt.Sscanf(m, inputf, &ax, &ay, &bx, &by, &px, &py)

		// i*ax + j*bx = px
		// i = (px - j*bx)/ax
		// i*ay + j*by = py
		// ((px - j*bx)*ay)/ax + j*by = py
		// (px-j*bx)ay + j*by*ax = py*ax
		// px*ay - bx*ay*j + by*ax*j = py*ax
		// (by*ax-bx*ay)*j = py*ax - px*ay
		// j = (py*ax - px*ay)/(by*ax-bx*ay)
		t := py*ax - px*ay
		b := by*ax - bx*ay

		if b == 0 || t % b != 0 {
			continue
		}
		j := t / b
		r := (px - j*bx)
		if ax == 0 || r % ax != 0 {
			continue
		}
		i := r / ax

		tokens += i * 3 + j
	}
	return fmt.Sprint(tokens)
}

func parttwo() string {
	pdiff := 10000000000000

	tokens := 0

	for _, m := range strings.Split(input, "\n\n") {
		var ax, ay, bx, by, px, py int
		fmt.Sscanf(m, inputf, &ax, &ay, &bx, &by, &px, &py)
		px += pdiff
		py += pdiff

		t := py*ax - px*ay
		b := by*ax - bx*ay

		if b == 0 || t % b != 0 {
			continue
		}
		j := t / b
		r := (px - j*bx)
		if ax == 0 || r % ax != 0 {
			continue
		}
		i := r / ax

		tokens += i * 3 + j
	}
	return fmt.Sprint(tokens)
}
