package sorting

import "fmt"

func quicksort(arr *[]int, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)
		fmt.Println(*arr)
		quicksort(arr, low, pivotIndex)
		quicksort(arr, pivotIndex+1, high)
	}
}

func partition(arr *[]int, low, high int) int {
	pivot := (*arr)[low]
	// 12, 1, 15, 6
	i, j := low, high
	for i < j {
		for (*arr)[i] <= pivot && i < high {
			i++
		}
		for (*arr)[j] > pivot && j > low {
			j--
		}
		if i < j {
			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
		}

	}
	(*arr)[j], (*arr)[low] = (*arr)[low], (*arr)[j]
	return j
}
