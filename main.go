package main

import (
	"fmt"
	"os"
	"path/filepath"

	asciiart "asciiart/functions"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Error: String not found. Usage: go run . \"<string>\" <flag_name>")
		return
	}
	input := os.Args[1]

	validated := asciiart.StrValidate(input)

	if !validated {
		return
	}

	input = asciiart.FormatSpChar(input)
	if input == "" {
		return
	}
	// if string contains New lines only:
	newLinesOnly := asciiart.NewLinesOnly(input)

	if newLinesOnly {
		for i := 0; i < len(input); i++ {
			fmt.Println()
		}
		return
	}

	bannerTemplate := filepath.Join("banners", "standard.txt")
	arguments := os.Args[1:]

	// Code checks if the length of the argument is exactly two and the second one is not a flag.
	if len(arguments) == 2 {
		flag := os.Args[2]
		if flag != "-sh" && flag != "-th" && flag != "-st" {
			fmt.Printf("Error: \"%v\" is not a flag_name. Usage: go run . \"<string>\" <flag_name> (flags: -sh, -st, -th)\n", flag)
			return
		}

		// Handling Flags
		switch flag {
		case "-sh":
			bannerTemplate = filepath.Join("banners", "shadow.txt")
		case "-th":
			bannerTemplate = filepath.Join("banners", "thinkertoy.txt")
		default:
			bannerTemplate = filepath.Join("banners", "standard.txt")
		}
	}
	if len(arguments) > 2 {
		fmt.Println("Error: Too many arguments. Usage: go run . \"<string>\" <flag_name>")
		return
	}

	asciiart.ValidateBanner(bannerTemplate)
	artStringstr := asciiart.WordDistributer(input, bannerTemplate)
	fmt.Println(artStringstr)
}
