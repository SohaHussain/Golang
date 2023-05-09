/*
	Write a program which prompts the user to enter a string. The program searches through the entered string for the

characters ‘i’, ‘a’, and ‘n’. The program should print “Found!” if the entered string starts with the character ‘i’,
ends with the character ‘n’, and contains the character ‘a’. The program should print “Not Found!” otherwise.
The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	var pre, suf = "i", "n"
	var inn = "a"
	fmt.Printf("enter a string")
	fmt.Scan(&s)
	s = strings.ToLower(s)
	cond1 := strings.HasPrefix(s, pre)
	cond2 := strings.HasSuffix(s, suf)
	cond3 := strings.Contains(s, inn)

	if cond1 && cond2 && cond3 {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
