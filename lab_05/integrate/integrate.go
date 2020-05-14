package integrate

import (
	"math"

	"gonum.org/v1/gonum/integrate/quad"
)

// Integrator type used to represent integrator function.
type Integrator func(func(float64) float64, float64, float64, int) float64

// Integrated type used to represent integrated function.
type Integrated func(float64, float64) float64

// Simpson function used to integrate with Simpson method.
func Simpson(f func(float64) float64, a, b float64, n int) float64 {
	if n < 3 || n&1 == 0 {
		panic("Error")
	}

	h := (b - a) / float64(n-1)
	x := a
	res := 0.0

	for i := 0; i < (n-1)/2; i++ {
		res += f(x) + 4*f(x+h) + f(x+2*h)
		x += 2 * h
	}

	return res * (h / 3)
}

// Gauss function used to integrate with Gauss-Legendre quadrature.
func Gauss(f func(float64) float64, a, b float64, n int) float64 {
	var (
		x, weight []float64
		l         quad.Legendre
		res       float64
	)
	x = make([]float64, n)
	weight = make([]float64, n)

	l.FixedLocations(x, weight, -1, 1)
	res = 0.0

	for i := 0; i < n; i++ {
		res += (b - a) / 2 * weight[i] * f(pToV(x[i], a, b))
	}

	return res
}

// IntegratedFunc function used to choose inner and outer direction of integration.
func IntegratedFunc(p float64) Integrated {
	sf := func(x, y float64) float64 {
		return 2 * math.Cos(x) / (1 - math.Pow(math.Sin(x), 2)*math.Pow(math.Cos(y), 2))
	}

	f := func(x, y float64) float64 {
		return 4 / math.Pi * (1 - math.Exp(-p*sf(x, y))) * math.Cos(x) * math.Sin(x)
	}

	return f
}

// Integrate function used to split integration directions.
func Integrate(f Integrated, lm [][]float64, n []int, fn []Integrator) float64 {
	inner := func(x float64) float64 {
		return fn[1](fWrap(f, x), lm[1][0], lm[1][1], n[1])
	}
	return fn[0](inner, lm[0][0], lm[0][1], n[0])
}

func fWrap(f func(float64, float64) float64, val float64) func(float64) float64 {
	return func(val1 float64) float64 {
		return f(val, val1)
	}
}

func pToV(p, a, b float64) float64 {
	return (b+a)/2 + (b-a)*p/2
}
