package time

func Annotate(lexemes []lexeme) []lexeme {
	/*
		Strategy:

		- First annotate obvious things: timezone, weekdays, month names.
		- Then find clock time, if any. These are very often digits separated by colons, and colons are usually not used to separate date elements.
		- Whatever digits are left should be date elements. We may be able to inspect the format of the input to see what order YMD come in. If not, we should not guess at what is what.
	*/

	// Annotate obvious things: timezones, weekdays, month names
	for i := range lexemes {
		if isTimezone(lexemes[i]) {
			lexemes[i].annotation = timezone
		} else if isMonth(lexemes[i]) {
			lexemes[i].annotation = monthLetters
		} else if isWeekday(lexemes[i]) {
			lexemes[i].annotation = dayLetters
		}
	}

	// TODO: fitf

	// Annotate clock time
	clockTimesSeen := 0
	clockAnnotation := func() lexemeAnnotation {
		var a lexemeAnnotation
		switch clockTimesSeen {
		case 0:
			a = hour
		case 1:
			a = minute
		case 2:
			a = second
		default:
			a = clockTime
		}
		clockTimesSeen++

		return a
	}

	for i, l := range lexemes {
		if l.isColon() {
			// Need a lexeme before or after colon
			if i == 0 || i == len(lexemes)-1 {
				continue
			}
			if lexemes[i-1].isDigit() && lexemes[i-1].annotation == unknown {
				lexemes[i-1].annotation = clockAnnotation()
			}
			if lexemes[i+1].isDigit() && lexemes[i+1].annotation == unknown {
				lexemes[i+1].annotation = clockAnnotation()
			}
		}
	}

	for i := range lexemes {
		// Need to have at least 2 more lexemes
		if len(lexemes)-i < 3 {
			break
		}
		if lexemes[i].annotation == second && lexemes[i+1].kind == dot && lexemes[i+2].isDigit() {
			lexemes[i+2].annotation = secondFraction
		}
	}

	// TODO: date elements

	return lexemes
}
