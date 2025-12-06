package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsStdin(t *testing.T) {
	mockStdin("data")
	assert.True(t, isStdin())
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
