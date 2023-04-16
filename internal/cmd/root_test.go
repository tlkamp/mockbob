package cmd

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsStdin(t *testing.T) {
	mockStdin("data")
	assert.True(t, isStdin())
}

func TestChanneler(t *testing.T) {
	b := new(bytes.Buffer)
	b.WriteString("test\n123")

	c := channeler(b)

	r := make([]string, 0, 2)

	for s := range c {
		r = append(r, s)
	}

	expected := []string{"test\n", "123"}

	assert.Equal(t, 2, len(r))
	assert.True(t, reflect.DeepEqual(r, expected))
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
