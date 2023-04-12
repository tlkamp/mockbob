package bobs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStandardBobifier(t *testing.T) {
	b := NewStandardBobifier(true)

	assert.NotNil(t, b)
	assert.True(t, b.startCaps)
}

func TestBobifyNoCaps(t *testing.T) {
	b := &Bobifier{}
	want := "hErPaDeRp"
	got := b.Bobify("herpaderp")
	if got != want {
		t.Errorf("got %s; want %s", got, want)
	}
}

func TestBobifyCaps(t *testing.T) {
	b := &Bobifier{startCaps: true}
	want := "HeRpAdErP"
	got := b.Bobify("herpaderp")
	if got != want {
		t.Errorf("got %s; want %s", got, want)
	}
}

func TestBobifySpaces(t *testing.T) {
	b := &Bobifier{}
	want := "dO yOu EvEn LiFt BrO"
	got := b.Bobify("do you even lift bro")
	if got != want {
		t.Errorf("got = %s; want %s", got, want)
	}
}
