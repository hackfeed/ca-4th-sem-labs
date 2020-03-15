package spline

import (
	"math"
)

// Spline used to find approximated value of function given in table representation.
func Spline(ds DotSet, xd *Dot) {
	aCoef := []float64{0}
	bCoef := []float64{0}
	cCoef := []float64{0}
	dCoef := []float64{0}
	hCoef := []float64{0}
	xiCoef := []float64{0, 0, 0}
	etaCoef := []float64{0, 0, 0}

	for _, d := range ds {
		aCoef = append(aCoef, d.Y)
	}

	for i := 1; i < len(ds); i++ {
		hCoef = append(hCoef, ds[i].X-ds[i-1].X)
		cCoef = append(cCoef, 0)
	}
	cCoef = append(cCoef, 0)

	for i := 3; i < len(ds)+1; i++ {
		xii := hCoef[i-1] /
			(-2*(hCoef[i-1]+hCoef[i-2]) - hCoef[i-2]*xiCoef[i-1])
		xiCoef = append(xiCoef, xii)

		fStrong := -3 * ((ds[i-1].Y-ds[i-2].Y)/hCoef[i-1] -
			(ds[i-2].Y-ds[i-3].Y)/hCoef[i-2])
		etai := (fStrong + hCoef[i-2]*etaCoef[i-1]) /
			(-2*(hCoef[i-1]+hCoef[i-2]) - hCoef[i-2]*xiCoef[i-1])
		etaCoef = append(etaCoef, etai)
	}

	for i := len(ds) - 1; i > 1; i-- {
		cCoef[i] = xiCoef[i+1]*cCoef[i+1] + etaCoef[i+1]
	}

	for i := 1; i < len(ds); i++ {
		bi := (aCoef[i+1]-aCoef[i])/hCoef[i] -
			hCoef[i]/3*(cCoef[i+1]+2*cCoef[i])
		bCoef = append(bCoef, bi)

		di := (cCoef[i+1] - cCoef[i]) / 3 / hCoef[i]
		dCoef = append(dCoef, di)
	}
	bCoef = append(bCoef, 0)
	dCoef = append(dCoef, 0)

	pos := ds.getPos(*xd)

	res := aCoef[pos+1] +
		bCoef[pos+1]*(xd.X-ds[pos].X) +
		cCoef[pos+1]*math.Pow((xd.X-ds[pos].X), 2) +
		dCoef[pos+1]*math.Pow((xd.X-ds[pos].X), 3)

	xd.Y = res
}
