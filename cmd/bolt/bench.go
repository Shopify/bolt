package main

import (
	"github.com/boltdb/bolt/bench"
)

// Run benchmarks on a given dataset.
func Bench(path string) {
	b, err := bench.New(path)
	if err != nil {
		fatal(err)
	}
	if err := b.Run(); err != nil {
		fatal(err)
	}
}
