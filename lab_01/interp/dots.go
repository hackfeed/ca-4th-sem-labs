package interp

// Dot type used to represent cartesian dots.
type Dot struct {
	X float64
	Y float64
}

// DotSet type used to represent amount of cartesian dots.
type DotSet []Dot

// FTable used to represent table of float numbers.
type FTable [][]float64

// START: sort.Interface satisfying receiver functions.

func (ds DotSet) Len() int {
	return len(ds)
}

func (ds DotSet) Less(i, j int) bool {
	return ds[i].X < ds[j].X
}

func (ds DotSet) Swap(i, j int) {
	ds[i], ds[j] = ds[j], ds[i]
}

// END: sort.Interface satisfying receiver functions.

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
