package main

import (
	"fmt"
)

func main() {
	var constructMode int
	fmt.Println("Pls input the testCase mode(1:manual, 2: direct)")
	fmt.Scanf("%d", &constructMode)
	var testCase []int
	if constructMode == 1 {
		var arrayLen int
		fmt.Println("Pls type data length")
		fmt.Scanf("%d", &arrayLen)
		testCase = constructByCmd(arrayLen)
	} else if constructMode == 2 {
		testCase = constructByDirectTestCase()
	} else {
		fmt.Println("unsupprted testCase")
		return
	}
	fmt.Println("Before sort")
	fmt.Println(testCase)
	fmt.Println("After sort")
	quickSort(testCase, 0, len(testCase)-1)
	fmt.Println(testCase)
}

func constructByDirectTestCase() []int {
	return []int{3, 5, 38, 20, 18, 1, 2, 10, 8, 6, 4}
}

func constructByCmd(dataLength int) []int {
	var res = make([]int, dataLength)
	for i := 0; i < dataLength; i++ {
		tmp := 0
		fmt.Scanf("%d", &tmp)
		res[i] = tmp
	}
	return res
}

func quickSort(list []int, startIndex, endIndex int) {
	if startIndex >= endIndex {
		return
	}
	baseIndex := sort(list, startIndex, endIndex)
	quickSort(list, startIndex, baseIndex-1)
	quickSort(list, baseIndex+1, endIndex)

}

func sort(list []int, startIndex, endIndex int) int {
	baseValue := list[endIndex]
	swapIndex := startIndex
	for i := startIndex; i <= endIndex; i++ {
		if list[i] < baseValue {
			list[swapIndex], list[i] = list[i], list[swapIndex]
			swapIndex++
		}
	}
	list[endIndex], list[swapIndex] = list[swapIndex], list[endIndex]
	return swapIndex
}
