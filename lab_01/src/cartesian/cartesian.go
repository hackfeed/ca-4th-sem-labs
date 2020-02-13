package cartesian

type Dot struct {
	X float64
	Y float64
}
type DotSet []Dot

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

func (ds *DotSet) Append(d Dot) {
	*ds = append(*ds, d)
}

func (ds DotSet) GetPos(d Dot) int {
	for i, el := range ds {
		if el.X == d.X {
			return i
		}
	}
	return -1
}
