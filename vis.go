package mtree

import (
	"bytes"
	"fmt"
	"unicode"
)

// Encode does this
func Encode(str string) string {
	characters := []byte(str)
	result := bytes.Buffer{}
	for _, c := range characters {
		if unicode.IsSpace(rune(c)) || unicode.IsControl(rune(c)) ||
			unicode.IsSymbol(rune(c)) || unicode.IsGraphic(rune(c)) {
			result.WriteString(fmt.Sprintf("\\%#o", c))
		} else {
			result.WriteByte(c)
		}
	}
	return result.String()
}

// Decode does this
func Decode(str string) string {
	return str
}
