package main

import (
	"fmt"
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
	var a, b, c int
	out := make([]string, 0, 64)
	fmt.Sscanf(input, "Register A: %d\nRegister B: %d\nRegister C: %d\n", &a, &b, &c)

	groups := strings.Split(input, "\n\n")
	numstr := strings.Replace(groups[len(groups)-1], "Program: ", "", 1)
	program := utils.MustSplitIntsSep(numstr, ",")

	combo := func(ip int) int {
		operand := program[ip+1]
		switch {
		case operand < 4:
			return operand
		case operand == 4:
			return a
		case operand == 5:
			return b
		case operand == 6:
			return c
		}
		panic("invalid program")
	}

	ip := 0
	for ip < len(program) {
		n := ip + 2
		switch program[ip] {
		case 0:
			operand := utils.Pow(2, combo(ip))
			a = a / operand
		case 1:
			b = b ^ program[ip+1]
		case 2:
			b = combo(ip) % 8
		case 3:
			if a != 0 {
				n = program[ip+1]
			}
		case 4:
			b = b ^ c
		case 5:
			out = append(out, strconv.Itoa(combo(ip)%8))
		case 6:
			operand := utils.Pow(2, combo(ip))
			b = a / operand
		case 7:
			operand := utils.Pow(2, combo(ip))
			c = a / operand
		}
		ip = n
	}

	return strings.Join(out, ",")
}

func parttwo() string {
	groups := strings.Split(input, "\n\n")
	numstr := strings.Replace(groups[len(groups)-1], "Program: ", "", 1)
	program := utils.MustSplitIntsSep(numstr, ",")

	t := program

	mlen := 0
	s := 0
mainloop:
	for range t {
	targetloop:
		for true {
			var a, b, c int

			a = s
			out := make([]int, 0, 64)

			combo := func(ip int) int {
				operand := program[ip+1]
				switch {
				case operand < 4:
					return operand
				case operand == 4:
					return a
				case operand == 5:
					return b
				case operand == 6:
					return c
				}
				panic("invalid program")
			}

			ip := 0
			for ip < len(program) {
				n := ip + 2
				switch program[ip] {
				case 0:
					operand := utils.Pow(2, combo(ip))
					a = a / operand
				case 1:
					b = b ^ program[ip+1]
				case 2:
					b = combo(ip) % 8
				case 3:
					if a != 0 {
						n = program[ip+1]
					}
				case 4:
					b = b ^ c
				case 5:
					out = append(out, combo(ip)%8)
					if len(out) > mlen {
						for i := 1; i <= len(out); i++ {
							if out[len(out)-i] != t[len(t)-i] {
								s++
								continue targetloop
							}
						}
						if len(out) == len(t) {
							break mainloop
						} else if len(out) > mlen {
							mlen = len(out)
							s = s << 3
							break targetloop
						}
					}
				case 6:
					operand := utils.Pow(2, combo(ip))
					b = a / operand
				case 7:
					operand := utils.Pow(2, combo(ip))
					c = a / operand
				}
				ip = n
			}
		}
	}

	return strconv.Itoa(s)
}
