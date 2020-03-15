package spline

import (
	"fmt"
	"io"
)

// Dot type used to represent plane dots.
type Dot struct {
	X float64
	Y float64
}

// DotSet type used to represent amount of plane dots.
type DotSet []Dot

// ReadX used to read function X coordinate.
func ReadX() Dot {
	var d Dot

	fmt.Scan(&d.X)

	return d
}

// ReadDots used to read Dot objects to DotSet object from file.
func ReadDots(f io.Reader) DotSet {
	var (
		ds     DotSet
		curDot Dot
	)

	for {
		_, err := fmt.Fscanln(f, &curDot.X, &curDot.Y)
		if err == io.EOF {
			break
		}
		ds = append(ds, curDot)
	}

	return ds
}

// PrintDots used to print dots in table form to standart output.
func (ds DotSet) PrintDots() {
	for i := range ds {
		fmt.Printf("%8.2f %8.2f\n", ds[i].X, ds[i].Y)
	}
}

func (ds DotSet) Len() int {
	return len(ds)
}

func (ds DotSet) Less(i, j int) bool {
	return ds[i].X < ds[j].X
}

func (ds DotSet) Swap(i, j int) {
	ds[i], ds[j] = ds[j], ds[i]
}

func (ds DotSet) getPos(d Dot) int {
	var pos int

	if d.X < ds[0].X {
		pos = 0
	} else if d.X > ds[len(ds)-1].X {
		pos = len(ds) - 1
	} else {
		for i := 1; i < len(ds)-2; i++ {
			if d.X > ds[i-1].X && d.X <= ds[i].X {
				pos = i - 1
			}
		}
	}

	return pos
}
