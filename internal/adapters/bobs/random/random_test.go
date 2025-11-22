package bobs

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRandomBobifier(t *testing.T) {
	b := NewRandomBobifier()
	assert.NotNil(t, b)
	assert.IsType(t, &RandomBobifier{}, b)
}

func TestBobify(t *testing.T) {
	input := "herp a derp"

	out1 := NewRandomBobifier().Bobify(input)
	out2 := NewRandomBobifier().Bobify(input)

	assert.Equal(t, input, strings.ToLower(out1))
	assert.Equal(t, input, strings.ToLower(out2))

	assert.NotEqual(t, input, out1)
	assert.NotEqual(t, input, out2)

	assert.NotEqual(t, out1, out2)
}
