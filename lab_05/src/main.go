package main

import (
	"fmt"
	"math"

	"./integrate"
)

func main() {
	var (
		n, m, md1, md2, end int
		p                   float64
		f1, f2              integrate.Integrator
		sc                  []float64
	)

	sc = []float64{0.05, 0.05, 10}
	end = 0
	pl := integrate.CreatePlot("Integration using meansquare method", "Tao", "Result")
	for end != 1 {
		fmt.Printf("Enter N: ")
		fmt.Scan(&n)
		fmt.Printf("Enter M: ")
		fmt.Scan(&m)
		fmt.Printf("Enter parameter (tao): ")
		fmt.Scan(&p)
		fmt.Printf("Choose mode for outer integration (0 - Gauss, 1 - Simpson): ")
		fmt.Scan(&md1)
		if md1 == 0 {
			f1 = integrate.Gauss
		} else {
			f1 = integrate.Simpson
		}
		fmt.Printf("Choose mode for inner integration (0 - Gauss, 1 - Simpson): ")
		fmt.Scan(&md2)
		if md2 == 0 {
			f2 = integrate.Gauss
		} else {
			f2 = integrate.Simpson
		}

		lm := [][]float64{{0, math.Pi / 2}, {0, math.Pi / 2}}
		ns := []int{n, m}
		igs := []integrate.Integrator{f1, f2}

		pint := func(p float64) float64 {
			return integrate.Integrate(integrate.IntegratedFunc(p), lm, ns, igs)
		}

		fmt.Printf("Result with %.2f as a parameter is %.7f\n", p, pint(p))
		ds := integrate.GenDots(pint, sc)
		integrate.DrawPlot(pl, ds, n, m, md1, md2)
		fmt.Printf("Stop execution?: ")
		fmt.Scan(&end)
	}
	integrate.SavePlot(pl, "points.png")
}
