package time

import (
	"testing"
)

func TestAnnotateClock(t *testing.T) {
	inputs := [][]lexeme{
		[]lexeme{
			lexeme{value: "15", kind: digit},
			lexeme{value: ":", kind: colon},
			lexeme{value: "04", kind: digit},
			lexeme{value: ":", kind: colon},
			lexeme{value: "05", kind: digit},
		},
		[]lexeme{
			lexeme{value: "15", kind: digit},
			lexeme{value: ":", kind: colon},
			lexeme{value: "04", kind: digit},
		},
		[]lexeme{
			lexeme{value: "15", kind: digit},
			lexeme{value: ":", kind: colon},
			lexeme{value: "04", kind: digit},
			lexeme{value: ":", kind: colon},
			lexeme{value: "05", kind: digit},
			lexeme{value: ".", kind: dot},
			lexeme{value: "123", kind: digit},
		},
		[]lexeme{
			lexeme{value: "15", kind: digit},
			lexeme{value: ":", kind: colon},
			lexeme{value: "04", kind: digit},
			lexeme{value: ":", kind: colon},
			lexeme{value: "05", kind: digit},
			lexeme{value: ".", kind: dot},
		},
	}

	expecteds := [][]lexeme{
		[]lexeme{
			lexeme{value: "15", kind: digit, annotation: hour},
			lexeme{value: ":", kind: colon},
			lexeme{value: "04", kind: digit, annotation: minute},
			lexeme{value: ":", kind: colon},
			lexeme{value: "05", kind: digit, annotation: second},
		},
		[]lexeme{
			lexeme{value: "15", kind: digit, annotation: hour},
			lexeme{value: ":", kind: colon},
			lexeme{value: "04", kind: digit, annotation: minute},
		},
		[]lexeme{
			lexeme{value: "15", kind: digit, annotation: hour},
			lexeme{value: ":", kind: colon},
			lexeme{value: "04", kind: digit, annotation: minute},
			lexeme{value: ":", kind: colon},
			lexeme{value: "05", kind: digit, annotation: second},
			lexeme{value: ".", kind: dot},
			lexeme{value: "123", kind: digit, annotation: secondFraction},
		},
		[]lexeme{
			lexeme{value: "15", kind: digit, annotation: hour},
			lexeme{value: ":", kind: colon},
			lexeme{value: "04", kind: digit, annotation: minute},
			lexeme{value: ":", kind: colon},
			lexeme{value: "05", kind: digit, annotation: second},
			lexeme{value: ".", kind: dot},
		},
	}

	for i, input := range inputs {
		got := Annotate(input)
		expected := expecteds[i]

		if !equal(got, expected) {
			t.Errorf("Input '%s':\nGot:     '%v'\nExpected '%v'", input, got, expected)
		}
	}
}

func TestAnnotateTimezone(t *testing.T) {
	input := []lexeme{
		lexeme{value: "GMT", kind: letter},
	}

	expected := []lexeme{
		lexeme{value: "GMT", kind: letter, annotation: timezone},
	}
	got := Annotate(input)

	if !equal(got, expected) {
		t.Errorf("Input '%s':\nGot:     '%v'\nExpected '%v'", input, got, expected)
	}
}

func TestAnnotateWeekday(t *testing.T) {
	inputs := []string{
		"mon",
		"tue",
		"wed",
		"thu",
		"fri",
		"sat",
		"sun",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}

	for _, str := range inputs {
		input := []lexeme{
			lexeme{value: str, kind: letter},
		}

		expected := []lexeme{
			lexeme{value: str, kind: letter, annotation: dayLetters},
		}
		got := Annotate(input)

		if !equal(got, expected) {
			t.Errorf("Input '%s':\nGot:     '%v'\nExpected '%v'", input, got, expected)
		}
	}
}

func TestAnnotateMonthday(t *testing.T) {
	inputs := []string{
		"Jan",
		"Feb",
		"Mar",
		"Apr",
		"May",
		"Jun",
		"Jul",
		"Aug",
		"Sep",
		"Oct",
		"Nov",
		"Dec",
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	for _, str := range inputs {
		input := []lexeme{
			lexeme{value: str, kind: letter},
		}

		expected := []lexeme{
			lexeme{value: str, kind: letter, annotation: monthLetters},
		}
		got := Annotate(input)

		if !equal(got, expected) {
			t.Errorf("Input '%s':\nGot:     '%v'\nExpected '%v'", input, got, expected)
		}
	}
}
