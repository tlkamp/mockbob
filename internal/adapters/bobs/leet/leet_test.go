package bobs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLeetBobifier(t *testing.T) {
	b := NewLeetBobifier()
	assert.NotNil(t, b)
	assert.IsType(t, &LeetBobifier{}, b)
}

func TestLeetBobify(t *testing.T) {
	b := NewLeetBobifier()

	tests := []struct {
		input    string
		expected string
	}{
		{"THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG", "7H3 QUICK BR0WN F0X JUMP5 0V3R 7H3 14ZY D0G"},
		{"the quick brown fox jumps over the lazy dog", "7h3 quick br0wn f0x jump5 0v3r 7h3 14zy d0g"},
		{"Testing 123", "7357ing 123"},
		{"Spongebob", "5p0ng3b0b"},
	}

	for _, tt := range tests {
		result := b.Bobify(tt.input)
		assert.Equal(t, tt.expected, result)
	}
}
