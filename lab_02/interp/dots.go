package interp

import "fmt"

// Dot type used to represent cartesian dots.
type Dot struct {
	X float64
	Y float64
}

// DotSet type used to represent amount of cartesian dots.
type DotSet []Dot

type DotMatrix []DotSet

func (dm DotMatrix) PrintMatrix() {
	for i := 0; i < len(dm); i++ {
		for j := 0; j < len(dm); j++ {
			fmt.Printf("%8.2f ", dm[i][j].X)
		}
		fmt.Printf("\n")
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
			if d.X > ds[i-1].X && d.X <= ds[i].X {
				pos = i - 1
			}
		}
	}

	return pos
}
