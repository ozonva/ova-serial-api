package utils

import (
	"fmt"
	"ova-serial-api/internal/model"
)

var FORBIDDEN_VALUES = map[int]bool{-2: true, 1: true, 2: true, 5: true}

func SplitSlice(slice []int, batchSize int) (newSlice [][]int) {
	remainder := len(slice) % batchSize
	newSlice = [][]int{}
	for i := 0; i < len(slice); i += batchSize {
		if i+batchSize <= len(slice) {
			newSlice = append(newSlice, slice[i:i+batchSize])
		}
	}
	if remainder > 0 {
		newSlice = append(newSlice, slice[len(slice)-remainder:])
	}
	return
}

func SplitSerialSlice(serials []model.Serial, batchSize int) [][]model.Serial {
	remainder := len(serials) % batchSize
	var newSlice [][]model.Serial
	for i := 0; i < len(serials); i += batchSize {
		if i+batchSize <= len(serials) {
			newSlice = append(newSlice, serials[i:i+batchSize])
		}
	}
	if remainder > 0 {
		newSlice = append(newSlice, serials[len(serials)-remainder:])
	}
	return newSlice
}

func FilterIntSlice(slice []int) (filtered []int) {
	for i := range slice {
		if !FORBIDDEN_VALUES[slice[i]] {
			filtered = append(filtered, slice[i])
		}
	}
	return
}

func Print2dIntSlice(slice [][]int) {
	for _, subSlice := range slice {
		Print1dIntSlice(subSlice)
	}
}

func Print1dIntSlice(slice []int) {
	fmt.Print("[ ")
	for _, element := range slice {
		fmt.Print(element)
		fmt.Print(" ")
	}
	fmt.Println("]")
}

func Print2dSerialSlice(slice [][]model.Serial) {
	for _, subSlice := range slice {
		Print1dSerialSlice(subSlice)
	}
}

func Print1dSerialSlice(slice []model.Serial) {
	fmt.Println("[ ")
	for _, element := range slice {
		fmt.Print(" " + element.String())
		fmt.Println(" ")
	}
	fmt.Println("]")
}
