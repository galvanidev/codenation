package caesar

import (
	"bytes"
)

type Caesar struct {
	numberPositions int
}

func NewCaesar(numberPositions int) Caesar {
	caesar := Caesar{}
	caesar.numberPositions = numberPositions
	return caesar
}

func (caesar Caesar) Decipher(cipher string) string {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8,
		"i": 9, "j": 10, "k": 11, "l": 12, "m": 13, "n": 14, "o": 15,
		"p": 16, "q": 17, "r": 18, "s": 19, "t": 20, "u": 21, "v": 22,
		"w": 23, "x": 24, "y": 25, "z": 26}

	var b bytes.Buffer

	for _, char := range cipher {
		v, exist := m[string(char)]
		if exist {
			if (v - caesar.numberPositions) < 0 {
				v = len(m) - (caesar.numberPositions - v)
			} else {
				v = v - caesar.numberPositions
			}
			b.WriteString(stringValueOf(v))
		} else {
			b.WriteString(string(char))
		}
	}
	return b.String()
}

func stringValueOf(i int) string {
	characters := "abcdefghijklmnopqrstuvwxyz"
	return string(characters[i-1])
}