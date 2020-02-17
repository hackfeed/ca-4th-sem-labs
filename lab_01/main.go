package main

import (
	"fmt"
	"io"
	"os"
	"sort"

	"./interp"
)

func main() {
	f, err := os.Open("data/dots.txt")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	getComputations(f)
}

func getComputations(f io.Reader) {
	ds, d, n := interp.ReadDots(f)
	sort.Sort(ds)
	base := interp.GetBase(ds, d, n)
	tb := interp.MakeTable(ds, d, n)
	p := interp.Interpolation(tb, d)
	fmt.Println(ds, d, base, tb, p)
}
