package meansquare

import (
	"fmt"
	"math"
)

// F64Matrix type used to represent float64 matrix.
type F64Matrix [][]float64

// SolveSLAE used to solve SLAE.
func SolveSLAE(ds DotSet, n int) F64Matrix {
	slae := MakeSLAE(ds, n)

	for i := 0; i < n+1; i++ {
		for j := 0; j < n+1; j++ {
			if i == j {
				continue
			}
			subCoeff := slae[j][i] / slae[i][i]
			for k := 0; k < n+2; k++ {
				slae[j][k] -= subCoeff * slae[i][k]
			}
		}
	}

	for i := 0; i < n+1; i++ {
		divider := slae[i][i]
		for j := 0; j < n+2; j++ {
			slae[i][j] /= divider
		}
	}

	return slae
}

// MakeSLAE used to make SLAE.
func MakeSLAE(ds DotSet, n int) F64Matrix {
	mat := make(F64Matrix, n+1)
	for i := range mat {
		mat[i] = make([]float64, n+2)
	}

	for i := 0; i < n+1; i++ {
		for j := 0; j < n+1; j++ {
			slaeCoeffs := 0.0
			expandedCoeff := 0.0
			for k := 0; k < len(ds); k++ {
				slaeCoeffs += ds[k].weight * math.Pow(ds[k].X, float64(i)) * math.Pow(ds[k].X, float64(j))
				expandedCoeff += ds[k].weight * ds[k].Y * math.Pow(ds[k].X, float64(i))
			}
			mat[i][j] = slaeCoeffs
			mat[i][n+1] = expandedCoeff
		}
	}

	return mat
}

// PrintMatrix used to print matrix in matrix form to standart output.
func (mat F64Matrix) PrintMatrix() {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat)+1; j++ {
			fmt.Printf("%15.1f ", mat[i][j])
		}
		fmt.Printf("\n")
	}
}
