package channeler

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChanneler(t *testing.T) {
	b := new(bytes.Buffer)
	b.WriteString("test\n123")

	c := Channeler(b)

	r := make([]string, 0, 2)

	for s := range c {
		r = append(r, s)
	}

	expected := []string{"test\n", "123"}

	assert.Equal(t, 2, len(r))
	assert.True(t, reflect.DeepEqual(r, expected))
}
