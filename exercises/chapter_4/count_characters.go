package main

import (
	"fmt"
	"unicode/utf8"
)

var str1 string = "asSASA ddd dsjkdsjs dk"
var str2 string = "asSASA ddd dsjkdsjsこん dk"

func mainCountCharacters() {
	// count number of characters:
	fmt.Printf("The number of bytes in string str1 is %d\n", len(str1))
	fmt.Printf("The number of characters in string str1 is %d\n", utf8.RuneCountInString(str1))

	fmt.Printf("The number of bytes in string str2 is %d\n", len(str2))
	fmt.Printf("The number of characters in string str2 is %d\n", utf8.RuneCountInString(str2))
}

/* Output:
The number of bytes in string str1 is 22
The number of characters in string str1 is 22
The number of bytes in string str2 is 28
The number of characters in string str2 is 24
*/
