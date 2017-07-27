package time

import (
	"testing"
	"unicode"
)

func equal(a, b []lexeme) bool {
	if len(a) != len(b) {
		return false
	}

	for i, x := range a {
		if x != b[i] {
			return false
		}
	}

	return true
}

func TestLexer(t *testing.T) {
	inputs := []string{
		"2006-1-2",
		"Mon, January 2 2006 15:04:05 -0700",
		"Mon, January 02, 2006, 15:04:05 MST",
		"2.1.2006 15:04:05",
		"",
		":",
	}
	expecteds := [][]lexeme{
		// "2006-1-2",
		[]lexeme{
			lexeme{value: "2006", kind: digit},
			lexeme{value: "-", kind: dash},
			lexeme{value: "1", kind: digit},
			lexeme{value: "-", kind: dash},
			lexeme{value: "2", kind: digit},
		},
		// "Mon, January 2 2006 15:04:05 -0700",
		[]lexeme{
			lexeme{value: "Mon", kind: letter},
			lexeme{value: ",", kind: other},
			lexeme{value: " ", kind: space},
			lexeme{value: "January", kind: letter},
			lexeme{value: " ", kind: space},
			lexeme{value: "2", kind: digit},
			lexeme{value: " ", kind: space},
			lexeme{value: "2006", kind: digit},
			lexeme{value: " ", kind: space},
			lexeme{value: "15", kind: digit},
			lexeme{value: ":", kind: colon},
			lexeme{value: "04", kind: digit},
			lexeme{value: ":", kind: colon},
			lexeme{value: "05", kind: digit},
			lexeme{value: " ", kind: space},
			lexeme{value: "-", kind: dash},
			lexeme{value: "0700", kind: digit},
		},
		// "Mon, January 02, 2006, 15:04:05 MST",
		[]lexeme{
			lexeme{value: "Mon", kind: letter},
			lexeme{value: ",", kind: other},
			lexeme{value: " ", kind: space},
			lexeme{value: "January", kind: letter},
			lexeme{value: " ", kind: space},
			lexeme{value: "02", kind: digit},
			lexeme{value: ",", kind: other},
			lexeme{value: " ", kind: space},
			lexeme{value: "2006", kind: digit},
			lexeme{value: ",", kind: other},
			lexeme{value: " ", kind: space},
			lexeme{value: "15", kind: digit},
			lexeme{value: ":", kind: colon},
			lexeme{value: "04", kind: digit},
			lexeme{value: ":", kind: colon},
			lexeme{value: "05", kind: digit},
			lexeme{value: " ", kind: space},
			lexeme{value: "MST", kind: letter},
		},
		// "2.1.2006 15:04:05",
		[]lexeme{
			lexeme{value: "2", kind: digit},
			lexeme{value: ".", kind: dot},
			lexeme{value: "1", kind: digit},
			lexeme{value: ".", kind: dot},
			lexeme{value: "2006", kind: digit},
			lexeme{value: " ", kind: space},
			lexeme{value: "15", kind: digit},
			lexeme{value: ":", kind: colon},
			lexeme{value: "04", kind: digit},
			lexeme{value: ":", kind: colon},
			lexeme{value: "05", kind: digit},
		},
		// ""
		[]lexeme{},
		// ":"
		[]lexeme{
			lexeme{value: ":", kind: colon},
		},
	}

	for i, input := range inputs {
		expected := expecteds[i]
		got := Lex(input)
		if !equal(got, expected) {
			t.Errorf("Input '%s':\nGot:     '%v'\nExpected '%v'", input, got, expected)
		}
	}
}

func TestIsLetter(t *testing.T) {
	notletters := []rune{
		'.',
		':',
		'/',
		' ',
		'-',
		'0',
		'1',
		'?',
		'!',
		'\\',
	}

	for _, nl := range notletters {
		if unicode.IsLetter(nl) {
			t.Error("'%s' shouldn't be letter", nl)
		}
	}
}
