package utils

import (
	"ova-serial-api/internal/model"
)

func InvertStrIntMap(mapToInvert map[string]int) map[int]string {
	newMap := make(map[int]string)
	for key, value := range mapToInvert {
		newMap[value] = key
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

func intSliceToMap(slice []int) map[int]struct{} {
	newMap := make(map[int]struct{})
	for _, element := range slice {
		newMap[element] = struct{}{}
	}
	return newMap
}
