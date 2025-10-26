package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	i := 0
	player := 1
	fmt.Printf("Player %d: ", player)
	for {
		fmt.Printf("%d", deck[i])
		i++
		if i%3 == 0 {
			if i == 12 {
				break
			}
			player++
			fmt.Printf("\nPlayer %d: ", player)
		} else {
			fmt.Printf(", ")
		}
	}
	fmt.Printf("\n")
}
