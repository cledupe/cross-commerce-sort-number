package usecase

import (
	"github.com/cledupe/cross-commerce-sort-number/application/ports/output"
	"github.com/cledupe/cross-commerce-sort-number/domain"
)

/*
	Fluxo para Listar dados na etapa Load
*/
func ShowNumbers(itensPerPage int, pageNumber int, storageData output.InterfaceData) []float64 {
	numbersLoaded, err := storageData.Load()
	if err != nil {
		return nil
	}
	listStNumber := domain.SetNumbers(numbersLoaded)
	return listStNumber.GetNumbers(itensPerPage, pageNumber)

}
