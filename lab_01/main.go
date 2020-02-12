package main

import (
	"fmt"
	"io"
	"os"
)

type dot struct {
	X float64
	Y float64
}

type cords []dot

func main() {
	var x dot
	var n int
	c := cords{}
	f, err := os.Open("dots.txt")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	readDotsFromFile(f, &c, &x, &n)
	pos := c.insertByX(x)
	fmt.Println(c, x, n, pos)
}

func readDotsFromFile(f *os.File, c *cords, x *dot, n *int) {
	var dots int
	fmt.Fscanln(f, &dots)
	for i := 0; i < dots; i++ {
		var d dot
		_, err := fmt.Fscanln(f, &d.X, &d.Y)
		if err == io.EOF {
			break
		}
		*c = append(*c, d)
	}
	fmt.Fscanln(f, &x.X)
	fmt.Fscanln(f, n)
}

func (c cords) isFuncAscending() bool {
	for i := 0; i < len(c)-1; i++ {
		if c[i].X > c[i+1].X {
			return true
		}
	}
	return false
}

func (c *cords) insertByX(d dot) int {
	var t dot
	*c = append(*c, t)
	isAsc := c.isFuncAscending()
	if isAsc == true && d.X < (*c)[0].X || isAsc == false && d.X > (*c)[0].X {
		copy((*c)[1:], (*c)[0:])
		(*c)[0] = d
		return 0
	}
	if isAsc == true && d.X > (*c)[len(*c)-2].X || isAsc == false && d.X < (*c)[len(*c)-2].X {
		(*c)[len(*c)-1] = d
		return len(*c) - 1
	}
	for i := 1; i < len(*c); i++ {
		if isAsc == true && d.X > (*c)[i-1].X && d.X < (*c)[i].X ||
			isAsc == false && d.X < (*c)[i-1].X && d.X > (*c)[i].X {
			copy((*c)[i+1:], (*c)[i:])
			(*c)[i] = d
			return i
		}
	}
	return 0
}
