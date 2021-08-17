package main

import (
	"fmt"
	"ova-serial-api/internal/config"
	"ova-serial-api/internal/model"
	"ova-serial-api/internal/utils"
	"time"
)

const configPath = "config/test_config.json"
const intervalSec = 10

func main() {
	fmt.Println("ova-serial-api")

	testTask2()
	testTask3()

	var cfg config.Config
	for {
		err := config.UpdateConfig(configPath, &cfg)
		if err != nil {
			fmt.Printf("Error occurred: %s\n", nil)
		} else {
			fmt.Printf("Config '%s' updated: %+v\n", configPath, cfg)
		}

		time.Sleep(intervalSec * time.Second)
	}
}

func testTask2() {
	slices := [][]int{
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
			fmt.Printf("%v\n", slice)
			dividedSlice := utils.SplitSlice(slice, batchSize)
			fmt.Printf("Slice divided into batches of size %d\n", batchSize)
			for _, subSlice := range dividedSlice {
				fmt.Printf("%v\n", subSlice)
			}
		}
	}

	var m map[string]int
	m = make(map[string]int)
	m["a"] = 1
	m["bc"] = 2
	fmt.Printf("%+v\n", m)
	newMap := utils.InvertStrIntMap(m)
	fmt.Printf("%+v\n", newMap)

	slices = [][]int{
		{-1, -3, 0, 4, 1, 6, -4},
		{-2, 2},
		{5, 1, 2},
		{5, 1, 2, 7},
	}
	forbiddenValues := []int{-2, 1, 2, 5}

	for _, slice := range slices {
		fmt.Println("Slice")
		fmt.Printf("%v\n", slice)
		filteredSlice := utils.FilterIntSlice(slice, forbiddenValues)
		fmt.Println("Slice filtered")
		fmt.Printf("%v\n", filteredSlice)
	}
}

func testTask3() {
	serial1 := model.Serial{UserID: 1, Title: "Firends", Genre: "comedy", Year: 1994, Seasons: 10}
	serial2 := model.Serial{UserID: 2, Title: "Game of Thrones", Genre: "fantasy", Year: 2011, Seasons: 8}
	serial3 := model.Serial{UserID: 3, Title: "Breaking Bad", Genre: "criminal", Year: 2008, Seasons: 5}
	serial4 := model.Serial{UserID: 4, Title: "The Big Bang Theory", Genre: "comedy", Year: 2007, Seasons: 12}

	serialSlices := [][]model.Serial{
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
			fmt.Printf("%v\n", slice)
			dividedSlice := utils.SplitSerialSlice(slice, batchSize)
			fmt.Printf("Slice divided into batches of size %d\n", batchSize)
			for _, subSlice := range dividedSlice {
				fmt.Printf("%v\n", subSlice)
			}
		}
	}

	serialsMap := utils.SerialSliceToMap([]model.Serial{
		serial1, serial2, serial3, serial4,
	})
	fmt.Printf("%v\n", serialsMap)
}
