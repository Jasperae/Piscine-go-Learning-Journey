package piscine

// This function returns true if the string passed as parameter contains only lowercase characters, otherwise, returns false.

func IsLower(s string) bool {
	r := []rune(s)
	for i := range r {
		if r[i] < 97 || r[i] > 122 {
			return false
		}
	}
	return true
}
