package main

import (
	"maps"
	"sknoslo/aoc2024/utils"
	"slices"
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

func partone() string {
	lines := strings.Split(input, "\n")
	computers := make(map[string][]string, len(lines))

	for _, line := range lines {
		parts := strings.Split(line, "-")
		computers[parts[0]] = append(computers[parts[0]], parts[1])
		computers[parts[1]] = append(computers[parts[1]], parts[0])
	}

	triples := set.New[string](len(lines))

	for a, pairs := range maps.All(computers) {
		if a[0] != 't' {
			continue
		}

		for i, b := range pairs[0 : len(pairs)-1] {
			for _, c := range pairs[i+1:] {
				if slices.Contains(computers[b], c) {
					network := []string{a, b, c}
					slices.Sort(network)
					triples.Insert(strings.Join(network, ","))
				}
			}
		}
	}

	return strconv.Itoa(triples.Size())
}

func parttwo() string {
	lines := strings.Split(input, "\n")
	computers := make(map[string]*set.Set[string], len(lines))

	for _, line := range lines {
		parts := strings.Split(line, "-")
		if _, ok := computers[parts[0]]; !ok {
			computers[parts[0]] = set.From([]string{parts[0]})
		}
		if _, ok := computers[parts[1]]; !ok {
			computers[parts[1]] = set.From([]string{parts[1]})
		}
		computers[parts[0]].Insert(parts[1])
		computers[parts[1]].Insert(parts[0])
	}

	maxlen := 0
	var maxpass string
mainloop:
	for a, anet := range maps.All(computers) {
		var counts [15]int

		for b := range anet.Items() {
			if a == b {
				continue
			}
			bnet := computers[b]
			inter := anet.Intersect(bnet)

			counts[inter.Size()-1]++
		}

		// hope that the answer will be the network that has the maximum number of
		// networks that have a certain count is equal to that count
		// ie, if 7 networks attached to anet have a size 7 intersection with anet... it might be a candidate
		// this shouldn't work in the general case, but seems like a reasonable heuristic for this input
		mustmatch := 0
		for i, c := range counts {
			if i == c && c >= maxlen {
				mustmatch = c
				goto nextstep
			}
		}
		continue
	nextstep:
		acc := anet.Copy()
		for b := range anet.Items() {
			if a == b {
				continue
			}
			bnet := computers[b]
			inter := anet.Intersect(bnet)
			if inter.Size() < mustmatch {
				continue
			}
			acc = acc.Intersect(bnet).(*set.Set[string])
			if acc.Size() < mustmatch {
				continue mainloop
			}
		}

		if acc.Size() > maxlen {
			maxlen = acc.Size()
			candidate := acc.Slice()
			slices.Sort(candidate)
			maxpass = strings.Join(candidate, ",")
		}
	}

	return maxpass
}
