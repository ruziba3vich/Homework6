package main

import (
	"fmt"
	"strings"
)

var str = `
	Alice was beginning to get very tired of sitting by her sister on the bank, 
	and of having nothing to do: once or twice she had peeped into the book her sister was reading, 
	but it had no pictures or conversations in it, 
	'and what is the use of a book,' thought Alice 'without pictures or conversation?'
	So she was considering in her own mind (as well as she could, 
	for the hot day made her feel very sleepy and stupid), 
	whether the pleasure of making a daisy-chain would be worth the trouble of getting up 
	and picking the daisies, when suddenly a White Rabbit with pink eyes ran close by her.
`

func main() {
	// Unli va undoshlarni sanash:
	// Str o'zgaruvchi ni ichida unli va undoshlar sonini hisoblaydigan dastur tuzing.
	// Bo'shliqlarga e'tibor bermang va katta va kichik harflarni ekvivalent deb hisoblang.

	vowels := map[byte] byte {
		'a': 'a',
		'o': 'o',
		'e': 'e',
		'i': 'i',
		'u': 'u',
	}

	var vowelscnt int
	var consonantscnt int

	for i := range str {
		if char := str[i]; char != ' ' {
			if _, exists := vowels[char]; exists {
				vowelscnt ++
			} else {
				consonantscnt ++
			}
		}
	}

	fmt.Printf("Number of vowels in the given string is %d, and number of consonants are %d\n", vowelscnt, consonantscnt)
	fmt.Println(len(str))
}

func ToLower(letter byte) string {
	lowercaseLetter := strings.ToLower(string(letter))

	return lowercaseLetter
}
