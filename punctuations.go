package goreloaded

import "strings"

func FormatText(text string) string {
	// Split the text into words
	words := strings.Fields(text)

	// Create a buffer to store the formatted text
	var buffer strings.Builder

	// Iterate over the words
	for i, word := range words {
		// Check if the current word ends with a punctuation mark
		if strings.ContainsAny(word, ".,!?;:") {
			// If it does, remove any trailing whitespace and add it to the buffer
			word = strings.TrimRight(word, " \t\n\r\f\v")
			buffer.WriteString(word)

			// Check if the next word is also a punctuation mark
			if i+1 < len(words) && strings.ContainsAny(words[i+1], ".,!?;:") {
				// If it is, add it to the buffer without any space
				buffer.WriteString(words[i+1])
			} else {
				// Otherwise, add a space after the punctuation mark
				buffer.WriteString(" ")
			}
		} else {
			// If the current word does not end with a punctuation mark, add it to the buffer with a space
			buffer.WriteString(word)
			buffer.WriteString(" ")
		}
	}

	// Replace groups of punctuation marks with the formatted version
	output := strings.ReplaceAll(buffer.String(), " .", ".")
	output = strings.ReplaceAll(output, " ,", ", ")
	output = strings.ReplaceAll(output, " !", "! ")
	output = strings.ReplaceAll(output, " ?", "? ")
	output = strings.ReplaceAll(output, " :", ": ")
	output = strings.ReplaceAll(output, " ;", "; ")
	output = strings.ReplaceAll(output, " ...", "...")
	output = strings.ReplaceAll(output, "!?", "!? ")
	output = strings.ReplaceAll(output, "?!", "?! ")
	output = strings.ReplaceAll(output, "...", "...")
	output = strings.ReplaceAll(output, ",,", ", ")

	return output
}
func FormatPunctuations(text string) {
	panic("unimplemented")
}