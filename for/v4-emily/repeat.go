package iteration

import "strings"

// Repeat takes a character and the number of times the character is repeated to return a string of repeated characters.
func Repeat(character string, repeatCount int) string {
	//var repeated string
	//for i := 0; i < repeatCount; i++ {
	//	repeated += character
	//}
	//return repeated
	return strings.Repeat(character, repeatCount)
}
