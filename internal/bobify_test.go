package internal

import (
	"testing"
)

func TestBobifyNoCaps(t *testing.T) {
	want := "hErPaDeRp"
	got := Bobify("herpaderp", false)
	if got != want {
		t.Errorf("Bobify(\"herpaderp\", false) = %s; want %s", got, want)
	}
}

func TestBobifyCaps(t *testing.T) {
	want := "HeRpAdErP"
	got := Bobify("herpaderp", true)
	if got != want {
		t.Errorf("Bobify(\"herpaderp\", true = %s; want %s", got, want)
	}
}

func TestBobifySpaces(t *testing.T) {
	want := "dO yOu EvEn LiFt BrO"
	got := Bobify("do you even lift bro", false)
	if got != want {
		t.Errorf("Bobify(\"do you even lift bro\", true = %s; want %s", got, want)
	}
}
