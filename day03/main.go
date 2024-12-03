package main

import (
	"fmt"
	"log"
	"regexp"
	"sknoslo/aoc2024/utils"
)

var input string

func init() {
	in, err := utils.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input = in
}

func main() {
	utils.Run(1, partone)
	utils.Run(2, parttwo)
}

func partone() string {
	mulregex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	muls := mulregex.FindAllString(input, -1)

	sum := 0

	for _, mul := range muls {
		var a, b int
		fmt.Sscanf(mul, "mul(%d,%d)", &a, &b)
		sum += a * b
	}

	return fmt.Sprint(sum)
}

func parttwo() string {
	instregex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|don't\(\)|do\(\)`)
	insts := instregex.FindAllString(input, -1)

	sum := 0
	do := true

	for _, inst := range insts {
		switch inst {
		case "don't()":
			do = false
		case "do()":
			do = true
		default:
			if do {
				var a, b int
				fmt.Sscanf(inst, "mul(%d,%d)", &a, &b)
				sum += a * b
			}
		}
	}

	return fmt.Sprint(sum)
}
