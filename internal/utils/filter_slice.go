package utils

var FORBIDDEN_VALUES = map[int]bool{-2: true, 1: true, 2: true, 5: true}

func FilterSlice(slice []int) (filtered []int) {
	for i := range slice {
		if !FORBIDDEN_VALUES[slice[i]] {
			filtered = append(filtered, slice[i])
		}
	}
	return
}
