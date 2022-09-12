package main

import (
	"fmt"
	"strings"
)

func MeanMedian(arrayInput []float64) (float64, float64) {
	var mean, median float64

	for i := 0; i < len(arrayInput); i++ {
		mean += arrayInput[i]
	}
	mean = mean / float64(len(arrayInput))

	var data float64 = float64(len(arrayInput)+1) / 2
	s := fmt.Sprintln(" ", data)
	if strings.Contains(s, ".5") {
		t := data - 1
		median = (arrayInput[int(t-0.5)] + arrayInput[int(t+0.5)]) / 2
	} else {
		median = arrayInput[int(data+1)/2]
	}
	// data := make([]float64, len(arrayInput))
	// l := len(data)
	// if l == 0 {
	// 	return 0
	// } else if l%2 == 0 {
	// 	median = (data[l/2-1] + data[l/2]) / 2
	// } else {
	// 	median = data[l/2]
	// }
	return mean, median
}

func main() {
	fmt.Println(MeanMedian([]float64{1, 2, 3, 4}))          // 2.5 2.5
	fmt.Println(MeanMedian([]float64{1, 2, 3, 4, 5}))       // 3 3
	fmt.Println(MeanMedian([]float64{7, 8, 9, 13, 15}))     // 10.4 9
	fmt.Println(MeanMedian([]float64{10, 20, 30, 40, 50}))  // 30 30
	fmt.Println(MeanMedian([]float64{15, 20, 30, 60, 120})) // 49 30
}
