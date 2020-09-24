package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// declare a string with both Chinese and English characters
	s := "世界 means world"
	// UTFMax is 4 - up to 4 bytes per encoded rune
	var buf [utf8.UTFMax]byte

	// iterate over the string
	for i, r := range s {
		runeLength := utf8.RuneLen(r)
		// calculate the slice offset for the bytes associated with this rune
		sliceOffset := i + runeLength

		// copy of rune from the string to our buffer
		copy(buf[:], s[i:sliceOffset])
		// display the details
		fmt.Printf("%2d: %q; codepint: %#6x; encoded bytes: %#v\n", i, r, r, buf[:runeLength])
	}
}
