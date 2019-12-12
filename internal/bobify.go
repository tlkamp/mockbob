package internal

import (
	"bufio"
	"os"
	"unicode"
)

func Bobify(input string, caps bool) string {
	cap := bool2int(!caps) // boolean inverted because 1 = true, and we're using modulus, so it will select the wrong character indicies if not inverted
	runes := []rune(input)

	for i := 0; i < len(runes); i++ {
		if unicode.IsSpace(runes[i]) {
			cap = (cap + 1) % 2 // change the value of cap to flip indicies if we encounter a space while processing
			continue
		}

		if i%2 == cap { // using cap this way allows us to keep code DRY and invert the boolean condition when we encounter spaces
			runes[i] = unicode.ToUpper(runes[i])
		}
	}
	return string(runes)
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
