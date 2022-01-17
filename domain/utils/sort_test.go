package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortMultiThread(t *testing.T) {
	assert := assert.New(t)
	testArray := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	resultArray := SortMultiThread([]float64{1, 3, 4, 6, 7, 9, 8, 5, 2})
	assert.Equal(testArray, resultArray, "The Two arrays should be the same")
}
