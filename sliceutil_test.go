package goji

import (
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
