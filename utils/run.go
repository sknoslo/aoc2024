package utils

import (
	"fmt"
	"time"
)

func Run(part int, runner func() string) {
	s := time.Now()
	res := runner()
	e := time.Now().Sub(s)
	fmt.Printf("Part %d: %s (%v)\n", part, res, e)
}
