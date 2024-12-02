package utils

func AbsDiff(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}
