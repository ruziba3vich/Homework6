package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Ketma-ket takrorlanuvchi belgilarni takrorlash sonidan keyin belgi bilan almashtirib,
	// siqib chiqaradigan dastur yarating.
	// Masalan, "aaabbbccc" "a3b3c3" ga aylanishi kerak.

	var s string
	fmt.Print("Enter the string : ")
	fmt.Scan(& s)
	var char byte = s[0]
	var result strings.Builder
	count := 1
	for i := 1; i < len(s); i ++ {
		if s[i] == char {
			count ++
		} else {
			result.WriteByte(char)
			result.WriteString(strconv.Itoa(count))
			count = 1
			char = s[i]
		}
	}

	result.WriteByte(char)
	result.WriteString(strconv.Itoa(count))

	fmt.Println("The result is :", result.String())
}
