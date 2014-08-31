package main
import (
	"strings"
	"strconv"
)


func Underscore(camelcase string) string {
	returnstring := ""
	for _, letter := range camelcase {
		if IsConsonant(letter) {
			returnstring += "_"
			runetostring, _ := strconv.Unquote(strconv.QuoteRuneToASCII(letter))
			returnstring += strings.ToLower(runetostring)
		} else {
			runetostring, _ := strconv.Unquote(strconv.QuoteRuneToASCII(letter))
			returnstring += runetostring
		}
	}
	return strings.Trim(returnstring, "_") 
}

func IsConsonant(s rune) bool {
	return strings.ContainsRune("ABCDEFGHIJKLMNOPQRSTUVWXYZ", s)
}


