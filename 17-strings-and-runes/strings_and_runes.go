package main

/*
A Go string is a read-only slice of bytes. The language and the standard library
treat strings specially - as containers of text encoded in UTF-8. In other languages,
strings are made of “characters”. In Go, the concept of a character is called a rune
- it’s an integer that represents a Unicode code point.
*/

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// s is a string assigned a literal value representing the word “hello”
	// in the Portuguese language. Go string literals are UTF-8 encoded text.
	const s = "olá"

	// Since strings are equivalent to []byte, this will produce the length
	// of the raw bytes stored within.
	fmt.Println("Len:", len(s))

	// Indexing into a string produces the raw byte values at each index.
	// This loop generates the hex values of all the bytes that constitute
	// the code points in s.
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// To count how many runes are in a string, we can use the utf8 package.
	// Note that the run-time of RuneCountInString dependes on the size of the string,
	// because it has to decode each UTF-8 rune sequentially. Some Thai characters are
	// represented by multiple UTF-8 code points, so the result of this count
	// may be surprising.
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// A range loop handles strings specially and decodes each rune
	// along with its offset in the string.
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// We can achieve the same iteration by using the utf8.DecodeRuneInString
	// function explicitly.
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		// This demonstrates passing a rune value to a function.
		examineRune(runeValue)
	}
}

func examineRune(r rune) {

	// Values enclosed in single quotes are rune literals.
	// We can compare a rune value to a rune literal directly.
	if r == 'o' {
		fmt.Println("found o")
	} else if r == 'l' {
		fmt.Println("found so l")
	}
}
