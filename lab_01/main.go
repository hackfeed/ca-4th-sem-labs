package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	var (
		d dot
		n int
	)
	ds := dotSet{}

	f, err := os.Open("dots.txt")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	readDots(f, &ds, &d, &n)
	ds.Append(d)
	sort.Sort(ds)
	pos := ds.GetPos(d)
	fmt.Println(ds, d, n, pos)
}
