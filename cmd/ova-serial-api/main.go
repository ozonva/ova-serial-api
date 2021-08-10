package main

import (
	"fmt"
	"ova-serial-api/internal/utils"
)

func main() {
	fmt.Println("ova-serial-api")

	var slices = [][]int{
		{10, 20, 30, 40, 50},
		{-1, -3, 0, 4, 1, 6, -4},
		{5, 1},
		{5, 1, 2},
		{5, 1, 2, 7},
	}
	batchSizeMax := 5
	batchSizeMin := 1

	for batchSize := batchSizeMin; batchSize < batchSizeMax; batchSize++ {
		for _, slice := range slices {
			fmt.Println("Slice")
			utils.Print1dSlice(slice)
			var dividedSlice = utils.DivideSlice(slice, batchSize)
			fmt.Printf("Slice divided into batches of size %d\n", batchSize)
			utils.Print2dSlice(dividedSlice)
		}
	}

	var m map[string]int
	m = make(map[string]int)
	m["a"] = 1
	m["bc"] = 2
	utils.PrintIntStrMap(m)
	newMap := utils.InvertMap(m)
	utils.PrintStrIntMap(newMap)

	slices = [][]int{
		{-1, -3, 0, 4, 1, 6, -4},
		{-2, 2},
		{5, 1, 2},
		{5, 1, 2, 7},
	}

	for _, slice := range slices {
		fmt.Println("Slice")
		utils.Print1dSlice(slice)
		var filteredSlice = utils.FilterSlice(slice)
		fmt.Println("Slice filtered")
		utils.Print1dSlice(filteredSlice)
	}
}
