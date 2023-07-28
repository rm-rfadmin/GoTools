package slice

func ReversalSlice(slice []int) []int {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func ReversalSliceV1[T any](slice []T) []T {
	lenNum := len(slice)
	if lenNum == 0 {
		return slice
	}

	for i, j := 0, lenNum-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
