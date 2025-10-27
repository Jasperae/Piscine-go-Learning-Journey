package piscine

// I created a function named NRune which is a string that would return a rune
func NRune(s string, n int) rune {
	if len(s) < n || n < 1 {
		return 0
	}
	runes := []rune(s)
	return runes[n-1]
}
