package utils

import (
	"os"
)

func ReadInput(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func MustReadInput(path string) string {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
