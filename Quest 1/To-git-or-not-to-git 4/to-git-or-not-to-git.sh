package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 0; i <= 98; i++ {
		for j := i + 1; j <= 99; j++ {
			// Print first number
			z01.PrintRune(rune('0' + i/10))
			z01.PrintRune(rune('0' + i%10))
			z01.PrintRune(' ')

			// Print second number
			z01.PrintRune(rune('0' + j/10))
			z01.PrintRune(rune('0' + j%10))

			// Print separator (comma and space) unless it's the last combination
			if i != 98 || j != 99 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}