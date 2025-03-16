package main

import (
		"fmt"
		"math/rand"
		)

func merge_sort(arr []int) []int {
	buf := make([]int, len(arr), len(arr))
	return _merge_sort(arr, buf)
}

func _merge_sort(arr []int, buf []int) []int {
	if len(arr) < 2 {
		return arr
	} else if len(arr) == 2 {
		a := arr[0]
		b := arr[1]
		if b < a {
			arr[0] = b
			arr[1] = a
		}
		return arr
	}

	pivot := len(arr)/2
	lower := _merge_sort(arr[:pivot], buf[:pivot])
	upper := _merge_sort(arr[pivot:], buf[pivot:])

	copy(buf[:pivot], lower)
	copy(buf[pivot:], upper)
	var i, j, k int = 0, pivot, 0
	var l, h int = 0, 0
	for i < pivot && j < len(buf) {
		l = buf[i]
		h = buf[j]
		if l <= h {
			arr[k] = l
			i++
		} else {
			arr[k] = h
			j++
		}
		k++
	}
	if i == pivot {
		for ; k < len(buf); k++ {
			arr[k] = buf[j]
			j++
		}
	} else {
		for ; k < len(buf); k++ {
			arr[k] = buf[i]
			i++
		}
	}

	return arr
}

func main() {
	n := 100
	a := make([]int, n, n)

	prefix := ""
	for i := 0; i < cap(a); i++ {
		a[i] = rand.Intn(n)
		fmt.Printf("%s%d", prefix, a[i])
		prefix = ", "
	}
	fmt.Println()

	merge_sort(a)
	prefix = ""
	for i := 0; i < len(a); i++ {
		fmt.Printf("%s%d", prefix, a[i])
		prefix = ", "
	}
	fmt.Println()
}
