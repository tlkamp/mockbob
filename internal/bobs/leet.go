package bobs

var mapping = map[rune]rune{
	'a': '4',
	'A': '4',
	'e': '3',
	'E': '3',
	'l': '1',
	'L': '1',
	'o': '0',
	'O': '0',
	's': '5',
	'S': '5',
	't': '7',
	'T': '7',
}

type LeetBobifier struct{}

func NewLeetBobifier() *LeetBobifier {
	return &LeetBobifier{}
}

func (l *LeetBobifier) Bobify(input string) string {
	runes := []rune(input)

	for i := 0; i < len(runes); i++ {
		if val, ok := mapping[runes[i]]; ok {
			runes[i] = val
		}
	}

	return string(runes)
}
