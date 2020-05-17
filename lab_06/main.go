package main

import "./differentiate"

func main() {
	x := []float64{1, 2, 3, 4, 5, 6}
	y := []float64{0.571, 0.889, 1.091, 1.231, 1.333, 1.412}
	h := 1.0

	differentiate.FmtPrintInit("X              :", x)
	differentiate.FmtPrintInit("Y              :", y)
	differentiate.FmtPrintRes("Onesided       :", differentiate.LeftDiff(y, h))
	differentiate.FmtPrintRes("Center         :", differentiate.CenterDiff(y, h))
	differentiate.FmtPrintRes("Second Runge   :", differentiate.SecondRungeDiff(y, h, 1))
	differentiate.FmtPrintRes("Aligned params :", differentiate.AlignedCoeffsDiff(x, y))
	differentiate.FmtPrintRes("Second onesided:", differentiate.SecondLeftDiff(y, h))
}
