package goji

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringSliceContains(t *testing.T) {
	sli := []string{"a", "b", "c"}
	s := "a"
	contains := StringSliceContains(sli, s)
	assert.Equal(t, true, contains)

	sli = []string{"a", "b", "c"}
	s = "d"
	contains = StringSliceContains(sli, s)
	assert.Equal(t, false, contains)
}

func TestSliceForEach(t *testing.T) {
	sli := []string{"a", "bb", "ccc"}
	charCount := 0
	SliceForEach(sli, func(i int, item interface{}) error {
		charCount += len(item.(string))
		return nil
	})
	assert.Equal(t, 6, charCount)

	sli2 := "aaa"
	err := SliceForEach(sli2, func(i int, item interface{}) error {
		return nil
	})
	assert.NotNil(t, err)
	fmt.Println(err)
}

func TestSliceForEachBackward(t *testing.T) {
	sli := []string{"a", "bb", "ccc"}
	charCount := 0
	SliceForEachBackward(sli, func(i int, item interface{}) error {
		charCount += len(item.(string))
		fmt.Println(i, item)
		return nil
	})
	assert.Equal(t, 6, charCount)

	sli2 := "aaa"
	err := SliceForEachBackward(sli2, func(i int, item interface{}) error {
		return nil
	})
	assert.NotNil(t, err)
	fmt.Println(err)
}