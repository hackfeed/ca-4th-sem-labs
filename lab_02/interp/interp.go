package interp

import (
	"fmt"
	"io"
)

// FTable used to represent table of float numbers.
type FTable [][]float64

// // Interpolation used to find value of X Dot coordinate
// // using Newton polynom.
// func Interpolation(ds DotSet, d Dot, n int) float64 {
// 	tb := MakeTable(ds, d, n)
// 	p := tb[1][0]
// 	var c float64

// 	for i := 2; i < len(tb); i++ {
// 		c = 1
// 		for j := 0; j < i-1; j++ {
// 			c *= d.X - tb[0][j]
// 		}
// 		c *= tb[i][0]
// 		p += c

// 	}

// 	return p
// }

// // MakeTable used to make table, which contains Newton
// // polynom coefficients.
// func MakeTable(ds DotSet, d Dot, n int) FTable {
// 	base := GetBase(ds, d, n)
// 	baselen := len(base)
// 	tb := make(FTable, n+2)
// 	tblen := len(tb)
// 	for i := range tb {
// 		tb[i] = make([]float64, baselen)
// 	}

// 	for i := range tb[0] {
// 		tb[0][i] = base[i].X
// 		tb[1][i] = base[i].Y
// 	}

// 	for i := 2; i < tblen; i++ {
// 		k := i - 2
// 		for j := 0; j < baselen-i+1; j++ {
// 			tb[i][j] = (tb[i-1][j] - tb[i-1][j+1]) /
// 				(tb[0][j] - tb[0][j+k+1])
// 		}
// 	}

// 	return tb
// }

// // GetBase used to make DotSet, which contains closest
// // dots to find Newton polynom's coefficients.
// func GetBase(ds DotSet, d Dot, n int) DotSet {
// 	base := DotSet{}
// 	pos := ds.GetPos(d)

// 	if pos <= n/2 {
// 		for i := 0; i < n+1; i++ {
// 			base = append(base, ds[i])
// 		}
// 	} else if len(ds)-pos-1 <= n/2 {
// 		for i := len(ds) - n - 1; i < len(ds); i++ {
// 			base = append(base, ds[i])
// 		}
// 	} else {
// 		lb := n/2 - 1
// 		rb := n - lb + 1
// 		if pos+rb > len(ds)-1 {
// 			rb--
// 			lb++
// 		}
// 		if pos-lb < 0 {
// 			rb++
// 			lb--
// 		}

// 		for i := pos - lb; i < pos+rb; i++ {
// 			base = append(base, ds[i])
// 		}
// 	}

// 	return base
// }

// ReadFuncData used to read function X coordinate and polynom degree.
func ReadFuncData() (Dot, int) {
	var (
		d Dot
		n int
	)

	fmt.Scan(&d.X, &d.Y, &n)

	return d, n
}

// ReadDots used to read Dot objects to DotSet object from file.
func ReadDots(f io.Reader) FTable {
	var (
		ds   int
		data FTable
	)

	fmt.Fscanln(f, &ds)

	data = make(FTable, ds)
	for i := range data {
		data[i] = make([]float64, ds)
	}

	for i := 0; i < ds; i++ {
		for j := 0; j < ds; j++ {
			fmt.Fscan(f, &data[i][j])
		}
	}

	fmt.Println(data)
	return data
}
