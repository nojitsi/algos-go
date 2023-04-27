package main

func SelectionSort(array []int) []int {

	for i := 0; i < len(array); i++ {
		minIndex := i

		for j := i; j < len(array); j++ {
			if array[j] < array[minIndex] {
				minIndex = j
			}
		}

		replaceableValue := array[i]
		array[i] = array[minIndex]
		array[minIndex] = replaceableValue
	}

	return array
}
