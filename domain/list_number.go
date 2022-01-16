package domain

import (
	"sync"
)

type ListStNumber struct {
	numbers []float64
}

func SetNumbers(listNumber []float64) *ListStNumber {
	var listStnumber ListStNumber
	listStnumber.numbers = listNumber
	return &listStnumber
}

func (listStNumber ListStNumber) GetAllNumbers() []float64 {
	return listStNumber.numbers
}

func (listStNumber ListStNumber) GetNumbers(itensPerPage int, pageNumber int) []float64 {
	startItemPage := itensPerPage * (pageNumber - 1)
	endItemPage := (itensPerPage * pageNumber)
	numberItems := len(listStNumber.numbers)

	if pageNumber > 0 && numberItems > startItemPage {
		if numberItems < endItemPage {
			endItemPage = numberItems
		}
		return listStNumber.numbers[startItemPage:endItemPage]
	}
	return []float64{}
}

func (listStNumber *ListStNumber) SortList() {
	listStNumber.quickSortMultiThread()
}

func (listStNumber *ListStNumber) quickSortMultiThread() {
	if len(listStNumber.numbers) > 1 {
		var wg sync.WaitGroup
		wg.Add(2)
		first, last := 0, len(listStNumber.numbers)-1
		pivotIndex := partition(listStNumber.numbers, first, last)
		go func() {
			quickSort(listStNumber.numbers, first, pivotIndex-1)
			wg.Done()
		}()

		go func() {
			quickSort(listStNumber.numbers, pivotIndex+1, last)
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
