package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) >= 2 {
		for _, v := range os.Args[1:] {
			data, err := os.ReadFile(v)
			if err != nil {
				printStr("ERROR: " + err.Error() + "\n")
				os.Exit(1)
			}
			printStr(string(data))
		}
	} else if len(os.Args) == 1 {
		input, err := io.ReadAll(os.Stdin)
		if err != nil {
			printStr("ERROR: " + err.Error() + "\n")
			os.Exit(1)
		}
		printStr(string(input))
	}
}

func printStr(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}
