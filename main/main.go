package main

import (
	"fmt"
	"goreloaded"
	"os"
)

func main() {
	// Read input file
	input, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}
	text := string(input)

	// Apply transformations
	text = goreloaded.ReplaceHexWithDec(text)
	text = goreloaded.ReplaceBinWithDec(text)
	text = goreloaded.ReplaceLowWithLowercase(text)
	text = goreloaded.ReplaceUpWithUppercase(text)
	text = goreloaded.ReplaceCapWithCapitalized(text)
	text = goreloaded.FormatText(text)
	text = goreloaded.ReplaceQuotes(text)

	// Write output file
	err = os.WriteFile("result.txt", []byte(text), 0644)
	if err != nil {
		fmt.Println("Error writing output file:", err)
		return
	}
}
