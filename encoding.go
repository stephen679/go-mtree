package mtree

import (
	"bytes"
	"fmt"
	"strconv"
	"unicode"
)

// Encode is a function that encodes a string in a similar manner to vis(3),
// specifically that it is a simpler version and does not cover all characters
// that need encoding. However, the set of characters that are encoded in this
// function cover most of commonly used characters. For now it serves its purpose in
// a general sense. Eventually needs to be improved.
func Encode(str string) string {
	characters := []byte(str)
	result := bytes.Buffer{}
	for _, c := range characters {
		r := rune(c)
		switch {
		case !unicode.IsGraphic(r): // default
			fallthrough
		case c == '*' || c == '?' || c == '[' || c == '#': // VIS_GLOB
			fallthrough
		case unicode.IsSpace(r): // VIS_WHITE
			fallthrough
		case unicode.IsControl(r): // VIS_SAFE
			result.WriteString(fmt.Sprintf("\\%03o", c))
		default:
			if c == '\\' {
				result.WriteByte('\\')
			}
			result.WriteByte(c) // don't encode
		}
	}
	return result.String()
}

// Decode decodes a string that potentially contains encoded characters that were
// encoded by Encode above. encodeLen is the length of an encoded character resulting from
// using the Encode function above
func Decode(str string) string {
	result := bytes.Buffer{}
	characters := []byte(str)
	i := 0
	for i < len(characters) {
		c := characters[i]
		if c == '\\' {
			if i+1 < len(characters) && characters[i+1] == '\\' {
				result.WriteByte(c)
			} else if i+3 < len(characters) {
				octal, err := octalStrToByte(string(characters[i+1 : i+4]))
				if err != nil {
					result.WriteByte(c)
				} else {
					result.WriteByte(octal)
					i += 4
					continue
				}
			} else {
				result.WriteByte(c)
			}
		} else {
			result.WriteByte(c)
		}
		i++
	}
	return result.String()
}

// octalByteToString takes in a sring that attempts to represent a character in
// octal format
func octalStrToByte(str string) (byte, error) {
	result := 0
	numChars := len(str)
	for i, r := range str {
		octalNum, err := strconv.Atoi(string(r))
		if err != nil {
			return ' ', err
		}
		result += (octalNum << uint(numChars*(numChars-1-i)))
	}
	return byte(result), nil
}
