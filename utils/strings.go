package utils

import (
	"strconv"
	"strings"
)

func SplitInts(in string) ([]int, error) {
	splits := strings.Fields(in)
	out := make([]int, len(splits))

	for i, s := range splits {
		v, err := strconv.Atoi(s)
		if err != nil {
			return out, err
		}
		out[i] = v
	}

	return out, nil
}

func MustSplitIntsSep(in, sep string) []int {
	splits := strings.Split(in, sep)
	out := make([]int, len(splits))

	for i, s := range splits {
		v, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		out[i] = v
	}

	return out
}
