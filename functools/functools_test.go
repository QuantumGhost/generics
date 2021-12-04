package temp

import (
	"testing"
	"strconv"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	intArr := []int{1, 2, 3}
	mappedArr := Map[int, int](intArr, func(i int) int {
		return i+1
	})
	assert.Equal(t, []int{2, 3, 4}, mappedArr)
}


func TestMapErr(t *testing.T) {
	strArr := []string{"1", "2", "3"}
	intArr, err := MapErr[string, int](strArr, strconv.Atoi)
	assert.NoError(t, err)
	assert.Equal(t, []int{1, 2, 3}, intArr)

	strArr = []string{"1", "2", "a"}
	intArr, err = MapErr[string, int](strArr, strconv.Atoi)
	assert.Error(t, err)
	assert.Equal(t, []int(nil), intArr)
}


func TestReduce(t *testing.T) {
	intArr := []int{1, 2, 3}
	total := Reduce(intArr, 0, func (a, b int) int {
		return a + b
	})
	assert.Equal(t, 6, total)
}

func TestFilter(t *testing.T) {
	intArr := []int{1, 2, 3}
	filtered := Filter(intArr, func(i int) bool {
		return i % 2 == 0
	})
	assert.Equal(t, []int{2}, filtered)
}
