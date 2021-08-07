package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("ova-serial-api")

	var slice = []int{10, 20, 30, 40, 50}
	var dividedSlice = divideSlice(slice, 2)

	for _, x := range dividedSlice {
		fmt.Print("[ ")
		for _, y := range x {
			fmt.Print(y)
			fmt.Print(" ")
		}
		fmt.Println("]")
	}
}

func divideSlice(slice []int, batchSize int) (newSlice [][]int) {
	var newSize = int(math.Ceil(float64(len(slice)) / float64(batchSize)))
	remainder := len(slice) % batchSize
	newSlice = make([][]int, newSize)
	for i := 0; i < newSize; i++ {
		newBatchSize := batchSize
		if i == newSize-1 {
			newBatchSize = remainder
		}
		newSlice[i] = make([]int, newBatchSize)
		for j := 0; j < newBatchSize; j++ {
			newSlice[i][j] = slice[i*batchSize+j]
		}
	}
	return
}
