package piscine

func AtoiBase(s string, base string) int {
	// Step 1: Validate base
	if len(base) < 2 {
		return 0
	}

	// Check for duplicates or invalid characters
	seen := make(map[rune]bool)
	for _, ch := range base {
		if ch == '+' || ch == '-' || seen[ch] {
			return 0
		}
		seen[ch] = true
	}

	// Step 2: Convert string to integer
	baseLen := len(base)
	result := 0

	for _, ch := range s {
		index := indexInBase(ch, base)
		if index == -1 {
			return 0 // character not in base
		}
		result = result*baseLen + index
	}

	return result
}

// Helper function to find index of a rune in the base string
func indexInBase(ch rune, base string) int {
	for i, b := range base {
		if b == ch {
			return i
		}
	}
	return -1
}
