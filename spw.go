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
	pw[0] = Upper[int(indicies[0])%len(Upper)]
	pw[1] = Lower[int(indicies[1])%len(Lower)]
	pw[2] = Digit[int(indicies[2])%len(Digit)]
	pw[3] = Special[int(indicies[3])%len(Special)]
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
