package main

import (
	"crypto/rand"
	"fmt"
)

var Upper = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var Lower = []byte("abcdefghijklmnopqrstuvwxyz")
var Digit = []byte("0123456789")
var Special = []byte("!@#$%^&*()-_=+,.?/:;{}[]`~\"")
var Chars = [][]byte{
	Upper, Lower, Digit, Special,
}
var All []byte

const Length = 16

func initConstants() {
	for _, r := range Chars {
		All = append(All, r...)
	}
}

func main() {
	initConstants()
	pw := make([]byte, Length)
	indicies := make([]byte, Length)
	_, err := rand.Read(indicies)
	if err != nil {
		panic("Wrong.")
	}
	for i := 0; i < 4; i++ {
		dict := Chars[i]
		pw[i] = dict[int(indicies[i])%len(dict)]
	}
	for i := 4; i < Length; i++ {
		pw[i] = All[int(indicies[i])%len(All)]
	}
	count := 0
	swap := make([]byte, 2)
	for count < 1000 {
		count += 1
		_, err := rand.Read(swap)
		if err != nil {
			panic("Wrong.")
		}
		i, j := swap[0]%Length, swap[1]%Length
		pw[i], pw[j] = pw[j], pw[i]
	}
	fmt.Println(string(pw))
}
