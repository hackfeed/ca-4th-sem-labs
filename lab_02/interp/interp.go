package interp

import (
	"fmt"
	"io"
)

// FTable used to represent table of float numbers.
type FTable [][]float64

// BiInterpolation used to find value of X,Y Dot coordinate
// using Newton polynom.
func BiInterpolation(dm DotMatrix, d Dot, nx, ny int) float64 {
	xcol, yrow, zdm := splitData(dm)
	masterSet := make(DotSet, 0)

	baseX, startX, endX := getDimBase(xcol, d, nx)
	d.X, d.Y = d.Y, d.X
	baseY, startY, endY := getDimBase(yrow, d, ny)

	ftb := make(DotMatrix, nx+1)
	for i := range ftb {
		ftb[i] = make(DotSet, ny+1)
	}

	k := 0
	for i := startX; i < endX+1; i++ {
		l := 0
		for j := startY; j < endY+1; j++ {
			ftb[k][l] = zdm[i][j]
			l++
		}
		k++
	}

	for i := 0; i < len(ftb); i++ {
		for j := 0; j < len(baseY); j++ {
			ftb[i][j].Y = baseY[j].X
			ftb[i][j].X, ftb[i][j].Y = ftb[i][j].Y, ftb[i][j].X
		}
		md := Dot{
			X: baseX[i].X,
			Y: Interpolation(ftb[i], d, ny),
		}
		masterSet = append(masterSet, md)
	}

	d.X, d.Y = d.Y, d.X

	return Interpolation(masterSet, d, nx)
}

// Interpolation used to find value of X Dot coordinate
// using Newton polynom.
func Interpolation(ds DotSet, d Dot, n int) float64 {
	tb := makeTable(ds, d, n)
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

// ReadFuncData used to read function X coordinate and polynom degree.
func ReadFuncData() (Dot, int, int) {
	var (
		d      Dot
		nx, ny int
	)

	fmt.Scan(&d.X, &d.Y, &nx, &ny)

	return d, nx, ny
}

// ReadDots used to read Dot objects to DotSet object from file.
func ReadDots(f io.Reader) DotMatrix {
	var (
		ds   int
		data DotMatrix
	)

	fmt.Fscanln(f, &ds)

	data = make(DotMatrix, ds)
	for i := range data {
		data[i] = make(DotSet, ds)
	}

	for i := 0; i < ds; i++ {
		for j := 0; j < ds; j++ {
			fmt.Fscan(f, &data[i][j].X)
		}
	}

	return data
}

func getDimBase(ds DotSet, d Dot, n int) (DotSet, int, int) {
	base := getBase(ds, d, n)
	start, end := 0, 0

	for i, el := range ds {
		if el == base[0] {
			start = i
			end = start + n
			break
		}
	}

	return base, start, end
}

func makeTable(ds DotSet, d Dot, n int) FTable {
	base := getBase(ds, d, n)
	baselen := len(base)
	tb := make(FTable, n+2)
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
			tb[i][j] = (tb[i-1][j] - tb[i-1][j+1]) /
				(tb[0][j] - tb[0][j+k+1])
		}
	}

	return tb
}

func getBase(ds DotSet, d Dot, n int) DotSet {
	base := DotSet{}
	pos := ds.getPos(d)

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
		rb := n - lb + 1
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

func splitData(dm DotMatrix) (DotSet, DotSet, DotMatrix) {
	xcol := make(DotSet, 0)
	for i := 1; i < len(dm); i++ {
		xcol = append(xcol, dm[i][0])
	}

	yrow := dm[0][1:]

	zdm := make(DotMatrix, len(dm)-1)
	for i := range zdm {
		zdm[i] = make(DotSet, len(dm)-1)
	}

	for i := 1; i < len(dm); i++ {
		for j := 1; j < len(dm); j++ {
			zdm[i-1][j-1] = dm[i][j]
		}
	}

	return xcol, yrow, zdm
}
