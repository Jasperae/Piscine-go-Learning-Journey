package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	counts := make(map[string]int)
	word := ""

	// We iterate through the string and the spaces act as the trigger
	for _, charRune := range str {
		char := string(charRune)

		if char != " " {
			// 1. Accumulate the character
			word += char
		} else {
			// 2. Tally the accumulated word (or empty string)
			// This will tally "" if the previous char was also a space (or if we started with one)
			counts[word]++

			// 3. Reset the word buffer to an empty string
			word = ""
		}
	}

	// 4. Final check for the last word (since the loop ended without a space)
	counts[word]++

	return counts
}
