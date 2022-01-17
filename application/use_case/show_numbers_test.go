package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockDb struct {
}

func (mockDb MockDb) Load() ([]float64, error) {
	return []float64{1, 2, 3, 4, 5, 6}, nil
}
func (mockDb MockDb) Save(mock []float64) {

}

func TestPagination(t *testing.T) {
	assert := assert.New(t)
	mockDb := MockDb{}
	testArray := []float64{2}
	resultArray := ShowNumbers(1, 2, mockDb)
	assert.Equal(testArray, resultArray, "The Two arrays should be the same")
}

func TestPaginationFirstElement(t *testing.T) {
	assert := assert.New(t)
	mockDb := MockDb{}
	testArray := []float64{1}
	resultArray := ShowNumbers(1, 1, mockDb)
	assert.Equal(testArray, resultArray, "The Two arrays should be the same")
}

func TestPaginationLastElement(t *testing.T) {
	assert := assert.New(t)
	mockDb := MockDb{}
	testArray := []float64{6}
	resultArray := ShowNumbers(1, 6, mockDb)
	assert.Equal(testArray, resultArray, "The Two arrays should be the same")
}

func TestPaginationOutOfElement(t *testing.T) {
	assert := assert.New(t)
	mockDb := MockDb{}
	testArray := []float64{}
	resultArray := ShowNumbers(1, 7, mockDb)
	assert.Equal(testArray, resultArray, "The Two arrays should be the same")
}
