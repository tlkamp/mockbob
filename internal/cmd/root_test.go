package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestIsStdin(t *testing.T) {
	mockStdin("data")
	want := true
	got := isStdin()
	if got != want {
		t.Errorf("IsStdin(), true = %v; want %v", got, want)
	}
}

func TestReadStdin(t *testing.T) {
	mockData := "data"
	mockStdin(mockData)
	want := mockData
	got, err := readStdin()
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
