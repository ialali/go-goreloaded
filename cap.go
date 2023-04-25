package goreloaded

import (
	"fmt"
	"strings"
)

func ReplaceCapWithCapitalized(text string) string {
	words := CustomSplit(text)
	for i, word := range words {
		if strings.Contains(word, "(cap,") || strings.Contains(word, "(cap, ") && i > 0 {
			var numWords int
			_, err := fmt.Sscanf(word, "(cap,%d)", &numWords)
			if err == nil && i-numWords >= 0 {
				for j := i - numWords; j < i; j++ {
					words[j] = strings.Title(words[j])
				}
			}
			words[i] = ""
		} else if word == "(cap)" && i > 0 {
			words[i-1] = strings.Title(words[i-1])
			words[i] = ""
		}
	}
	return strings.Join(words, " ")
}
