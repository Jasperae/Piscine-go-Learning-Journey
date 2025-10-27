package piscine

// This unction that returns true if the string passed as the parameter only contains alphanumerical characters or is empty, otherwise, and returns false.
func IsAlpha(s string) bool {
	if len(s) == 0 {
		return false
	}
	for i := 0; i < len(s); i++ {
		if !(s[i] >= 'A' && s[i] <= 'Z') &&
			!(s[i] >= 'a' && s[i] <= 'z') && !(s[i] >= '0' && s[i] <= '9') {
			return false
		}
	}
	return true
}
