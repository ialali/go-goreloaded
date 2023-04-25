package goreloaded

import (
	"bytes"
	"strings"
)

func ReplaceQuotes(text string) string {
	words := strings.Fields(text)
	var buffer bytes.Buffer
	for i, word := range words {
		if word == "'" && i > 0 && i < len(words)-1 {
			prevWord := words[i-1]
			nextWord := words[i+1]
			if len(prevWord) > 0 && prevWord[len(prevWord)-1] != '\'' {
				buffer.Truncate(buffer.Len() - 1)
				buffer.WriteString("'")
			}
			buffer.WriteString(prevWord)
			buffer.WriteString("'")
			if len(nextWord) > 0 && nextWord[0] != '\'' {
				buffer.WriteString("'")
			}
			buffer.WriteString(nextWord)
			buffer.WriteString(" ")
			continue
		}
		buffer.WriteString(word)
		buffer.WriteString(" ")
	}
	return strings.TrimSpace(buffer.String())
}
