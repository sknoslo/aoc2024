package utils

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

var cpuprofile = flag.Bool("cpuprofile", false, "Write cpu profile to file")

func Run(part int, runner func() string) {
	flag.Parse()
	if *cpuprofile {
		f, err := os.Create(fmt.Sprintf("part_%d.prof", part))
		if err != nil {
			log.Fatal("Could not create cpu profile", err)
		}
		defer f.Close()

		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("Could not start cpu profile", err)
		}
		defer pprof.StopCPUProfile()
	}

	s := time.Now()
	res := runner()
	e := time.Now().Sub(s)
	fmt.Printf("Part %d: %s (%v)\n", part, res, e)
}
