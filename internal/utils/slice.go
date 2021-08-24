package utils

import (
	"ova-serial-api/internal/model"
)

func SplitSlice(slice []int, batchSize uint) (newSlice [][]int) {
	if batchSize == 0 {
		batchSize = 1
	}
	remainder := len(slice) % int(batchSize)
	newSlice = [][]int{}
	for i := 0; i < len(slice); i += int(batchSize) {
		if i+int(batchSize) <= len(slice) {
			newSlice = append(newSlice, slice[i:i+int(batchSize)])
		}
	}
	if remainder > 0 {
		newSlice = append(newSlice, slice[len(slice)-remainder:])
	}
	return
}

func SplitSerialSlice(serials []model.Serial, batchSize uint) [][]model.Serial {
	if batchSize == 0 {
		batchSize = 1
	}
	remainder := len(serials) % int(batchSize)
	newSlice := make([][]model.Serial, 0)
	for i := 0; i < len(serials); i += int(batchSize) {
		if i+int(batchSize) <= len(serials) {
			newSlice = append(newSlice, serials[i:i+int(batchSize)])
		}
	}
	if remainder > 0 {
		newSlice = append(newSlice, serials[len(serials)-remainder:])
	}
	return newSlice
}

func FilterIntSlice(slice []int, forbiddenValues []int) (filtered []int) {
	forbiddenMap := intSliceToMap(forbiddenValues)
	filtered = make([]int, 0)
	for _, element := range slice {
		if _, exists := forbiddenMap[element]; !exists {
			filtered = append(filtered, element)
		}
	}
	return
}
