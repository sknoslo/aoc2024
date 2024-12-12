package utils

import (
	"os"
	"strings"
)

func ReadInput(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(bytes)), nil
}

func MustReadInput(path string) string {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(bytes))
}
