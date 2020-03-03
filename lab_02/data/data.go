package main

import (
	"fmt"
	"os"
	"strconv"

	"../interp"
)

func main() {
	if len(os.Args) < 5 {
		fmt.Printf("USAGE: data <start> <end> <step> <datafile>\n")
		os.Exit(1)
	}

	f, err := os.Create(os.Args[4])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	start, _ := strconv.ParseFloat(os.Args[1], 64)
	end, _ := strconv.ParseFloat(os.Args[2], 64)
	step, _ := strconv.ParseFloat(os.Args[3], 64)

	data := genData(start, end, step)
	ds := len(data)

	f.WriteString(fmt.Sprintf("%d\n", ds))
	for i := 0; i < ds; i++ {
		for j := 0; j < ds; j++ {
			f.WriteString(fmt.Sprintf("%8.2f", data[i][j]))
		}
		f.WriteString("\n")
	}

	f.Close()
}

func genData(start, end, step float64) interp.FTable {
	size := int((end + start) / step)
	data := make(interp.FTable, size)
	for i := range data {
		data[i] = make([]float64, size)
	}

	for i := 1; i < size; i++ {
		data[0][i] = start + step*float64(i-1)
		data[i][0] = start + step*float64(i-1)
	}

	for i := 1; i < size; i++ {
		for j := 1; j < size; j++ {
			data[i][j] = mathFunc(data[0][i], data[j][0])
		}
	}

	return data
}

func mathFunc(x, y float64) float64 {
	return 2 * x * y
}
