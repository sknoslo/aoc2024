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
