package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var Data *[]float64

//Mock de um storage
type MockDb struct {
}

func (mockDb MockDb) Load() ([]float64, error) {
	return []float64{1, 2, 3, 4, 5, 6}, nil
}
func (mockDb MockDb) Save(mock []float64) {
	Data = &mock
}

//Mock de um provedor de dados
type MockGetData struct {
}

func (mockGetData MockGetData) GetData(page int, result chan<- []float64) {
	if page > 1 {
		result <- []float64{}
	}
	result <- []float64{9, 8, 7, 6, 5, 4, 3, 2, 1}
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

func TestInsertion(t *testing.T) {
	assert := assert.New(t)
	mockDb := MockDb{}
	mockGetData := MockGetData{}
	testArray := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	InsertNumbers(mockGetData, mockDb)
	resultArray := *Data
	assert.Equal(testArray, resultArray, "The Two arrays should be the same")
}
