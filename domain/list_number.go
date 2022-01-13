package domain

type ListStNumber struct {
	numbers []float32
}

func SetNumbers(listNumber []float32) *ListStNumber {
	var listStnumber ListStNumber
	listStnumber.numbers = listNumber
	return &listStnumber
}

func (listStNumber *ListStNumber) Join(listJoin ListStNumber) {
	if !listJoin.isEmpty() {
		listStNumber.numbers = append(listStNumber.numbers, listJoin.numbers...)
	}
}

func (listStNumber ListStNumber) isEmpty() bool {
	if len(listStNumber.numbers) == 0 {
		return true
	} else {
		return false
	}
}

func (listStNumber ListStNumber) GetAllNumbers() []float32 {
	return listStNumber.numbers
}

func (listStNumber *ListStNumber) SortList() {

}
