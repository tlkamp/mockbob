package bobs

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestBobifyNoCaps(t *testing.T) {
	b := &Bobifier{}
	want := "hErPaDeRp"
	got := b.Bobify("herpaderp")
	if got != want {
		t.Errorf("got %s; want %s", got, want)
	}
}

func TestBobifyCaps(t *testing.T) {
	b := &Bobifier{StartCaps: true}
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

func TestIsStdin(t *testing.T) {
	mockStdin("data")
	want := true
	got := IsStdin()
	if got != want {
		t.Errorf("IsStdin(), true = %v; want %v", got, want)
	}
}

func TestReadStdin(t *testing.T) {
	mockData := "data"
	mockStdin(mockData)
	want := mockData
	got, err := ReadStdin()
	if err != nil {
		t.Error(err)
	}
	if got != want {
		t.Errorf("ReadStdin(), true = %v; want %v", got, want)
	}
}

func mockStdin(content string) {
	data := []byte(content)
	tmpfile, err := ioutil.TempFile("", "mockbob")

	if err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Write(data); err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	os.Stdin = tmpfile
}
