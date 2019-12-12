package internal

import (
	"io/ioutil"
	"log"
	"os"
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
		t.Errorf("Bobify(\"herpaderp\", true) = %s; want %s", got, want)
	}
}

func TestBobifySpaces(t *testing.T) {
	want := "dO yOu EvEn LiFt BrO"
	got := Bobify("do you even lift bro", false)
	if got != want {
		t.Errorf("Bobify(\"do you even lift bro\", false) = %s; want %s", got, want)
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
