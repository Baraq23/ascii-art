package main

import (
	"fmt"
	"os"


	asciiart "asciiart/files"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Error: Input string not typed")
		return
	}
	if len(os.Args) > 3 {
		fmt.Println("Error: too many arguments")
		fmt.Println("Write 'st', 'sh' or 'th' flags for standard, shadow or thinkertoy banner files respectively")
		return
	}

	flagFile := ""
	if len(os.Args) == 2 {
		flagFile = "standard.txt"
	}
	if len(os.Args) == 3 {

		switch os.Args[2] {
		case "sh":
			flagFile = "shadow.txt"
		case "th":
			flagFile = "thinkertoy.txt"
		case "st":
			flagFile = "standard.txt"
		default:
			fmt.Println("Write 'st', 'sh' or 'th' flags for standard, shadow or thinkertoy banner files respectively")
			fmt.Printf("%v is not dfined as a flag standard(st) file has been selected\n", os.Args[2])
			flagFile = "standard.txt"
		}
	}
	
	
	// flagFile := "standard.txt"
		
	input := os.Args[1]
	validated := asciiart.StrValidate(input)

	// if len(os.Args) == 1 {
	// 	fmt.Println("Error: Input string not typed")
	// 	return
	// }
	// input1 := os.Args[1:]
	// inputStr := strings.Join(input1, " ")
	// validated := asciiart.StrValidate(inputStr)

	if !validated {
		return
	}
	

	input = asciiart.FormatSpChar(input)
	if input == "" {
		return
	}
	// print newlines
	nlsOnly := nlsOnly(input)

	if nlsOnly {
		for i := 0; i < len(input); i++ {
			fmt.Println()
		}
		return
	}
	str := asciiart.Start(input, flagFile)
	fmt.Sprintln(str)
}

//check if trings only contains '\n's
func nlsOnly(input string) bool {
	for _, v := range input {
		if v != '\n' {
			return false
		} else {
			continue
		}
	}
	return true
}
