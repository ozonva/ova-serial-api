package utils

import (
	"ova-serial-api/internal/model"
)

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

func FilterIntSlice(slice []int, forbiddenValues []int) (filtered []int) {
	forbiddenMap := intSliceToMap(forbiddenValues)
	for _, element := range slice {
		if _, exists := forbiddenMap[element]; !exists {
			filtered = append(filtered, element)
		}
	}
	return
}
