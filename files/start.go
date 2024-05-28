package asciiart

import (
	"fmt"
	"strings"
)

func Start(input, fileFlag string) string {
	
	asciiMap := ReadBanner(fileFlag)
	slice := ""
	words := strings.Split(input, "\n")
	
	for _, word := range words {
		if word == "" {
			fmt.Println()
		} else {
			art := PrintArt(asciiMap, word)
			slice = art
		}
	}
	// This return value is used for testin purposes
	return slice
}
