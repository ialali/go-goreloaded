package goreloaded

import "strings"

func FormatVowels(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words)-1; i++ {
		if strings.ToLower(words[i]) == "a" {
			nextWord := strings.ToLower(words[i+1])
			if strings.HasPrefix(nextWord, "a") ||
				strings.HasPrefix(nextWord, "e") ||
				strings.HasPrefix(nextWord, "i") ||
				strings.HasPrefix(nextWord, "o") ||
				strings.HasPrefix(nextWord, "u") ||
				(strings.HasPrefix(nextWord, "h") && len(nextWord) > 1 && !strings.HasPrefix(words[i+1], "hour")) {
				words[i] = "an"
			}
		}
	}
	return strings.Join(words, " ")
}
 