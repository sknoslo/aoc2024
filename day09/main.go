package main

import (
	"fmt"
	"sknoslo/aoc2024/utils"
	"slices"
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

func partone() string {
	digits := strings.Split(input, "")

	fs := make([]int, 0, len(input))

	id := 0
	for i := 0; i < len(digits); i += 2 {
		for _ = range utils.MustAtoi(digits[i]) {
			fs = append(fs, id)
		}
		if i+1 < len(digits) {
			for _ = range utils.MustAtoi(digits[i+1]) {
				fs = append(fs, -1)
			}
		}
		id++
	}

	cfs := make([]int, 0, len(fs))

	s, e := 0, len(fs)-1
	for s <= e {
		if fs[s] != -1 {
			cfs = append(cfs, fs[s])
		} else {
			for fs[e] == -1 {
				e--
			}
			cfs = append(cfs, fs[e])
			e--
		}
		s++
	}

	res := 0
	for pos, id := range cfs {
		res += pos * id
	}
	return fmt.Sprint(res)
}

func parttwo() string {
	digits := strings.Split(input, "")

	fs := make([]int, 0, len(input))
	free := make(map[int][]int, 9)

	id := 0
	for i := 0; i < len(digits); i += 2 {
		for _ = range utils.MustAtoi(digits[i]) {
			fs = append(fs, id)
		}
		if i+1 < len(digits) {
			freestart := len(fs)
			freelen := utils.MustAtoi(digits[i+1])
			if freelen > 0 {
				free[freelen] = append(free[freelen], freestart)
			}
			for _ = range freelen {
				fs = append(fs, -1)
			}
		}
		id++
	}

	i := len(fs) - 1
	for i > 0 {
		id := fs[i]
		j := i
		for j >= 0 && fs[j] == id {
			j--
		}
		size := i - j

		firstspace := len(fs)
		firstspacesize := 0

		for s := size; s < 10; s++ {
			if len(free[s]) > 0 && free[s][0] < firstspace {
				firstspace = free[s][0]
				firstspacesize = s
			}
		}

		if firstspace < i && firstspacesize > 0 {
			pos := free[firstspacesize][0]
			for k := range size {
				fs[pos+k] = id
				fs[j+k+1] = -1
			}

			free[firstspacesize] = free[firstspacesize][1:]
			remainder := firstspacesize - size
			if remainder > 0 {
				free[remainder] = append(free[remainder], pos+size)
				slices.Sort(free[remainder])

			}
		}

		for j >= 0 && fs[j] == -1 {
			j--
		}
		i = j
	}

	res := 0
	for pos, id := range fs {
		if id == -1 {
			continue
		}
		res += pos * id
	}
	return fmt.Sprint(res)
}
