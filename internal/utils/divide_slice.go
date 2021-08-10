package utils

func DivideSlice(slice []int, batchSize int) (newSlice [][]int) {
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
