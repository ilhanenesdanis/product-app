package infrasracture

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {

	t.Run("Test Add", func(t *testing.T) {
		actual := Add(10, 20)
		assert.Equal(t, 30, actual)
	})
}

func Add(x int, y int) int {
	return x + y
}
