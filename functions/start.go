package asciiart

import (
	"fmt"
	"strings"
)

// This function is the starting point of the program.
// Banner file and input string has already been validated in main.go.
func Start(input, fileFlag string) string {
	asciiMap := ReadBanner(fileFlag)
	artString := ""
	words := strings.Split(input, "\n")
	fmt.Println(words)

	for _, word := range words {
		if word == "" {
			fmt.Println()
		} else {
			art := PrintArt(asciiMap, word)
			artString = art
		}
	}
	// This return value is used for testing purposes.
	return artString
}
