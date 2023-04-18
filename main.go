package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Read input file
	input, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}
	text := string(input)

	// Apply transformations
	text = replaceHexWithDec(text)
	text = replaceBinWithDec(text)
	text = replaceLowWithLowercase(text)
	text = replaceUpWithUppercase(text)
	text = replaceCapWithCapitalized(text)
	text = formatPunctuation(text)
	text = replaceQuotes(text)

	// Write output file
	err = ioutil.WriteFile("result.txt", []byte(text), 0644)
	if err != nil {
		fmt.Println("Error writing output file:", err)
		return
	}
}

func compare(a, b string) int {
	return strings.Compare(a, b)
}

func toUpper(str string) string {
	return strings.ToUpper(str)
}

func toLower(str string) string {
	return strings.ToLower(str)
}

func capitalise(str string) string {
	return strings.Title(str)
}

func replaceHexWithDec(text string) string {
	words := strings.Fields(text)
	prevWord := ""
	for i, word := range words {
		if word == "(hex)" && i > 0 {
			// Replace previous word with its decimal equivalent
			prevHex := strings.TrimPrefix(prevWord, "0x")
			decNum, err := strconv.ParseInt(prevHex, 16, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(decNum, 10)
			}
		} else {
			// Store current word as previous word for next iteration
			prevWord = word
		}
	}
	// Remove any remaining instances of "(hex)"
	replacedText := strings.Join(words, " ")
	replacedText = strings.ReplaceAll(replacedText, "(hex)", "")
	return replacedText
}

func replaceBinWithDec(text string) string {
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

func replaceLowWithLowercase(text string) string {
	words := strings.Fields(text)
	for i, word := range words {
		if strings.Contains(word, "(low,") && i > 0 {
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

func replaceUpWithUppercase(text string) string {
	words := strings.Fields(text)
	for i, word := range words {
		if strings.Contains(word, "(up,") && i > 0 {
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

func replaceCapWithCapitalized(text string) string {
	words := strings.Fields(text)
	for i, word := range words {
		if strings.Contains(word, "(cap,") && i > 0 {
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

func formatPunctuation(text string) string {
	// Define regular expressions for matching specific punctuation marks and groups
	re1 := regexp.MustCompile(`(\w)[ ]*([.,!?:;])([.,!?:;])+([ ]|\n|\r|\t)*(\w)`) // Match punctuation followed by multiple punctuation marks
	re2 := regexp.MustCompile(`(\w)[ ]*([.,!?:;])([ ]|\n|\r|\t)*(\w)`)            // Match punctuation followed by a space and a word

	// Replace matching instances with modified versions
	text = re1.ReplaceAllStringFunc(text, func(str string) string {
		return strings.ReplaceAll(str, " ", "") // Remove all spaces between punctuation marks and the surrounding words
	})

	text = re2.ReplaceAllStringFunc(text, func(str string) string {
		return strings.Trim(str, " ") + " " // Remove any extra spaces before the punctuation mark and add one space after it
	})

	return text
}

func replaceQuotes(text string) string {
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

func formatText(text string) string {
	// Convert a to an if the next word begins with a vowel or h
	text = strings.ReplaceAll(text, " a ", " an ")
	text = strings.ReplaceAll(text, " A ", " An ")
	words := strings.Fields(text)
	for i := 0; i < len(words)-1; i++ {
		if strings.ToLower(words[i]) == "an" {
			if strings.HasPrefix(strings.ToLower(words[i+1]), "a") ||
				strings.HasPrefix(strings.ToLower(words[i+1]), "e") ||
				strings.HasPrefix(strings.ToLower(words[i+1]), "i") ||
				strings.HasPrefix(strings.ToLower(words[i+1]), "o") ||
				strings.HasPrefix(strings.ToLower(words[i+1]), "u") ||
				strings.HasPrefix(strings.ToLower(words[i+1]), "h") {
				words[i] = "a"
			}
		}
	}
	text = strings.Join(words, " ")

	return text
}

func formatPunctuations(text string) {
	panic("unimplemented")
}
