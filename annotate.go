package time

func Annotate(lexemes []lexeme) []lexeme {
	for i := range lexemes {
		if isTimezone(lexemes[i]) {
			lexemes[i].annotation = timezone
		} else if isMonth(lexemes[i]) {
			lexemes[i].annotation = monthLetters
		} else if isWeekday(lexemes[i]) {
			lexemes[i].annotation = dayLetters
		}
	}

	return lexemes
}
