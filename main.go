package main

import "fmt"

func quickSort(arr []int, first int, last int) {
	if first < last {
		p := partition(arr, first, last)
		quickSort(arr, first, p-1)
		quickSort(arr, p+1, last)
	}
}

func partition(arr []int, first int, last int) int {
	pivot := arr[last]
	i := first - 1
	for j := first; j < last; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	i++
	arr[i], arr[last] = pivot, arr[i]
	return i

}

func main() {
	var arr = []int{4, 123412, 5, 3, 1234, 7, 6, 345, 345345, 321, -2}
	n := len(arr) - 1
	quickSort(arr, 0, n)
	fmt.Println(arr)
}
