package mtree

import (
	"fmt"
	"testing"
)

func encodingErr(str string) error {
	return fmt.Errorf("Encoding of '%s' does not align with vis(3) VIS_OCTAL encoding", str)
}
func invertibleErr(str string) error {
	return fmt.Errorf("Encoding of '%s' wasn't invertible", str)
}
func TestEncode(t *testing.T) {

	// not all symbols are encoded
	s := "hello there go-mtree users!"
	if Encode(s) != "hello\\040there\\040go-mtree\\040users!" {
		t.Fatal(encodingErr(" "))
	}
	if Decode(Encode(s)) != s {
		t.Fatal(invertibleErr(s))
	}

	s = "[ "
	if Encode(s) != "\\133\\040" {
		t.Fatal(encodingErr(s))
	}
	if len(Encode(s)) != (4 * len(s)) {
		t.Fatal(encodingErr(s))
	}
	if Decode(Encode(s)) != s {
		t.Fatal(invertibleErr(s))
	}
}
