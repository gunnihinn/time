package time

import (
	"testing"
)

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
