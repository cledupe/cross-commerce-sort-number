package utils

import (
	"runtime"
	"sync"
)

func QuickSortMultiThread(arrayNumbers []float64) {
	if len(arrayNumbers) > 1 {
		var wg sync.WaitGroup
		wg.Add(2)
		first, last := 0, len(arrayNumbers)-1
		pivotIndex := partition(arrayNumbers, first, last)
		go func() {
			quickSort(arrayNumbers, first, pivotIndex-1)
			wg.Done()
		}()

		go func() {
			quickSort(arrayNumbers, pivotIndex+1, last)
			wg.Done()
		}()

		wg.Wait()
	}

}

func partition(numberArray []float64, first int, last int) int {
	pivotValue := numberArray[last]
	for j := first; j < last; j++ {
		if numberArray[j] < pivotValue {
			numberArray[j], numberArray[first] = numberArray[first], numberArray[j]
			first++
		}
	}
	numberArray[first], numberArray[last] = numberArray[last], numberArray[first]
	return first
}

func quickSort(numberArray []float64, first int, last int) {
	if first > last {
		return
	}
	pivotIndex := partition(numberArray, first, last)
	quickSort(numberArray, first, pivotIndex-1)
	quickSort(numberArray, pivotIndex+1, last)
}

func SortMultiThread(arrayNumbers []float64) []float64 {
	var result []float64
	coreNumber := runtime.NumCPU()
	channels := make(chan []float64, coreNumber)

	rangeValues := (len(arrayNumbers) + coreNumber - 1) / coreNumber
	coreUsed := 0

	for i := 0; i < len(arrayNumbers); i += rangeValues {
		end := i + rangeValues
		coreUsed += 1

		if end > len(arrayNumbers) {
			end = len(arrayNumbers)
		}
		go func(i int) {
			channels <- mergeSort(arrayNumbers[i:end])
		}(i)
	}

	for c := 0; c < coreUsed; c++ {
		result = merge(result, <-channels)
	}

	return result

}

func mergeSort(items []float64) []float64 {
	if len(items) < 2 {
		return items
	}
	part1 := mergeSort(items[:len(items)/2])
	part2 := mergeSort(items[len(items)/2:])
	return merge(part1, part2)
}

func merge(part1 []float64, part2 []float64) []float64 {
	result := []float64{}
	i := 0
	j := 0
	for i < len(part1) && j < len(part2) {
		if part1[i] < part2[j] {
			result = append(result, part1[i])
			i++
		} else {
			result = append(result, part2[j])
			j++
		}
	}
	for ; i < len(part1); i++ {
		result = append(result, part1[i])
	}
	for ; j < len(part2); j++ {
		result = append(result, part2[j])
	}
	return result
}
