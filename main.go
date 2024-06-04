package main

import (
	"fmt"
	"os"

	asciiart "asciiart/functions"
)

func main() {
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

	bannerTemplate := "standard.txt"
	arguments := os.Args[1:]

	// code checks if the length of the argument is exactly two and the second one is not a flag.
	if len(arguments) == 2 {
		flag := os.Args[2]
		if flag != "-sh" && flag != "-th" && flag != "-st" {
			fmt.Printf("Error: \"%v\" is not a flag_name. Usage: go run . \"<string>\" <flag_name>\n", flag)
			return
		}

		// Handling Flags
		switch flag {
		case "-st":
			bannerTemplate = "standard.txt"
		case "-sh":
			bannerTemplate = "shadow.txt"
		case "-th":
			bannerTemplate = "thinkertoy.txt"
		default:
			bannerTemplate = "standard.txt"
		}
	}
	if len(arguments) > 2 {
		fmt.Println("Error: Too many arguments. Usage: go run . \"<string>\" <flag_name>")
		return
	}

	asciiart.ValidateBanner(bannerTemplate)
	str := asciiart.Start(input, bannerTemplate)
	fmt.Sprintln(str)
}
