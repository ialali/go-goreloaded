package goreloaded

import (
	"fmt"
	"regexp"
	"strings"
)

func ReplaceLowWithLowercase(text string) string {
	words := CustomSplit(text)
	for i, word := range words {
		if strings.Contains(word, "(low,") || strings.Contains(word, "(low, ") && i > 0 {
			var numWords int
			_, err := fmt.Sscanf(word, "(low,%d)", &numWords)
			if err == nil && i-numWords >= 0 {
				for j := i - numWords; j < i; j++ {
					words[j] = strings.ToLower(words[j])
				}
			}
			words[i] = ""
		} else if word == "(low)" && i > 0 {
			words[i-1] = strings.ToLower(words[i-1])
			words[i] = ""
		}
	}
	return strings.Join(words, " ")
}
func CustomSplit(line string) []string {
	// Use regular expressions to match words and content inside parentheses.
	re := regexp.MustCompile(`\([^\)]+\)|\S+`)
	return re.FindAllString(line, -1)
}
