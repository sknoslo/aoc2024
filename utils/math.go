package utils

func AbsDiff(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}

func Gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func Pow(a, b int) int {
	res := 1
	for range b {
		res *= a
	}

	return res
}
