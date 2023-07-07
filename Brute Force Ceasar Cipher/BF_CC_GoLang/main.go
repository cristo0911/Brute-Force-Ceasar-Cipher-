package main

import (
	"fmt"
)

func bruteForce(message string) {
	for key := 1; key <= 26; key++ {
		decryption := decryptMessage(message, key)
		fmt.Printf("Decryption with key %d: %s\n", key, decryption)
	}
}

func decryptMessage(message string, key int) string {
	decryption := ""
	for _, char := range message {
		if char >= 'a' && char <= 'z' {
			decryption += string(shiftChar(char, key, 'a', 26))
		} else if char >= 'A' && char <= 'Z' {
			decryption += string(shiftChar(char, key, 'A', 26))
		} else {
			decryption += string(char)
		}
	}
	return decryption
}

func shiftChar(char rune, key int, baseChar rune, alphabetSize int) rune {
	return ((char - baseChar + rune(key)) % rune(alphabetSize)) + baseChar
}

func main() {
	targetMessage := "Ugew gnwj zwjw Oslkgf"
	bruteForce(targetMessage)
}
