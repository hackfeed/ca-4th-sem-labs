package main

import (
	"fmt"
	"os"
	"sort"

	"./interp"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("USAGE: lab_02 <datafile>\n")
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	ds := interp.ReadDots(f)
	fmt.Printf("Table loaded from file:\n\n")
	ds.PrintDots()

	fmt.Printf("\nEnter X value and polynom degree: ")
	d, n := interp.ReadFuncData()

	sort.Sort(ds)

	d.Y = interp.Interpolation(ds, d, n)
	fmt.Printf("\nFunction value in %5.2f dot is %5.4f\n\n", d.X, d.Y)
}
