package main

import (
	"fmt"
	"os"
	"sort"

	"./interp"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("USAGE: lab_01 <datafile>\n")
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	eps := 0.00001
	bisErr := -1.111

	ds := interp.ReadDots(f)
	fmt.Printf("Table loaded from file:\n\n")
	ds.PrintDots()

	fmt.Printf("\nEnter X value and polynom degree: ")
	d, n := interp.ReadFuncData()

	sort.Sort(ds)

	d.Y = interp.Interpolation(ds, d, n)
	fmt.Printf("\nFunction value in %5.2f dot is %5.4f\n\n", d.X, d.Y)

	rootBis := interp.Bisection(ds, n, eps)
	if rootBis-bisErr < eps {
		fmt.Printf("Can't find root by bisection method in given dot set\n\n")
	} else {
		fmt.Printf("Root found by bisection method is %.4f\n", rootBis)
	}

	rootInvInterp := interp.InvInterpolation(ds, n)
	fmt.Printf("Root found by inverted interpolation method is %.4f\n\n",
		rootInvInterp)
}
