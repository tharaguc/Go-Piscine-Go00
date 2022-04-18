package main

import "fmt"

func main() {
	var orig = []int{0, 1, 2, 3, 4, 5}
	var ret []int

	ret = subSlice(orig, 0, 3, 3)
	fmt.Printf("ret = %v, len = %d, cap = %d\n", ret, len(ret), cap(ret))
	ret = subSlice(orig, 2, 7, 10)
	fmt.Printf("ret = %v, len = %d, cap = %d\n", ret, len(ret), cap(ret))
	ret = subSlice(orig, 2, 7, 3)
	fmt.Printf("ret = %v, len = %d, cap = %d\n", ret, len(ret), cap(ret))
}

func subSlice(slice []int, begin int, length int, capacity int) []int {
	if capacity < length {
		capacity = length
	}
	ret := make([]int, length, capacity)
	i := 0
	for i < (len(slice) - begin) && i < length {
		ret[i] = slice[begin + i]
		i++
	}
	return ret
}