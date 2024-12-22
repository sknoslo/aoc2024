package main

import (
	"fmt"
	"sknoslo/aoc2024/utils"
	"strings"
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

func operate(target, curr int, operands []int) bool {
	if len(operands) == 0 {
		return target == curr
	}

	return operate(target, curr+operands[0], operands[1:]) || operate(target, curr*operands[0], operands[1:])
}

func partone() string {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ": ")
		target, operands := utils.MustAtoi(parts[0]), utils.MustSplitIntsSep(parts[1], " ")

		if operate(target, 0, operands) {
			sum += target
		}
	}
	return fmt.Sprint(sum)
}

func concat(left, right int) int {
	pad := 10
	for right/pad > 0 {
		pad *= 10
	}
	return left*pad + right
}

func operate2(target, curr int, operands []int) bool {
	if curr > target {
		return false
	}

	if len(operands) == 0 {
		return target == curr
	}

	return operate2(target, curr+operands[0], operands[1:]) ||
		operate2(target, curr*operands[0], operands[1:]) ||
		operate2(target, concat(curr, operands[0]), operands[1:])
}

func parttwo() string {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ": ")
		target, operands := utils.MustAtoi(parts[0]), utils.MustSplitIntsSep(parts[1], " ")

		if operate2(target, 0, operands) {
			sum += target
		}
	}
	return fmt.Sprint(sum)
}
