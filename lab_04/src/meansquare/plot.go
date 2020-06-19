package meansquare

import (
	"fmt"
	"image/color"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func convertDots(ds DotSet) plotter.XYs {
	var conv plotter.XYs

	for _, d := range ds {
		cd := plotter.XY{
			X: d.X,
			Y: d.Y,
		}
		conv = append(conv, cd)
	}

	return conv
}

// DrawPlot used to draw plot by approximated dots.
func DrawPlot(ds, approx DotSet) {
	p, err := plot.New()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	p.Title.Text = "Approximation using meansquare method"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	p.Add(plotter.NewGrid())

	conv := convertDots(approx)
	def := convertDots(ds)

	l, err := plotter.NewLine(conv)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Color = color.RGBA{B: 255, A: 255}

	s, err := plotter.NewScatter(def)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}

	p.Add(l, s)
	p.Legend.Add("approximated", l)
	p.Legend.Add("data", s)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}
