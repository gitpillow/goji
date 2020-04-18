package goji

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestStringSliceContains(t *testing.T) {
	sli := []string{"a", "b", "c"}
	s := "a"
	contains := StringSliceContains(sli, s)
	assert.Equal(t, true, contains)
}
