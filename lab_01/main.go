package main

import (
	"fmt"
	"os"
	"sort"

	"./interp"
)

func main() {
	var (
		d interp.Dot
		n int
	)
	ds := interp.DotSet{}

	f, err := os.Open("data/dots.txt")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	interp.ReadDots(f, &ds, &d, &n)
	sort.Sort(ds)
	base := interp.GetBase(ds, d, n)
	tb := interp.MakeTable(ds, d, n)
	p := interp.Interpolation(tb, d)
	fmt.Println(ds, d, base, tb, p)
}
