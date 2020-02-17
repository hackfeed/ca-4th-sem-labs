package interp

import (
	"fmt"
	"io"
)

// Interpolation used to find value of X Dot coordinate
// using Newton polynom.
func Interpolation(tb [][]float64, d Dot) float64 {
	p := tb[1][0]
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

// MakeTable used to make table, which contains Newton
// polynom coefficients.
func MakeTable(ds DotSet, d Dot, n int) [][]float64 {
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

// GetBase used to make DotSet, which contains closest
// dots to find Newton polynom's coefficients.
func GetBase(ds DotSet, d Dot, n int) DotSet {
	base := DotSet{}
	pos := ds.GetPos(d)

	if pos <= n/2 {
		for i := 0; i < n+1; i++ {
			base = append(base, ds[i])
		}
	} else if len(ds)-pos-1 <= n/2 {
		for i := len(ds) - n - 1; i < len(ds); i++ {
			base = append(base, ds[i])
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
			base = append(base, ds[i])
		}
	}
	return base
}

// ReadDots used to read Dot objects to DotSet object from file.
func ReadDots(f io.Reader) (DotSet, Dot, int) {
	var (
		dotsNum int
		n       int
		ds      DotSet
		d       Dot
		curDot  Dot
	)

	fmt.Fscanln(f, &dotsNum)
	for i := 0; i < dotsNum; i++ {
		_, err := fmt.Fscanln(f, &curDot.X, &curDot.Y)
		if err == io.EOF {
			break
		}
		ds = append(ds, curDot)
	}
	fmt.Fscanln(f, &d.X)
	fmt.Fscanln(f, &n)

	return ds, d, n
}
