package mtree

import (
	"bytes"
	"fmt"
	"unicode"
)

// graphic can contain letters
// look into unicode.RangeTable, figure out which Ranges need to be encoded

// Encode does this
func Encode(str string) string {
	characters := []byte(str)
	result := bytes.Buffer{}
	for _, c := range characters {
		r := rune(c)
		switch {
		case unicode.IsSpace(r):
			fallthrough
		case unicode.IsLetter(r):
			fallthrough
		case unicode.IsControl(r):
			result.WriteString(fmt.Sprintf("\\%03o", c))
		default:
			result.WriteByte(c)
		}
	}
	return result.String()
}

// Decode does this
func Decode(str string) string {
	return str
}
