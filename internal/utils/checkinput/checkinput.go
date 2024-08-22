package checkinput

var romans = map[string]struct{}{
	"I":    {},
	"II":   {},
	"III":  {},
	"IV":   {},
	"V":    {},
	"VI":   {},
	"VII":  {},
	"VIII": {},
	"IX":   {},
	"X":    {},
}

func IsRoman(s string) bool {
	if _, ok := romans[s]; !ok {
		return false
	}
	return true
}
