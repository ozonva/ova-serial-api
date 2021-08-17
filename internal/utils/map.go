package utils

import (
	"ova-serial-api/internal/model"
)

func InvertStrIntMap(mapToInvert map[string]int) map[int]string {
	var newMap map[int]string
	newMap = make(map[int]string)
	for key, value := range mapToInvert {
		newMap[value] = key
	}
	return newMap
}

func IntSliceToMap(slice []int) map[int]bool {
	var newMap map[int]bool
	newMap = make(map[int]bool)
	for _, element := range slice {
		newMap[element] = true
	}
	return newMap
}

func SerialSliceToMap(serials []model.Serial) map[uint64]model.Serial {
	newMap := make(map[uint64]model.Serial, len(serials))
	for _, serial := range serials {
		newMap[serial.UserID] = serial
	}
	return newMap
}
