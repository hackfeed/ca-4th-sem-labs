package interpolation

import (
	"cartesian"
	"fmt"
	"io"
)

func GetBase(ds cartesian.DotSet, d cartesian.Dot, n int) cartesian.DotSet {
	base := cartesian.DotSet{}
	pos := ds.GetPos(d)
	var t int

	if pos < n/2 {
		for i := 0; i < pos; i++ {
			base.Append(ds[i])
			t = n - pos
		}
		for i := pos + 1; i < pos+t+1; i++ {
			base.Append(ds[i])
		}
	}
	if len(ds)-pos < n/2 {
		rb := len(ds) - pos - 1
		lb := pos - n + rb

		for i := lb; lb < pos; i++ {
			base.Append(ds[i])
		}
		for i := pos + 1; i < rb; i++ {
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
