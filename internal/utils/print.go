package utils

import "fmt"

func Print2dSlice(slice [][]int) {
	for _, subSlice := range slice {
		Print1dSlice(subSlice)
	}
}

func Print1dSlice(slice []int) {
	fmt.Print("[ ")
	for _, element := range slice {
		fmt.Print(element)
		fmt.Print(" ")
	}
	fmt.Println("]")
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
