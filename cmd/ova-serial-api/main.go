package main

import (
	"fmt"
	"ova-serial-api/internal/config"
	"ova-serial-api/internal/model"
	"ova-serial-api/internal/utils"
	"time"
)

var cfg config.Config

func main() {
	fmt.Println("ova-serial-api")

	testTask2()
	testTask3()

	const configPath = "config/test_config.json"
	const intervalSec = 10

	for {
		err := config.UpdateConfig(configPath, &cfg)
		if err != nil {
			fmt.Printf("Error occurred: %s\n", nil)
		} else {
			fmt.Printf("Config '%s' updated: %s\n", configPath, cfg)
		}

		time.Sleep(intervalSec * time.Second)
	}
}

func testTask2() {
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
			utils.Print1dIntSlice(slice)
			var dividedSlice = utils.SplitSlice(slice, batchSize)
			fmt.Printf("Slice divided into batches of size %d\n", batchSize)
			utils.Print2dIntSlice(dividedSlice)
		}
	}

	var m map[string]int
	m = make(map[string]int)
	m["a"] = 1
	m["bc"] = 2
	utils.PrintIntStrMap(m)
	newMap := utils.InvertStrIntMap(m)
	utils.PrintStrIntMap(newMap)

	slices = [][]int{
		{-1, -3, 0, 4, 1, 6, -4},
		{-2, 2},
		{5, 1, 2},
		{5, 1, 2, 7},
	}

	for _, slice := range slices {
		fmt.Println("Slice")
		utils.Print1dIntSlice(slice)
		var filteredSlice = utils.FilterIntSlice(slice)
		fmt.Println("Slice filtered")
		utils.Print1dIntSlice(filteredSlice)
	}
}

func testTask3() {
	serial1 := model.Serial{UserID: 1, Title: "Firends", Genre: "comedy", Year: 1994, Seasons: 10}
	serial2 := model.Serial{UserID: 2, Title: "Game of Thrones", Genre: "fantasy", Year: 2011, Seasons: 8}
	serial3 := model.Serial{UserID: 3, Title: "Breaking Bad", Genre: "criminal", Year: 2008, Seasons: 5}
	serial4 := model.Serial{UserID: 4, Title: "The Big Bang Theory", Genre: "comedy", Year: 2007, Seasons: 12}

	var serialSlices = [][]model.Serial{
		{serial1, serial2},
		{serial1, serial2, serial3},
		{serial1, serial2, serial3, serial4},
		{serial1, serial2, serial3, serial4, serial1},
		{serial4},
	}

	batchSizeMax := 5
	batchSizeMin := 1

	for batchSize := batchSizeMin; batchSize < batchSizeMax; batchSize++ {
		for _, slice := range serialSlices {
			fmt.Println("Slice")
			utils.Print1dSerialSlice(slice)
			var dividedSlice = utils.SplitSerialSlice(slice, batchSize)
			fmt.Printf("Slice divided into batches of size %d\n", batchSize)
			utils.Print2dSerialSlice(dividedSlice)
		}
	}

	serialsMap := utils.SerialSliceToMap([]model.Serial{
		serial1, serial2, serial3, serial4,
	})
	utils.PrintUintSerialMap(serialsMap)
}
