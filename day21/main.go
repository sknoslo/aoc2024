package main

import (
	"sknoslo/aoc2024/utils"
	"sknoslo/aoc2024/vec2"
	"strconv"
	"strings"
)

var input string

var keypad map[rune]vec2.Vec2
var dpad map[rune]vec2.Vec2

func init() {
	input = utils.MustReadInput("input.txt")
	keypad = make(map[rune]vec2.Vec2, 11)
	keypad['7'] = vec2.New(0, 0)
	keypad['8'] = vec2.New(1, 0)
	keypad['9'] = vec2.New(2, 0)
	keypad['4'] = vec2.New(0, 1)
	keypad['5'] = vec2.New(1, 1)
	keypad['6'] = vec2.New(2, 1)
	keypad['1'] = vec2.New(0, 2)
	keypad['2'] = vec2.New(1, 2)
	keypad['3'] = vec2.New(2, 2)
	keypad['0'] = vec2.New(1, 3)
	keypad['A'] = vec2.New(2, 3)

	dpad = make(map[rune]vec2.Vec2, 5)
	dpad['^'] = vec2.New(1, 0)
	dpad['A'] = vec2.New(2, 0)
	dpad['<'] = vec2.New(0, 1)
	dpad['v'] = vec2.New(1, 1)
	dpad['>'] = vec2.New(2, 1)
}

func main() {
	utils.Run(1, partone)
	utils.Run(2, parttwo)
}

func traverseKeypad(from, to rune) []rune {
	a, b := keypad[from], keypad[to]
	delta := b.Sub(a)
	dist := vec2.Distance(a, b)
	path := make([]rune, 0, dist+1)

	dx := delta.X
	dy := delta.Y

	if a.Y == 3 && b.X == 0 {
		// going left first is normally better, but...
		// when going left on bottom row, to left col
		// go up first to avoid the gap
		for range -dy {
			path = append(path, '^')
			dy++
		}
	}

	if a.X == 0 && b.Y == 3 {
		// going down first is normally better, but...
		// when going down on left col, to bottom row
		// go right first to avoid the gap
		for range dx {
			path = append(path, '>')
			dx--
		}
	}

	if dx < 0 {
		for range -dx {
			path = append(path, '<')
		}
	}

	if dy > 0 {
		for range dy {
			path = append(path, 'v')
		}
	}

	if dy < 0 {
		for range -dy {
			path = append(path, '^')
		}
	}

	if dx > 0 {
		for range dx {
			path = append(path, '>')
		}
	}

	return append(path, 'A')
}

func traverseDpad(from, to rune) []rune {
	if from == 'A' && to == '<' {
		return []rune{'v', '<', '<', 'A'}
	}
	a, b := dpad[from], dpad[to]
	delta := b.Sub(a)
	dist := vec2.Distance(a, b)
	path := make([]rune, 0, dist+1)

	dx := delta.X
	dy := delta.Y

	if a.Y == 0 && b.X == 0 {
		// going left first is normally better, but...
		// when going left on top row, to left col
		// go down first to avoid the gap
		for range dy {
			path = append(path, 'v')
			dy--
		}
	}

	if a.X == 0 && b.Y == 0 {
		// going up first is normally better, but...
		// when going right on left col, to top row
		// go right first to avoid the gap
		for range dx {
			path = append(path, '>')
			dx--
		}
	}

	if dx < 0 {
		for range -dx {
			path = append(path, '<')
		}
	}

	if dy > 0 {
		for range dy {
			path = append(path, 'v')
		}
	}

	if dy < 0 {
		for range -dy {
			path = append(path, '^')
		}
	}

	if dx > 0 {
		for range dx {
			path = append(path, '>')
		}
	}

	return append(path, 'A')
}

func partone() string {
	sum := 0

	for _, code := range strings.Split(input, "\n") {
		curr := 'A'
		robot := make([]rune, 0, 20)

		for _, key := range code {
			robot = append(robot, traverseKeypad(curr, key)...)
			curr = key
		}

		for range 2 {
			nextRobot := make([]rune, 0, len(robot)*3)

			curr = 'A'
			for _, key := range robot {
				nextRobot = append(nextRobot, traverseDpad(curr, key)...)
				curr = key
			}
			robot = nextRobot
		}

		sum += utils.MustAtoi(code[0:3]) * len(robot)
	}

	return strconv.Itoa(sum)
}

type Memory struct {
	robot    int
	from, to rune
}

var cache map[Memory]int = make(map[Memory]int, 1000)

func calculateLength(robot int, from, to rune) int {
	if robot == 25 {
		return len(traverseDpad(from, to))
	}

	memkey := Memory{robot, from, to}

	if v, ok := cache[memkey]; ok {
		return v
	}

	complexity := 0
	curr := 'A'
	for _, key := range traverseDpad(from, to) {
		complexity += calculateLength(robot+1, curr, key)
		curr = key
	}

	cache[memkey] = complexity

	return complexity
}

func parttwo() string {
	sum := 0

	for _, code := range strings.Split(input, "\n") {
		robot := make([]rune, 0, 20)
		curr := 'A'
		for _, key := range code {
			robot = append(robot, traverseKeypad(curr, key)...)
			curr = key
		}

		complexity := 0
		curr = 'A'
		for _, key := range robot {
			complexity += calculateLength(1, curr, key)
			curr = key
		}

		sum += utils.MustAtoi(code[0:3]) * complexity
	}

	return strconv.Itoa(sum)
}
