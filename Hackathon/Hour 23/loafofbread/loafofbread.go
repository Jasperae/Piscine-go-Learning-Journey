package piscine

func LoafOfBread(str string) string {
	// Count non-space characters
	count := 0
	for _, r := range str {
		if r != ' ' {
			count++
		}
	}
	if count < 5 {
		return "Invalid Output\n"
	}

	result := ""
	wordLen := 0
	skip := false

	for _, r := range str {
		if r == ' ' {
			continue
		}
		if skip {
			skip = false
			continue
		}
		result += string(r)
		wordLen++
		if wordLen == 5 {
			result += " "
			wordLen = 0
			skip = true
		}
	}

	// Trim trailing space if present
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}
	return result + "\n"
}
