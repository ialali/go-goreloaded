package main

import (
	"fmt"
	"os"
	"strings"
)

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

func main() {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(data))
	input := string(data)
	//fmt.Println(input)
	result := strings.Fields(input)
	fmt.Println(result)

	for i, v := range result {

		if v == "(up)" {
			result[i-1] = toUpper(result[i-1])
		}

		if compare(v, "(cap)") == 0 {
			result[i-1] = capitalise(result[i-1])

		}

		if compare(v, "(lower)") == 0 {
			result[i-1] = toLower(result[i-1])
		}
q
	}

	// fmt.Println(toUpper("gopher"))
	// fmt.Println(toLower("01FOUNDERS"))
	// fmt.Println(toTitle("hey how are you?"))
	// input := "1E (hex) files were added"
	// output := replaceHex(input)
	// fmt.Println(output) // Output: 30 files were added
}
