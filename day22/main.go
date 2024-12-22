package main

import (
	"sknoslo/aoc2024/utils"
	"strconv"
	"strings"

	"github.com/hashicorp/go-set/v3"
)

var input string

func init() {
	input = utils.MustReadInput("input.txt")
}

func main() {
	utils.Run(1, partone)
	utils.Run(2, parttwo)
}

func mix(secret, value int) int {
	return secret ^ value
}

func prune(secret int) int {
	return secret % 16_777_216
}

func partone() string {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		secret := utils.MustAtoi(line)

		for range 2_000 {
			secret = prune(mix(secret, secret * 64))
			secret = prune(mix(secret, secret / 32))
			secret = prune(mix(secret, secret * 2048))
		}

		sum += secret
	}
	return strconv.Itoa(sum)
}

func parttwo() string {
	bananas := make(map[int]int, 4096)

	best := 0

	for _, line := range strings.Split(input, "\n") {
		secret := utils.MustAtoi(line)
		s := set.New[int](2048)

		last := secret % 10
		sequence := 0

		for i := range 2_000 {
			secret = prune(mix(secret, secret * 64))
			secret = prune(mix(secret, secret / 32))
			secret = prune(mix(secret, secret * 2048))

			next := secret % 10
			change := next - last + 10 // don't deal with negative numbers
			sequence &= 0x3ffff // keep the newest 3 changes
			sequence <<= 6 // make space for next change
			sequence |= change

			if i > 2 && !s.Contains(sequence) {
				n := bananas[sequence] + next
				bananas[sequence] = n
				if n > best {
					best = n
				}
				s.Insert(sequence)
			}

			last = next
		}
	}

	return strconv.Itoa(best)
}
