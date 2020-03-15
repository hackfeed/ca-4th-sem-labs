package main

import (
	"fmt"
	"os"
	"sort"

	"./spline"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("USAGE: lab_03 <datafile>\n")
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	ds := spline.ReadDots(f)
	fmt.Printf("Table loaded from file:\n\n")
	ds.PrintDots()

	fmt.Printf("\nEnter X value: ")
	d := spline.ReadX()

	sort.Sort(ds)
	spline.Spline(ds, &d)

	fmt.Printf("\nFunction value in %5.2f dot is %5.4f\n\n", d.X, d.Y)
}
