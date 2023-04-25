package goreloaded

import (
	"fmt"
	"strings"
)

func ReplaceUpWithUppercase(text string) string {
	words := CustomSplit(text)
	for i, word := range words {
		if strings.Contains(word, "(up,") || strings.Contains(word, "(up, )") && i > 0 {
			var numWords int
			_, err := fmt.Sscanf(word, "(up,%d)", &numWords)
			if err == nil && i-numWords >= 0 {
				for j := i - numWords; j < i; j++ {
					words[j] = strings.ToUpper(words[j])
				}
			}
			words[i] = ""
		} else if word == "(up)" && i > 0 {
			words[i-1] = strings.ToUpper(words[i-1])
			words[i] = ""
		}
	}
	return strings.Join(words, " ")
}
