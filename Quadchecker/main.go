package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// getDimensions calculates width and height from input string
func getDimensions(input string) (int, int) {
	if input == "" {
		return 0, 0
	}

	lines := strings.Split(input, "\n")
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	height := len(lines)
	width := 0
	if height > 0 {
		width = len(lines[0])
	}

	return width, height
}

// QuadA pattern
func generateQuadA(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	result := ""
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if (row == 1 || row == y) && (col == 1 || col == x) {
				result += "o"
			} else if row == 1 || row == y {
				result += "-"
			} else if col == 1 || col == x {
				result += "|"
			} else {
				result += " "
			}
		}
		result += "\n"
	}
	return result
}

// QuadB pattern
func generateQuadB(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	result := ""
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if row == 1 && col == 1 {
				result += "/"
			} else if row == 1 && col == x {
				result += "\\"
			} else if row == y && col == 1 {
				result += "\\"
			} else if row == y && col == x {
				result += "/"
			} else if row == 1 || row == y || col == 1 || col == x {
				result += "*"
			} else {
				result += " "
			}
		}
		result += "\n"
	}
	return result
}

// QuadC pattern
func generateQuadC(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	result := ""
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if row == 1 && (col == 1 || col == x) {
				result += "A"
			} else if row == y && (col == 1 || col == x) {
				result += "C"
			} else if row == 1 || row == y || col == 1 || col == x {
				result += "B"
			} else {
				result += " "
			}
		}
		result += "\n"
	}
	return result
}

// QuadD pattern
func generateQuadD(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	result := ""
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if col == 1 && (row == 1 || row == y) {
				result += "A"
			} else if col == x && (row == 1 || row == y) {
				result += "C"
			} else if row == 1 || row == y || col == 1 || col == x {
				result += "B"
			} else {
				result += " "
			}
		}
		result += "\n"
	}
	return result
}

// QuadE pattern
func generateQuadE(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	result := ""
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if (row == 1 && col == 1) || (row == y && col == x) {
				result += "A"
			} else if (row == 1 && col == x) || (row == y && col == 1) {
				result += "C"
			} else if row == 1 || row == y || col == 1 || col == x {
				result += "B"
			} else {
				result += " "
			}
		}
		result += "\n"
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	input := strings.Join(lines, "\n")
	if input != "" {
		input += "\n"
	}

	width, height := getDimensions(input)
	if width <= 0 || height <= 0 {
		fmt.Println("Not a quad function")
		return
	}

	patterns := map[string]string{
		"quadA": generateQuadA(width, height),
		"quadB": generateQuadB(width, height),
		"quadC": generateQuadC(width, height),
		"quadD": generateQuadD(width, height),
		"quadE": generateQuadE(width, height),
	}

	var matches []string
	for _, name := range []string{"quadA", "quadB", "quadC", "quadD", "quadE"} {
		if input == patterns[name] {
			matches = append(matches, fmt.Sprintf("[%s] [%d] [%d]", name, width, height))
		}
	}

	if len(matches) == 0 {
		fmt.Println("Not a quad function")
	} else {
		fmt.Println(strings.Join(matches, " || "))
	}
}
