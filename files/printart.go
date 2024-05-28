package asciiart

import "fmt"

// This function stores the ascii art to then prints the ascii art
func PrintArt(asciiMap map[rune][]string, input string) string {
	// slice of a slice to store the Art
	store := [][]string{}	
	for _, chr := range input {
		store = append(store, asciiMap[chr])
	}
	// Hold the word in a variable
	word := ""
	for i := 0; i < 8; i++ {
		for j := 0; j < len(store); j++ {
			word += store[j][i]
		}
		if i != 7 {
			word += "\n"
		}
	}
	// write the Art
	fmt.Println(word)
	//This return value will be usefeul during testing
	return word

}
