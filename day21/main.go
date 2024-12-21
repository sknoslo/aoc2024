package main

import (
	"fmt"
	"sknoslo/aoc2024/utils"
	"sknoslo/aoc2024/vec2"
	"strconv"
	"strings"
)

var input string

var keypad map[rune]vec2.Vec2
var dpad map[rune]vec2.Vec2

func init() {
	input = utils.MustReadInput("example.txt")
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
	path := make([]rune, 0, dist + 1)

	horizontal := '<'
	if delta.X > 0 {
		horizontal = '>'
	}

	if delta.Y < 0 {
		for range -delta.Y {
			path = append(path, '^')
		}
	}

	for range utils.Abs(delta.X) {
		path = append(path, horizontal)
	}

	if delta.Y > 0 {
		for range delta.Y {
			path = append(path, 'v')
		}
	}

	return append(path, 'A')
}

func traverseDpad(from, to rune) []rune {
	a, b := dpad[from], dpad[to]
	delta := b.Sub(a)
	dist := vec2.Distance(a, b)
	path := make([]rune, 0, dist + 1)

	horizontal := '<'
	if delta.X > 0 {
		horizontal = '>'
	}

	if delta.Y > 0 {
		for range delta.Y {
			path = append(path, 'v')
		}
	}

	for range utils.Abs(delta.X) {
		path = append(path, horizontal)
	}

	if delta.Y < 0 {
		for range -delta.Y {
			path = append(path, '^')
		}
	}

	return append(path, 'A')
}

func partone() string {
	// To deal with gaps...
	// when using keypad, going up and left, prioritize going up first
	//                    going down and right, prioritize going right first
	// when using d-pad, going down and left, prioritize going down first
	//                   going up and right, prioritize going right first
	// the shortest path should always be to go in straight lines, never zig or zag
	// which means we just need the x,y dist from each key a to b and we can calculate a min path?

	fmt.Println(string(traverseDpad('A', '<')))
	fmt.Println(string(traverseDpad('<', 'A')))
	fmt.Println()
	fmt.Println(string(traverseDpad('A', '>')))
	fmt.Println(string(traverseDpad('>', 'A')))
	fmt.Println()
	fmt.Println(string(traverseDpad('A', '^')))
	fmt.Println(string(traverseDpad('^', 'A')))
	fmt.Println()
	fmt.Println(string(traverseDpad('A', 'v')))
	fmt.Println(string(traverseDpad('v', 'A')))

	sum := 0

	for _, code := range strings.Split(input, "\n") {
		curr := 'A'
		robot1 := make([]rune, 0, 20)

		for _, key := range code {
			robot1 = append(robot1, traverseKeypad(curr, key)...)
			curr = key
		}

		fmt.Println(code, string(robot1))

		robot2 := make([]rune, 0, len(robot1) * 3)

		curr = 'A'
		for _, key := range robot1 {
			robot2 = append(robot2, traverseDpad(curr, key)...)
			curr = key
		}

		fmt.Println(code, string(robot2))

		robot3 := make([]rune, 0, len(robot2) * 3)

		curr = 'A'
		for _, key := range robot2 {
			robot3 = append(robot3, traverseDpad(curr, key)...)
			curr = key
		}

		fmt.Println(code, string(robot3))

		fmt.Println(len(robot3), "*", utils.MustAtoi(code[0:3]))
		sum += utils.MustAtoi(code[0:3]) * len(robot3)
	}

	return strconv.Itoa(sum)
}

func parttwo() string {
	return "incomplete"
}

//                                         *
// <v<A>>^AvA^A 3 <vA<AA>>^AAvA<^A>AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A

//                                        *
// v<<A>>^AvA^A 3 v<<A>>^AAv<A<A>>^AAvAA<^A>Av<A>^AA<A>Av<A<A>>^AAAvA<^A>A
