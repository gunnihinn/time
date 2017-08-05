package time

import (
	"strconv"
	"strings"
)

func isTimezone(l lexeme) bool {
	timezones := map[string]bool{
		// wikipedia.org/wiki/List_of_time_zone_abbreviations | uniq
		// Because of course time zone abbreviations aren't unique.
		"acdt":  true,
		"acst":  true,
		"act":   true,
		"adt":   true,
		"aedt":  true,
		"aest":  true,
		"aft":   true,
		"akdt":  true,
		"akst":  true,
		"amst":  true,
		"amt":   true,
		"art":   true,
		"ast":   true,
		"awst":  true,
		"azost": true,
		"azot":  true,
		"azt":   true,
		"bdt":   true,
		"biot":  true,
		"bit":   true,
		"bot":   true,
		"brst":  true,
		"brt":   true,
		"bst":   true,
		"btt":   true,
		"cat":   true,
		"cct":   true,
		"cdt":   true,
		"cest":  true,
		"cet":   true,
		"chadt": true,
		"chast": true,
		"chost": true,
		"chot":  true,
		"chst":  true,
		"chut":  true,
		"cist":  true,
		"cit":   true,
		"ckt":   true,
		"clst":  true,
		"clt":   true,
		"cost":  true,
		"cot":   true,
		"cst":   true,
		"ct":    true,
		"cvt":   true,
		"cwst":  true,
		"cxt":   true,
		"davt":  true,
		"ddut":  true,
		"dft":   true,
		"easst": true,
		"east":  true,
		"eat":   true,
		"ect":   true,
		"edt":   true,
		"eest":  true,
		"eet":   true,
		"egst":  true,
		"egt":   true,
		"eit":   true,
		"est":   true,
		"fet":   true,
		"fjt":   true,
		"fkst":  true,
		"fkt":   true,
		"fnt":   true,
		"galt":  true,
		"gamt":  true,
		"get":   true,
		"gft":   true,
		"gilt":  true,
		"git":   true,
		"gmt":   true,
		"gst":   true,
		"gyt":   true,
		"hadt":  true,
		"haec":  true,
		"hast":  true,
		"hkt":   true,
		"hmt":   true,
		"hovst": true,
		"hovt":  true,
		"ict":   true,
		"idt":   true,
		"iot":   true,
		"irdt":  true,
		"irkt":  true,
		"irst":  true,
		"ist":   true,
		"jst":   true,
		"kgt":   true,
		"kost":  true,
		"krat":  true,
		"kst":   true,
		"lhst":  true,
		"lint":  true,
		"magt":  true,
		"mart":  true,
		"mawt":  true,
		"mdt":   true,
		"mest":  true,
		"met":   true,
		"mht":   true,
		"mist":  true,
		"mit":   true,
		"mmt":   true,
		"msk":   true,
		"mst":   true,
		"mut":   true,
		"mvt":   true,
		"myt":   true,
		"nct":   true,
		"ndt":   true,
		"nft":   true,
		"npt":   true,
		"nst":   true,
		"nt":    true,
		"nut":   true,
		"nzdt":  true,
		"nzst":  true,
		"omst":  true,
		"orat":  true,
		"pdt":   true,
		"pet":   true,
		"pett":  true,
		"pgt":   true,
		"phot":  true,
		"pht":   true,
		"pkt":   true,
		"pmdt":  true,
		"pmst":  true,
		"pont":  true,
		"pst":   true,
		"pyst":  true,
		"pyt":   true,
		"ret":   true,
		"rott":  true,
		"sakt":  true,
		"samt":  true,
		"sast":  true,
		"sbt":   true,
		"sct":   true,
		"sdt":   true,
		"sgt":   true,
		"slst":  true,
		"sret":  true,
		"srt":   true,
		"sst":   true,
		"syot":  true,
		"taht":  true,
		"tft":   true,
		"tha":   true,
		"tjt":   true,
		"tkt":   true,
		"tlt":   true,
		"tmt":   true,
		"tot":   true,
		"trt":   true,
		"tvt":   true,
		"ulast": true,
		"ulat":  true,
		"usz1":  true,
		"utc":   true,
		"uyst":  true,
		"uyt":   true,
		"uzt":   true,
		"vet":   true,
		"vlat":  true,
		"volt":  true,
		"vost":  true,
		"vut":   true,
		"wakt":  true,
		"wast":  true,
		"wat":   true,
		"west":  true,
		"wet":   true,
		"wit":   true,
		"wst":   true,
		"yakt":  true,
		"yekt":  true,
	}

	return searchStringStackExact(l, timezones)
}

func isMonth(l lexeme) bool {
	shortMonths := []string{
		"jan",
		"feb",
		"mar",
		"apr",
		"jun",
		"jul",
		"aug",
		"sep",
		"oct",
		"nov",
		"dec",
	}

	return searchStringStack(l, shortMonths)
}

func isWeekday(l lexeme) bool {
	shortDays := []string{
		"mon",
		"tue",
		"wed",
		"thu",
		"fri",
		"sat",
		"sun",
	}

	return searchStringStack(l, shortDays)
}

func searchStringStack(needle lexeme, haystack []string) bool {
	if !needle.isLetter() {
		return false
	}

	value := strings.ToLower(needle.value)
	for _, hay := range haystack {
		if strings.Contains(value, hay) {
			return true
		}
	}

	return false
}

func searchStringStackExact(needle lexeme, haystack map[string]bool) bool {
	if !needle.isLetter() {
		return false
	}

	value := strings.ToLower(needle.value)
	_, found := haystack[value]

	return found
}

// maybeMonth can return false positives, but not false negatives.
func maybeMonth(l lexeme) bool {
	if l.isDigit() {
		d, err := strconv.Atoi(l.value)
		if err != nil {
			return false
		}
		return 1 <= d && d <= 12
	} else if l.isLetter() {
		return isMonth(l)
	}

	return false
}

// maybeDay can return false positives, but not false negatives.
func maybeDay(l lexeme) bool {
	if l.isDigit() {
		d, err := strconv.Atoi(l.value)
		if err != nil {
			return false
		}
		return 1 <= d && d <= 31
	}

	return false
}
