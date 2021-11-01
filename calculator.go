package strrvrs

//Calculator provides way to count number of characters in the string
func Calculator(s string) int {
	return len([]rune(s))
}
