package interp

import (
	"fmt"
	"io"
	"math"
	"sort"
)

// InvInterpolation used to find root of function using inverted interpolation.
func InvInterpolation(ds DotSet, n int) float64 {
	ds.AxisSwap()
	sort.Sort(ds)

	d := Dot{
		X: 0,
	}

	return Interpolation(ds, d, n)
}

// Bisection used to find root of function using bisection method.
func Bisection(ds DotSet, n int, eps float64) float64 {
	if math.Signbit(ds[0].Y) == math.Signbit(ds[len(ds)-1].Y) {
		return -1.111
	}
	low, high := ds[0].X, ds[len(ds)-1].X
	mid := (low + high) / 2

	for math.Abs(high-low) >= eps*mid+eps {
		lowd := Dot{
			X: low,
		}
		midd := Dot{
			X: mid,
		}
		if Interpolation(ds, lowd, n)*Interpolation(ds, midd, n) < 0 {
			high = mid
		} else {
			low = mid
		}
		mid = (low + high) / 2
	}

	return mid
}

// Interpolation used to find value of X Dot coordinate
// using Newton polynom.
func Interpolation(ds DotSet, d Dot, n int) float64 {
	tb := MakeTable(ds, d, n)
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

// ReadFuncData used to read function X coordinate and polynom degree.
func ReadFuncData() (Dot, int) {
	var (
		d Dot
		n int
	)

	fmt.Scan(&d.X, &n)

	return d, n
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
