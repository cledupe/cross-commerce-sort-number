package domain

import "github.com/cledupe/cross-commerce-sort-number/domain/utils"

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
	listStNumber.numbers = utils.SortMultiThread(listStNumber.numbers)
}
