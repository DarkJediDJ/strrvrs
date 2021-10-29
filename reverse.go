// Package strrvrs implements some functions to edit strings.
package strrvrs

//Reverse provides way to reverse string
func Reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

//Calculator provides way to count number of characters in the string
func Calculator(s string) int {
	return len([]rune(s))
}
