package asciiart

// Func GetArt() takes the map of ascii art and and string then returns the art in string form.
func GetArt(asciiMap map[rune][]string, input string) string {
	store := [][]string{}
	for _, chr := range input {
		store = append(store, asciiMap[chr])
	}
	// Hold the word in a variable.
	word := ""
	for i := 0; i < 8; i++ {
		for j := 0; j < len(store); j++ {
			word += store[j][i]
		}
		if i < 7 {
			word += "\n"
		}
	}
	return word
}

// Check if strings only contains '\n's.
func NewLinesOnly(input string) bool {
	for _, v := range input {
		if v != '\n' {
			return false
		} else {
			continue
		}
	}
	return true
}
