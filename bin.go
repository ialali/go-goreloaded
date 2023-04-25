package goreloaded

import (
	"strconv"
	"strings"
)

func ReplaceBinWithDec(text string) string {
	words := strings.Fields(text)
	for i, word := range words {
		if word == "(bin)" && i > 0 {
			prevWord := strings.TrimPrefix(words[i-1], "0b")
			decNum, err := strconv.ParseInt(prevWord, 2, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(decNum, 10)
			}
		}
	}
	replacedText := strings.Join(words, " ")
	replacedText = strings.ReplaceAll(replacedText, "(bin)", "")
	return replacedText
}
