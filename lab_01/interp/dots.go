package interp

import "fmt"

// Dot type used to represent cartesian dots.
type Dot struct {
	X float64
	Y float64
}

// DotSet type used to represent amount of cartesian dots.
type DotSet []Dot

// FTable used to represent table of float numbers.
type FTable [][]float64

// PrintDots used to print dots in table form to standart output.
func (ds DotSet) PrintDots() {
	for i := range ds {
		fmt.Printf("%8.2f %8.2f\n", ds[i].X, ds[i].Y)
	}
}

// GetPos used to find place where should insert dot to
// make yet set be sorted.
func (ds DotSet) GetPos(d Dot) int {
	var pos int

	if d.X < ds[0].X {
		pos = 0
	} else if d.X > ds[len(ds)-1].X {
		pos = len(ds) - 1
	} else {
		for i := 1; i < len(ds)-2; i++ {
			if d.X > ds[i-1].X && d.X < ds[i].X {
				pos = i
			}
		}
	}

	return pos
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

// AxisSwap used to swap dot's axises: X <-> Y.
func (ds DotSet) AxisSwap() {
	for i := range ds {
		ds[i].X, ds[i].Y = ds[i].Y, ds[i].X
	}
}
