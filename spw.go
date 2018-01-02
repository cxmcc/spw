package main

import (
	"crypto/rand"
	"fmt"
	"github.com/atotto/clipboard"
)

var upper = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var lower = []byte("abcdefghijklmnopqrstuvwxyz")
var digit = []byte("0123456789")
var special = []byte("`~!@#$%^&*()-=_+[]\\{};':\",./<>?")
var chars = [][]byte{
	upper, lower, digit, special,
}
var allChars []byte

const length = 16
const numSwaps = 1024

func init() {
	for _, r := range chars {
		allChars = append(allChars, r...)
	}
}

func generateBytes(a []byte) {
	_, err := rand.Read(a)
	if err != nil {
		panic("Error generating random bytes.")
	}
}

func generatePw() string {
	pw := make([]byte, length)
	indicies := make([]byte, length)
	swap := make([]byte, numSwaps*2)
	generateBytes(indicies)
	generateBytes(swap)
	// Make sure the password contain at least one char for each class of characters
	for i := 0; i < 4; i++ {
		dict := chars[i]
		pw[i] = dict[int(indicies[i])%len(dict)]
	}
	for i := 4; i < length; i++ {
		pw[i] = allChars[int(indicies[i])%len(allChars)]
	}
	for c := 0; c < numSwaps; c++ {
		i, j := swap[c*2]%length, swap[c*2+1]%length
		pw[i], pw[j] = pw[j], pw[i]
	}
	return string(pw)
}

func main() {
	pw := generatePw()
	fmt.Println(pw)
	clipboard.WriteAll(pw)
}
