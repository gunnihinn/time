package time

import (
	"testing"
)

func TestIsTimezone(t *testing.T) {
	input := lexeme{value: "GMT", kind: letter}

	if !isTimezone(input) {
		t.Errorf("Input '%s':\nShould be timezone", input)
	}
}

func TestIsWeekday(t *testing.T) {
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
		input := lexeme{value: str, kind: letter}

		if !isWeekday(input) {
			t.Errorf("Input '%s':\nShould be weekday", input)
		}
	}
}

func TestIsMonth(t *testing.T) {
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
		input := lexeme{value: str, kind: letter}

		if !isMonth(input) {
			t.Errorf("Input '%s':\nShould be month", input)
		}
	}
}
