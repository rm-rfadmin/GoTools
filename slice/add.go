package slice

func AddSliceIndexV1(slice []int, index int, value int) []int {
	if index < 0 || index > len(slice) {
		return nil
	}

	s1 := slice[:index]

	new_slice := append([]int{}, s1...) // 创建新切片

	new_slice = append(new_slice, value) // 添加元素

	new_slice = append(new_slice, slice[index:]...) // TODO: 性能偏弱，子切片的创建和拷贝

	return new_slice
}

func AddSliceIndexV2[T any](slice []T, index int, value T) []T {
	if index < 0 || index > len(slice) {
		return nil
	}

	s1 := slice[:index]

	new_slice := append([]T{}, s1...) // 创建新切片

	new_slice = append(new_slice, value) // 添加元素

	new_slice = append(new_slice, slice[index:]...) // TODO: 性能偏弱，子切片的创建和拷贝

	return new_slice
}
