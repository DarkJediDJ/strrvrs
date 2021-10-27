// Package reverse implements some functions to edit strings.
package reverse

import (
	"fmt"
	"regexp"
)

//Reverse provides way to reverse string
func Reverse(s string) string {
	matched, err := regexp.MatchString(`[\x{1F600}-\x{1F64F}]`, s)
	if err != nil {
		fmt.Printf(err.Error())
	}
	if matched {
		return s
	}
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}
