package asciiart

import (
	"fmt"
	"strings"
)

// Func WordDistributer splits the string int a slice then sends eaach part to be printed.
func WordDistributer(input, bannerTemplate string) string {
	asciiMap := ReadBanner(bannerTemplate)
	artString := strings.Builder{}
	words := strings.Split(input, "\n")
	fmt.Printf("%q\n", words)
	for _, word := range words {
		if word == "" {
			artString.WriteString("\n")
		} else {

			art := GetArt(asciiMap, word)
			artString.WriteString(art)

		}
		// artString+="\n"
	}
	return artString.String()
}
