package internal

import (
	"bufio"
	"os"
	"strings"
	"unicode"
)

func Bobify(input string, caps bool) string {
	cap := bool2int(!caps)
	spaces := spaceIndices(input)
	input = strings.ReplaceAll(strings.ToLower(input), " ", "")
	runes := []rune(input)

	// auto-invert the boolean
	for i := 0; i < len(runes); i++ {
		if i%2 == cap {
			runes[i] = unicode.ToUpper(runes[i])
		}
	}

	for i := 0; i < len(spaces); i++ {
		s := spaces[i]
		runes = append(runes, 0 /* use the zero value of the element type */)
		copy(runes[s+1:], runes[s:])
		runes[s] = rune(' ')
	}

	return string(runes)
}

func spaceIndices(input string) []int {
	runes := []rune(input)
	var ind []int

	for i := 0; i < len(runes); i++ {
		if unicode.IsSpace(runes[i]) {
			ind = append(ind, i)
		}
	}
	return ind
}

func bool2int(in bool) int {
	if in {
		return 1
	}

	return 0
}

func IsStdin() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}

func ReadStdin() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	text := ""

	if scanner.Err() != nil {
		return text, scanner.Err()
	}

	for scanner.Scan() {
		text += scanner.Text()
	}

	return text, nil
}
