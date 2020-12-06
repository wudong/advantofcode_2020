package day5

import "testing"
import "github.com/stretchr/testify/assert"

func TestDiv(t *testing.T) {
	low, high := div(0, 127, true)
	assert.Equal(t, low, 0)
	assert.Equal(t, high, 63)

	low, high = div(low, high, false)
	assert.Equal(t, low, 32)
	assert.Equal(t, high, 63)

	low, high = div(low, high, true)
	assert.Equal(t, low, 32)
	assert.Equal(t, high, 47)

	low, high = div(low, high, false)
	assert.Equal(t, low, 40)
	assert.Equal(t, high, 47)

	low, high = div(low, high, false)
	assert.Equal(t, low, 44)
	assert.Equal(t, high, 47)

	low, high = div(low, high, true)
	assert.Equal(t, low, 44)
	assert.Equal(t, high, 45)

	low, high = div(low, high, true)
	assert.Equal(t, low, 44)
	assert.Equal(t, high, 44)
}
