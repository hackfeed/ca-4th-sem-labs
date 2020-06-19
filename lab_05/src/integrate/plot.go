package integrate

import (
	"fmt"
	"image/color"
	"math/rand"
	"os"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func genLabel(n, m, md1, md2 int) string {
	var f1s, f2s string

	if md1 == 0 {
		f1s = "Gauss"
	} else {
		f1s = "Simpson"
	}

	if md2 == 0 {
		f2s = "Gauss"
	} else {
		f2s = "Simpson"
	}

	return fmt.Sprintf("N = %v, M = %v, Methods = %v-%v", n, m, f1s, f2s)
}

func genRandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// GenDots used to generate dots for plot representation.
func GenDots(f func(p float64) float64, sc []float64) plotter.XYs {
	ds := plotter.XYs{}
	for i := sc[0]; i < sc[2]; i += sc[1] {
		d := plotter.XY{
			X: i,
			Y: f(i),
		}
		ds = append(ds, d)
	}

	return ds
}

// CreatePlot used to create base plot object.
func CreatePlot(title, x, y string) *plot.Plot {
	pl, err := plot.New()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	pl.Title.Text = title
	pl.X.Label.Text = x
	pl.Y.Label.Text = y
	pl.Add(plotter.NewGrid())

	return pl
}

// DrawPlot used to draw plot with given dots.
func DrawPlot(p *plot.Plot, ds plotter.XYs, n, m, md1, md2 int) {

	l, err := plotter.NewLine(ds)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Color = color.RGBA{B: uint8(genRandomNumber(0, 255)), A: uint8(genRandomNumber(0, 255))}

	p.Add(l)
	p.Legend.Add(genLabel(n, m, md1, md2), l)
}

// SavePlot used to save gonum plot to file.
func SavePlot(p *plot.Plot, f string) {
	if err := p.Save(8*vg.Inch, 4*vg.Inch, f); err != nil {
		panic(err)
	}
}
