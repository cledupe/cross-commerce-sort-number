package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert := assert.New(t)
	numberArray := []float64{0.333333, 0.22222222, 0.55555555, 0.66, 0.11, 0.99}
	testedArray := []float64{0.11, 0.22222222, 0.333333, 0.55555555, 0.66, 0.99}
	listStNumber := SetNumbers(numberArray)
	listStNumber.SortList()
	assert.Equal(testedArray, listStNumber.GetAllNumbers(), "The Two arrays should be the same")
}
