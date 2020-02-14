package interpolation

import (
	"cartesian"
	"fmt"
	"io"
)

func GetBase(ds cartesian.DotSet, d cartesian.Dot, n int) cartesian.DotSet {
	base := cartesian.DotSet{}
	pos := ds.GetPos(d)

	if pos < n/2 {
		for i := 0; i < n+1; i++ {
			if i == pos {
				continue
			}
			base.Append(ds[i])
		}
	} else if len(ds)-pos-1 < n/2 {
		for i := len(ds) - n - 1; i < len(ds); i++ {
			if i == pos {
				continue
			}
			base.Append(ds[i])
		}
	} else {
		lb := n / 2
		rb := n - lb + 1

		for i := pos - lb; i < pos+rb; i++ {
			if i == pos {
				continue
			}
			base.Append(ds[i])
		}
	}
	return base
}

func ReadDots(f io.Reader, ds *cartesian.DotSet, d *cartesian.Dot, n *int) {
	var (
		dotsNum int
		curDot  cartesian.Dot
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
