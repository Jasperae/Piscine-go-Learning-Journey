package piscine

// This function that returns true if the string passed as parameter contains only uppercase characters, otherwise, the reslt to be returned is false.

func IsUpper(s string) bool {
	r := []rune(s)
	for i := range r {
		if r[i] < 65 || r[i] > 90 {
			return false
		}
	}
	return true
}
