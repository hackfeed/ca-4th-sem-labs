package main

import (
	"fmt"
	"io"
)

type dot struct {
	X float64
	Y float64
}
type dotSet []dot

// START: sort.Interface satisfying receiver functions.

func (ds dotSet) Len() int {
	return len(ds)
}

func (ds dotSet) Less(i, j int) bool {
	return ds[i].X < ds[j].X
}

func (ds dotSet) Swap(i, j int) {
	ds[i], ds[j] = ds[j], ds[i]
}

// END: sort.Interface satisfying receiver functions.

func (ds *dotSet) Append(d dot) {
	*ds = append(*ds, d)
}

func (ds dotSet) GetPos(d dot) int {
	for i, el := range ds {
		if el.X == d.X {
			return i
		}
	}
	return -1
}

func readDots(f io.Reader, ds *dotSet, d *dot, n *int) {
	var (
		dotsNum int
		curDot  dot
	)

	fmt.Fscanln(f, &dotsNum)
	for i := 0; i < dotsNum; i++ {
		_, err := fmt.Fscanln(f, &curDot.X, &curDot.Y)
		if err == io.EOF {
			break
		}
		*ds = append(*ds, curDot)
	}
	fmt.Fscanln(f, &d.X)
	fmt.Fscanln(f, n)
}
