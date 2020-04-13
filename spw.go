package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/nbutton23/zxcvbn-go"
	"os"
)

var withoutSpecial bool
var doPrint bool
var length int
var upper = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var lower = []byte("abcdefghijklmnopqrstuvwxyz")
var digit = []byte("0123456789")
var special = []byte("`~!@#$%^&*()-=_+[]\\{};':\",./<>?")
var charsWithSpecial = [][]byte{
	upper, lower, digit, special,
}
var charsWithoutSpecial = [][]byte{
	upper, lower, digit,
}
var allCharsWithSpecial []byte
var allCharsWithoutSpecial []byte

const numSwaps = 1024

func init() {
	flag.BoolVar(&withoutSpecial, "w", false, "without special characters")
	flag.BoolVar(&doPrint, "p", false, "print the password, in addition to copying to clipboard)")
	flag.IntVar(&length, "n", 32, "length of password")

	flag.Parse()
	if length < 12 {
		fmt.Fprintln(os.Stderr, "password length too short.")
		os.Exit(1)
	}
	for _, r := range charsWithSpecial {
		allCharsWithSpecial = append(allCharsWithSpecial, r...)
	}
	for _, r := range charsWithoutSpecial {
		allCharsWithoutSpecial = append(allCharsWithoutSpecial, r...)
	}
}

func generateBytes(a []byte) {
	_, err := rand.Read(a)
	if err != nil {
		panic("Error generating random bytes.")
	}
}

func generatePw() string {
	var chars [][]byte
	var allChars []byte

	if withoutSpecial {
		chars = charsWithoutSpecial
		allChars = allCharsWithoutSpecial
	} else {
		chars = charsWithSpecial
		allChars = allCharsWithSpecial
	}

	pw := make([]byte, length)
	indicies := make([]byte, length)
	swap := make([]byte, numSwaps*2)

	generateBytes(indicies)
	generateBytes(swap)
	// Make sure the password contain at least one char for each class of characters
	for i := 0; i < len(chars); i++ {
		dict := chars[i]
		pw[i] = dict[int(indicies[i])%len(dict)]
	}
	for i := len(chars); i < length; i++ {
		pw[i] = allChars[int(indicies[i])%len(allChars)]
	}
	for c := 0; c < numSwaps; c++ {
		i, j := int(swap[c*2])%length, int(swap[c*2+1])%length
		pw[i], pw[j] = pw[j], pw[i]
	}
	return string(pw)
}

func main() {
	var pw string
	for {
		pw = generatePw()
		match := zxcvbn.PasswordStrength(pw, nil)
		if match.Score == 4 {
			break
		}
	}

	err := clipboard.WriteAll(pw)
	if err != nil {
		panic("Error when writing to clipboard.")
	}

	if doPrint {
		fmt.Println(pw)
	} else {
		fmt.Fprintln(os.Stderr, "NOTE: Generated password sent to clipboard.")
	}
}
