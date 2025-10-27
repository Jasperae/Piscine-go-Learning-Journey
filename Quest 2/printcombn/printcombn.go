package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n > 10 { // Ensure valid n (digits 1-9)
		return
	}

	// Generate initial sequence [0, 1, 2, ..., n]
	var digits [10]rune
	for i := range digits[:n] {
		digits[i] = '0' + rune(i)
	}

	printComb(digits[:n], n)

	z01.PrintRune('\n') // Ensure newline at the end
}

// Helper function to generate and print combinations recursively
func printComb(digits []rune, n int) {
	printDigits(digits) // Print current combination

	if digits[0] == '9'+1-rune(n) { // Stop condition: highest possible start digit
		return
	}

	// Find the rightmost position that can be incremented
	for i := n - 1; i >= 0; i-- {
		if digits[i] < '9'+rune(i+1-n) {
			digits[i]++
			for j := i + 1; j < n; j++ {
				digits[j] = digits[j-1] + 1
			}
			printComb(digits, n) // Recursively call for next combination
			break
		}
	}
}

// Prints digits and adds ", " unless it's the last combination
func printDigits(digits []rune) {
	for _, d := range digits {
		z01.PrintRune(d)
	}

	// Ensure ", " is printed except for the last combination
	if digits[0] != '9'+1-rune(len(digits)) {
		z01.PrintRune(',')
		z01.PrintRune(' ')
	}
}
