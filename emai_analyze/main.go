package main

import "fmt"

func main() {
	// Elektron pochta manzillarini tahlil qiling:
	// emailList o'zgaruvchidan domen nomlarini
	// (masalan, "example.com") chiqaradigan dasturni ishlab chiqing.
	// Keyin, har bir domen nomining takrorlanishini hisoblang va natijalarni chop eting.

	emailList := []string{
		"john@example.com",
		"alice@example.com",
		"bob@example.com",
		"sam@example.net",
		"lisa@example.org",
		"test@example.com",
		"example@example.com",
		"user1@gmail.com",
		"user2@gmail.com",
		"user3@gmail.com",
		"user4@yahoo.com",
		"user5@yahoo.com",
		"user6@outlook.com",
		"user7@outlook.com",
		"admin@example.com",
		"info@example.com",
		"contact@example.com",
		"support@example.com",
		"sales@example.com",
		"feedback@example.com",
		"webmaster@example.com",
	}

	m := map[string] int {}

	for _, str := range emailList {
		nstr := convert(str)
		m[nstr] = getOrDefault(m, nstr) + 1
	}

	fmt.Println("The number of domeins in the given list are :")

	for domein, occurance := range m {
		fmt.Println(domein, ":", occurance)
	}

}

func convert(s string) string {
	var i int

	for s[i] != '@' {
		i ++
	}
	return s[i + 1:]
}

func getOrDefault(m map[string] int, s string) int {
	val, e := m[s]

	if e {
		return val
	}
	return 0
}
