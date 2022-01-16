package usecase

import (
	"runtime"

	"github.com/cledupe/cross-commerce-sort-number/application/ports/input"
	"github.com/cledupe/cross-commerce-sort-number/application/ports/output"
	"github.com/cledupe/cross-commerce-sort-number/domain"
)

func InsertNumbers(dataInput input.InterfaceData, dataSave output.InterfaceData) {
	coreNumber := runtime.NumCPU()
	channels := make(chan []float64, coreNumber)
	page := 0
	valid := true
	var numbers []float64
	for valid {
		for channelSize := 0; channelSize < coreNumber; channelSize++ {
			page += 1
			go func(page int) {
				dataInput.GetData(page, channels)
			}(page)
		}

		for channelSize := 0; channelSize < coreNumber; channelSize++ {
			listNumber := <-channels
			if len(listNumber) == 0 {
				valid = false
				break
			} else {
				numbers = append(numbers, listNumber...)
			}
		}

	}
	listStNumber := domain.SetNumbers(numbers)
	listStNumber.SortList()
	dataSave.Save(listStNumber.GetAllNumbers())

}
