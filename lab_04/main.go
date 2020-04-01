package main

import (
	"fmt"
	"os"

	"./meansquare"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("USAGE: lab_04 <datafile>\n")
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	ds := meansquare.ReadDots(f)
	fmt.Printf("Table loaded from file:\n\n")
	ds.PrintDots()

	fmt.Printf("\nEnter polynom degree: ")
	n := meansquare.ReadPolynomDegree()

	slae := meansquare.MakeSLAE(ds, n)
	fmt.Printf("\nSLAE to solve:\n\n")
	slae.PrintMatrix()

	sol := meansquare.SolveSLAE(ds, n)
	fmt.Printf("\nSolved SLAE:\n\n")
	sol.PrintMatrix()
}
