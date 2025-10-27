package piscine

func LoafOfBread(str string) string {
	// First, count non-space characters
	count := 0
	for _, r := range str {
		if r != ' ' {
			count++
		}
	}

	// If fewer than 5 valid characters, return error
	if count < 5 {
		return "Invalid Output\n"
	}

	// Build result manually
	result := ""
	wordLen := 0
	skipNext := false

	for _, r := range str {
		if r == ' ' {
			continue
		}
		if skipNext {
			skipNext = false
			continue
		}
		result += string(r)
		wordLen++
		if wordLen == 5 {
			result += " "
			wordLen = 0
			skipNext = true
		}
	}

	// Trim trailing space if needed and add newline
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}
	return result + "\n"
}
