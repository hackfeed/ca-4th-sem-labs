package main

import (
	"fmt"
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

	ds := interp.ReadDots(f)
	fmt.Printf("Table loaded from file:\n\n")
	ds.PrintDots()

	fmt.Printf("\nEnter X value and polynom degree: ")
	d, n := interp.ReadFuncData()

	sort.Sort(ds)

	d.Y = interp.Interpolation(ds, d, n)
	fmt.Printf("Function value in %5.2f dot is %5.4f\n\n", d.X, d.Y)

	rootBis := interp.Bisection(ds, n, 0.001)
	fmt.Printf("Root found by bisection method is %.4f\n", rootBis)
	fmt.Printf("If root value is equivalent to -1.1110, that means that\n" +
		"there is no possibility to find root by bisection method in given dot set.\n\n")

	rootInvInterp := interp.InvInterpolation(ds, n)
	fmt.Printf("Root found by inverted interpolation method is %.4f\n\n", rootInvInterp)
}
