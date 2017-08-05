package time

import (
	"unicode"
)

type lexemeType int

const (
	digit  lexemeType = iota // [0-9]+
	letter                   // [A-Za-z]+
	slash                    // /
	colon                    // :
	dash                     // -
	space                    // ' '
	dot                      // .
	other
)

type lexemeAnnotation int

const (
	unknown        lexemeAnnotation = iota
	year                            // [0-9]+ (no BC/AD)
	month                           // 1-12
	monthLetters                    // Jan, ...
	day                             // 1-31
	dayLetters                      // Mon, ..
	hour                            // 1-24
	minute                          // 1-59
	second                          // 1-60 [sic]
	secondFraction                  // [0-9]+
	clockTime                       // we know it's on a clock, but not exactly what it is
	fitf                            // AM/PM
	timezone                        // GMT, ...
)

func (l lexeme) isDigit() bool {
	return l.kind == digit
}

func (l lexeme) isLetter() bool {
	return l.kind == letter
}

func (l lexeme) isColon() bool {
	return l.kind == colon
}

func getType(r rune) lexemeType {
	switch r {
	case '/':
		return slash
	case ':':
		return colon
	case '-':
		return dash
	case ' ':
		return space
	case '.':
		return dot
	}

	if unicode.IsDigit(r) {
		return digit
	} else if unicode.IsLetter(r) {
		return letter
	} else {
		return other
	}
}

type lexeme struct {
	value      string
	kind       lexemeType
	annotation lexemeAnnotation
}

type lexerState int

const (
	lexDigit lexerState = iota
	lexLetter
	lexSingle
)

func Lex(timestamp string) []lexeme {
	runes := []rune(timestamp)
	parts := make([]lexeme, 0)
	state := lexSingle
	position := 0
	head := 0

LEXER:
	for position < len(runes) {
		switch state {

		case lexSingle:
			r := runes[position]

			if unicode.IsLetter(r) {
				state = lexLetter
			} else if unicode.IsDigit(r) {
				state = lexDigit
			} else {
				parts = append(parts, lexeme{
					value: string(r),
					kind:  getType(r),
				})
				position++
				head++
			}

		case lexLetter:
			r := runes[head]
			for unicode.IsLetter(r) {
				if head == len(runes)-1 {
					parts = append(parts, lexeme{
						value: string(runes[position : head+1]),
						kind:  letter,
					})
					break LEXER
				} else {
					head++
					r = runes[head]
				}
			}
			parts = append(parts, lexeme{
				value: string(runes[position:head]),
				kind:  letter,
			})
			position = head
			state = lexSingle

		case lexDigit:
			r := runes[head]
			for unicode.IsDigit(r) {
				if head == len(runes)-1 {
					parts = append(parts, lexeme{
						value: string(runes[position : head+1]),
						kind:  digit,
					})
					break LEXER
				} else {
					head++
					r = runes[head]
				}
			}
			parts = append(parts, lexeme{
				value: string(runes[position:head]),
				kind:  digit,
			})
			position = head
			state = lexSingle
		}
	}

	return parts
}
