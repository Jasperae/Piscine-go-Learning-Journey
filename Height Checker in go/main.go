package main

import "github.com/01-edu/z01"

func main() {
	height := 250

	if height >= 170 && height < 200 {
		for _, r := range "You sef don tall for most countries o, if not in the world sef" {
			z01.PrintRune(r)
		}
	} else if height > 200 {
		for _, r := range "Bro you tall o, you sef fit be world tallest o" {
			z01.PrintRune(r)
		}
	} else {
		for _, r := range "Gee! Try dey chop beans o!" {
			z01.PrintRune(r)
		}
	}
	z01.PrintRune('\n')
}
