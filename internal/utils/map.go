package utils

import (
	"fmt"
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

func SerialSliceToMap(serials []model.Serial) map[uint64]model.Serial {
	newMap := make(map[uint64]model.Serial, len(serials))
	for _, serial := range serials {
		newMap[serial.UserID] = serial
	}
	return newMap
}

func PrintIntStrMap(mapToPrint map[string]int) {
	fmt.Print("{ ")
	for key, value := range mapToPrint {
		fmt.Printf("%s: %d ", key, value)
	}
	fmt.Println("}")
}

func PrintStrIntMap(mapToPrint map[int]string) {
	fmt.Print("{ ")
	for key, value := range mapToPrint {
		fmt.Printf("%d: %s ", key, value)
	}
	fmt.Println("}")
}

func PrintUintSerialMap(mapToPrint map[uint64]model.Serial) {
	fmt.Println("{ ")
	for key, value := range mapToPrint {
		fmt.Printf(" %d: %s \n", key, value.String())
	}
	fmt.Println("}")
}
