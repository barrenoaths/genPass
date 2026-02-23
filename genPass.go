package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("generate da passwords")
	numbers := flag.Bool("n", false, "Include numbers in the password")
	capitals := flag.Bool("c", false, "Include capital letters in the password")
	symbols := flag.Bool("s", false, "Include symbols in the password")
	length := flag.Int("l", 12, "Length of the password")

	flag.Parse()

	inclCharacters := "abcdefghijklmnopqrstuvwxyz"
	inclNumbers := "0123456789"
	inclUpperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	inclSymbols := "!@#$%^&*()"
	//generatedPassword := ""

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

	for i := range *length {
		randomInt, err := rand.Int(rand.Reader, big.NewInt(int64(len(inclCharacters))))
		if err != nil {
			fmt.Println("Error while generating random characters")
			return
		}
		fmt.Println("Counter:", i)
		fmt.Println(randomInt)
	}
	fmt.Println(inclCharacters)
	//bs above to be changed, just learnin
}
