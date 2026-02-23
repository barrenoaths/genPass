package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("generate da passwords")
	numbers := flag.Bool("n", false, "Include numbers in the password")
	capitals := flag.Bool("c", false, "Include capital letters in the password")
	symbols := flag.Bool("s", false, "Include symbols in the password")

	flag.Parse()

	inclCharacters := "abcdefghijklmnopqrstuvwxyz"
	inclNumbers := "0123456789"
	inclUpperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	inclSymbols := "!@#$%^&*()"

	appliedFlags := map[*bool]string{
		numbers:  inclNumbers,
		capitals: inclUpperCase,
		symbols:  inclSymbols,
	}

	for flag, charSet := range appliedFlags {
		if *flag {
			inclCharacters += charSet
		}
	}

	fmt.Println(inclCharacters)

}
