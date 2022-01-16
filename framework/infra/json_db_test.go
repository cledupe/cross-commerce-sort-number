package infra

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveAndLoad(t *testing.T) {
	assert := assert.New(t)
	var jsonDb JsonDb
	testArray := []float64{1, 2, 3, 4, 5, 6}
	jsonDb.Save(testArray)
	arrayNumber, _ := jsonDb.Load()
	assert.Equal(testArray, arrayNumber, "The Two arrays should be the same")
}
