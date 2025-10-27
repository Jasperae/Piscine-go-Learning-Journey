package piscine

// This function returns true if the string passed as a parameter contains only printable characters, otherwise, returns false.
func IsPrintable(s string) bool {
	if len(s) == 0 {
		return false
	}
	// ASCII'den alinan degeler
	for i := 0; i < len(s); i++ {
		if s[i] < 32 || s[i] > 126 {
			return false
		}
	}
	return true
}
