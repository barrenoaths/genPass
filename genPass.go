package main

import (
	"crypto/rand"
	"encoding/binary"
	"flag"
	"fmt"
)

func main() {
	numbers := flag.Bool("n", false, "Include numbers in the password")
	capitals := flag.Bool("c", false, "Include capital letters in the password")
	symbols := flag.Bool("s", false, "Include symbols in the password")
	length := flag.Int64("l", 12, "Length of the password")

	flag.Parse()

	inclCharacters := "abcdefghijklmnopqrstuvwxyz"
	inclNumbers := "0123456789"
	inclUpperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	inclSymbols := "!@#$%^&*()"
	generatedPassword := ""

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

	for range *length {
		randomIndex, _ := genRandomIndex(len(inclCharacters))
		generatedPassword += string(inclCharacters[randomIndex])
	}

	fmt.Println(generatedPassword)

}

func genRandomIndex(n int) (int, error) {
	var randomBytes [4]byte
	_, err := rand.Read(randomBytes[:])
	if err != nil {
		return 0, err
	}

	randomValue := int(binary.LittleEndian.Uint32(randomBytes[:]))
	return randomValue % n, nil
}
