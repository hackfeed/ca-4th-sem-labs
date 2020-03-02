package main

import (
	"fmt"
	"os"

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

	dm := interp.ReadDots(f)
	dm.PrintMatrix()

	fmt.Printf("\nEnter X, Y value and X, Y polynom degrees: ")
	d, nx, ny := interp.ReadFuncData()
	interp.BiInterpolation(dm, d, nx, ny)

	// sort.Sort(ds)

	// d.Y = interp.Interpolation(ds, d, n)
	// fmt.Printf("\nFunction value in %5.2f dot is %5.4f\n\n", d.X, d.Y)
}
