package slice

import "fmt"

func CheckSliceIndex[T any](slice []T, index int) bool {
	length := len(slice)
	if index < 0 || index >= length {
		return false
	}
	return true
}

func DeleteSliceIndex[T any](slice []T, index int) ([]T, error) {
	if !CheckSliceIndex(slice, index) {
		return nil, fmt.Errorf("下标超出范围，长度 %d, 下标 %d", len(slice), index)
	} else {

		copy(slice[index:], slice[index+1:]) // TODO: 优化，不需要拷贝，直接覆盖

		newLen := len(slice) - 1
		slice = slice[:len(slice)-1]

		if newLen < cap(slice)/2 || newLen > 1024 {
			newSlice := make([]T, newLen) // TODO：缩容后的容量应该比长度大，否则会重新分配内存，以后再优化
			copy(newSlice, slice)
			slice = newSlice
		}

		return slice, nil
	}
}
