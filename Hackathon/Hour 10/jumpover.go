package piscine

func JumpOver(str string) string {
	result := ""
	index := 0

	for _, char := range str {
		if index%3 == 2 {
			result += string(char)
		}
		index++
	}

	return result + "\n"
}
