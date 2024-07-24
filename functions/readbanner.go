package asciiart

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

// This function ReadBanner() maps the ascii charracters to their corresponding art.
func ReadBanner(file string) map[rune][]string {
	bannerOk := checkBanner(file)
	if !bannerOk {
		fmt.Printf("%v file is corrupt\n", file)
		os.Exit(1)
	}

	asciiMap := make(map[rune][]string)
	bannerFile, err := os.Open(file)
	if err != nil {
		fmt.Printf("Error: Could not open %v file\n", file)
		os.Exit(1)
	}
	defer bannerFile.Close()

	scanner := bufio.NewScanner(bannerFile)

	for i := 32; i <= 126; i++ {
		runeValue := i
		scanner.Scan()
		var art []string
		for i := 0; i < 8; i++ {
			scanner.Scan()
			line := scanner.Text()
			art = append(art, line)
		}
		asciiMap[rune(runeValue)] = art
	}
	return asciiMap
}

// Function checkBanner() uses the crypto package to check the validity of the banner
func checkBanner(f string) bool {
	var hashMap map[string]string = map[string]string{
		"banners/standard.txt":   "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf",
		"banners/shadow.txt":     "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73",
		"banners/thinkertoy.txt": "092d0cde973bfbb02522f18e00e8612e269f53bac358bb06f060a44abd0dbc52",
	}

	fileBytes, err := os.ReadFile(f)
	if err != nil {
		fmt.Println("Error Checking banner")
		os.Exit(1)
	}

	hasher := sha256.New()
	hasher.Write(fileBytes)
	hashInBytes := hasher.Sum(nil)
	fHash := hex.EncodeToString(hashInBytes)

	return fHash == hashMap[f]
}
