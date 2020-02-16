package main

import (
	"cartesian"
	"fmt"
	"interpolation"
	"os"
	"sort"
)

func main() {
	var (
		d cartesian.Dot
		n int
	)
	ds := cartesian.DotSet{}

	f, err := os.Open("data/dots.txt")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	interpolation.ReadDots(f, &ds, &d, &n)
	ds.Append(d)
	sort.Sort(ds)
	base := interpolation.GetBase(ds, d, n)
	tb := interpolation.MakeTable(ds, d, n)
	p := interpolation.Interpolation(tb, d)
	fmt.Println(ds, d, base, tb, p)
}
