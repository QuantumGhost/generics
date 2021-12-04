package iter

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterSlice(t *testing.T) {
	s := []int{1, 2, 3}
	iter := Slice[int](s)
	collected := []int{}
	for {
		v, err := iter.Next()
		if errors.Is(err, ErrNoMore) {
			break
		}
		collected = append(collected, v)
	}
	assert.Equal(t, s, collected)
}

func TestIterMap(t *testing.T) {
	makeTestMap := func() map[int]string {
		return map[int]string{
			1: "1",
			2: "2",
			3: "3",
		}
	}

	{
		collected := map[int]string{}
		m := makeTestMap()
		iter := Map[int, string](m)
		for {
			v, err := iter.Next()
			if errors.Is(err, ErrNoMore) {
				break
			}
			collected[v.Key] = v.Value
		}
		assert.Equal(t, m, collected)
	}

	{
		collected := map[int]string{}
		m := makeTestMap()
		iter := Map[int, string](m)
		delete(m, 3)
		for {
			v, err := iter.Next()
			if errors.Is(err, ErrNoMore) {
				break
			}
			collected[v.Key] = v.Value
		}
		expectedAfterDel := map[int]string{
			1: "1", 2: "2",
		}
		assert.Equal(t, expectedAfterDel, collected)
	}

	{
		collected := map[int]string{}
		m := makeTestMap()
		iter := Map[int, string](m)
		m[4] = "4"
		for {
			v, err := iter.Next()
			if errors.Is(err, ErrNoMore) {
				break
			}
			collected[v.Key] = v.Value
		}
		expectedAfterAdd := map[int]string{
			1: "1", 2: "2", 3: "3",
		}
		assert.Equal(t, expectedAfterAdd, collected)
	}
}
