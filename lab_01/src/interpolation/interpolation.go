package interpolation

import (
	"cartesian"
	"fmt"
	"io"
)

func Interpolation(tb [][]float64, d cartesian.Dot) float64 {
	p := 1.0
	var c float64

	for i := 2; i < len(tb); i++ {
		c = 1
		for j := 0; j < i-1; j++ {
			c *= d.X - tb[0][j]
		}
		c *= tb[i][0]
		p += c
	}
	return p
}

func MakeTable(ds cartesian.DotSet, d cartesian.Dot, n int) [][]float64 {
	base := GetBase(ds, d, n)
	baselen := len(base)
	tb := make([][]float64, n+2)
	tblen := len(tb)
	for i := range tb {
		tb[i] = make([]float64, baselen)
	}

	for i := range tb[0] {
		tb[0][i] = base[i].X
		tb[1][i] = base[i].Y
	}

	for i := 2; i < tblen; i++ {
		k := i - 2
		for j := 0; j < baselen-i+1; j++ {
			tb[i][j] = (tb[i-1][j] - tb[i-1][j+1]) / (tb[0][j] - tb[0][j+k+1])
		}
	}
	return tb
}

func GetBase(ds cartesian.DotSet, d cartesian.Dot, n int) cartesian.DotSet {
	base := cartesian.DotSet{}
	pos := ds.GetPos(d)

	if pos <= n/2 {
		for i := 0; i < n+1; i++ {
			if i == pos {
				continue
			}
			base.Append(ds[i])
		}
	} else if len(ds)-pos-1 <= n/2 {
		for i := len(ds) - n - 1; i < len(ds); i++ {
			if i == pos {
				continue
			}
			base.Append(ds[i])
		}
	} else {
		lb := n / 2
		rb := n - lb + 2
		if pos+rb > len(ds)-1 {
			rb--
			lb++
		}
		if pos-lb < 0 {
			rb++
			lb--
		}

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
