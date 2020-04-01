package meansquare

import (
	"fmt"
	"io"
)

// Dot type used to represent plane dots with dot weight.
type Dot struct {
	X      float64
	Y      float64
	weight float64
}

// DotSet type used to represent amount of plane dots.
type DotSet []Dot

// ReadPolynomDegree used to read polynom degree.
func ReadPolynomDegree() int {
	var n int

	fmt.Scan(&n)

	return n
}

// ReadDots used to read Dot objects to DotSet object from file.
func ReadDots(f io.Reader) DotSet {
	var (
		ds     DotSet
		curDot Dot
	)

	for {
		_, err := fmt.Fscanln(f, &curDot.X, &curDot.Y, &curDot.weight)
		if err == io.EOF {
			break
		}
		ds = append(ds, curDot)
	}

	return ds
}

// PrintDots used to print dots in table form to standart output.
func (ds DotSet) PrintDots() {
	fmt.Printf("%8s %8s %8s\n", "X", "Y", "Weight")
	for i := range ds {
		fmt.Printf("%8.2f %8.2f %8.2f\n", ds[i].X, ds[i].Y, ds[i].weight)
	}
}
