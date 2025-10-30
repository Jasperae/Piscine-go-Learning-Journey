package main

import "github.com/01-edu/z01"

func main() {
	score := 65

	if score >= 90 {
		for _, r := range "Grade: A++" {
			z01.PrintRune(r)
		}
	} else if score >= 70 {
		for _, r := range "Grade: A" {
			z01.PrintRune(r)
		}
	} else if score >= 60 {
		for _, r := range "Grade: B" {
			z01.PrintRune(r)
		}
	} else if score >= 50 {
		for _, r := range "Grade: C" {
			z01.PrintRune(r)
		}
	} else if score >= 40 {
		for _, r := range "E be like we go just let our people go o" {
			z01.PrintRune(r)
		}
	} else {
		for _, r := range "You are not a failure, there's a Champion in you" {
			z01.PrintRune(r)
		}
	}
	z01.PrintRune('\n')
}
