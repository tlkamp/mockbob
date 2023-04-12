package bobs

import (
	"math/rand"
	"time"
	"unicode"
)

// RandomBobifier is an implementation of nondeterministic bobified strings.
type RandomBobifier struct{}

func NewRandomBobifier() *RandomBobifier {
	return &RandomBobifier{}
}

func (r *RandomBobifier) Bobify(input string) string {
	s := rand.NewSource(time.Now().UnixNano())
	ran := rand.New(s)
	runes := []rune(input)

	for i := 0; i < len(runes); i++ {
		if unicode.IsSpace(runes[i]) {
			continue
		}

		// Random number 1-100
		if randInt := ran.Intn(100); randInt <= 50 {
			continue
		}

		runes[i] = unicode.ToUpper(runes[i])
	}

	return string(runes)
}
