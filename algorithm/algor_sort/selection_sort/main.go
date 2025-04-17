package main

import "fmt"

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
	selectionSort(testCase)
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

func selectionSort(list []int) {
	maxVal := -1
	maxIndex := -1
	for i := len(list) - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			if list[j] > maxVal {
				maxIndex = j
				maxVal = list[j]
			}
		}
		if maxIndex != -1 {
			tmp := list[i]
			list[i] = maxVal
			list[maxIndex] = tmp
			maxVal = -1
			maxIndex = -1
		}
	}
}
