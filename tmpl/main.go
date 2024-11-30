package main

import (
	"log"
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
	return "incomplete"
}

func parttwo() string {
	return "incomplete"
}
