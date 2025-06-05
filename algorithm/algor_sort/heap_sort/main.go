package main

import "fmt"

func main() {
	arr := []int{4, 10, 3, 5, 1}
	fmt.Println("original array:", arr)

	heapSort(arr)
	fmt.Println("array heapSort:", arr)
}

func heapSort(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		heapify(arr, i, len(arr))
	}

	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, 0, i)
	}
}

func heapify(arr []int, index, len int) {
	largest := index
	leftPos := 2*index + 1
	rightPos := 2*index + 2
	if leftPos < len && arr[leftPos] > arr[largest] {
		largest = leftPos
	}
	if rightPos < len && arr[rightPos] > arr[largest] {
		largest = rightPos
	}
	if largest != index {
		arr[index], arr[largest] = arr[largest], arr[index]
		heapify(arr, largest, len)
	}
}
