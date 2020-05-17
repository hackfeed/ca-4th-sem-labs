package differentiate

import (
	"fmt"
	"math"
)

func leftDiffInter(y, yl, h float64) float64 {
	return (y - yl) / h
}

// LeftDiff used to represent onesided subtractive derivative.
func LeftDiff(y []float64, h float64) (res []DFloat64) {
	for i := range y {
		if i == 0 {
			res = append(res, newNil())
		} else {
			res = append(res, newVal(leftDiffInter(y[i], y[i-1], h)))
		}
	}
	return
}

// CenterDiff used to represent center derivative.
func CenterDiff(y []float64, h float64) (res []DFloat64) {
	for i := range y {
		if i == 0 || i == len(y)-1 {
			res = append(res, newNil())
		} else {
			res = append(res, newVal((y[i+1]-y[i-1])/2*h))
		}
	}
	return
}

// SecondRungeDiff used to represent second Runge derivative.
func SecondRungeDiff(y []float64, h, p float64) (res []DFloat64) {
	var y2h []DFloat64

	for i := range y {
		if i < 2 {
			y2h = append(y2h, newVal(0))
		} else {
			y2h = append(y2h, newVal((y[i]-y[i-2])/(2.*h)))
		}
	}

	yh := LeftDiff(y, h)
	for i := range yh {
		if i < 2 {
			res = append(res, newNil())
		} else {
			res = append(res, newVal(yh[i].value+(yh[i].value-y2h[i].value)/(math.Pow(2., p)-1)))
		}
	}
	return
}

// AlignedCoeffsDiff used to represent derivative with aligned parameters.
func AlignedCoeffsDiff(x, y []float64) (res []DFloat64) {
	for i := range y {
		if i == len(y)-1 {
			res = append(res, newNil())
		} else {
			k := y[i] * y[i] / x[i] / x[i]
			res = append(res, newVal(k*leftDiffInter(-1./y[i+1], -1./y[i], -1./x[i+1]- -1./x[i])))
		}
	}
	return
}

// SecondLeftDiff used to represent second subtractive derivative.
func SecondLeftDiff(y []float64, h float64) (res []DFloat64) {
	for i := range y {
		if i == 0 || i == len(y)-1 {
			res = append(res, newNil())
		} else {
			res = append(res, newVal((y[i-1]-2*y[i]+y[i+1])/(h*h)))
		}
	}
	return
}

// FmtPrintInit used to init formatted output.
func FmtPrintInit(text string, init []float64) {
	fmt.Printf("%s ", text)
	for _, v := range init {
		fmt.Printf("%7.4f ", v)
	}
	fmt.Printf("\n")
}

// FmtPrintRes used to init formatted output.
func FmtPrintRes(text string, res []DFloat64) {
	fmt.Printf("%s ", text)
	for _, v := range res {
		fmt.Printf("%7.4f ", v.retvalue())
	}
	fmt.Printf("\n")
}
