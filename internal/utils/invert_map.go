package utils

func InvertMap(mapToInvert map[string]int) map[int]string {
	var newMap map[int]string
	newMap = make(map[int]string)
	for key, value := range mapToInvert {
		newMap[value] = key
	}
	return newMap
}
