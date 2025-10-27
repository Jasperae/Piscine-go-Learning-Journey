package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')             // Print negative sign
		if n == -9223372036854775808 { // Edge case for smallest int64
			// Handle division by 10 manually to avoid negative recursive call
			PrintNbr(922337203685477580) // Print all but last digit
			z01.PrintRune('8')           // Print last digit manually
			return
		}
		n = -n // Convert to positive
	}

	if n >= 10 {
		PrintNbr(n / 10) // Recursive call for higher digits
	}
	z01.PrintRune('0' + rune(n%10)) // Print last digit as a rune
}
