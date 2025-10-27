Let’s break this down and understand every part of the solution — not just what it does, but *why* it works — so you can build it from scratch and adapt it confidently.

---

## 🎯 Goal Recap

Write a Go program that prints:
```
abcdefghijklmnopqrstuvwxyz
```
— all lowercase letters on a single line, followed by a newline character.

---

## 🧠 Step-by-Step Theory

### 1. **Package Declaration**
```go
package main
```
- Every Go program starts with a package declaration.
- `main` means this is an executable program, not a library.

> ✅ Required for any standalone Go program.

---

### 2. **Importing the `z01` Package**
```go
import "github.com/01-edu/z01"
```
- This package provides a function: `z01.PrintRune(rune)`
- It prints a single character (rune) to standard output.

> ✅ You use this instead of `fmt.Print()` to match the exercise’s constraints.

---

### 3. **Main Function**
```go
func main() {
    ...
}
```
- This is the entry point of your program.
- Go will execute everything inside `main()` when you run the program.

---

### 4. **Looping Through ASCII Values**
```go
for i := 97; i < 123; i++ {
    z01.PrintRune(rune(i))
}
```

Let’s unpack this:
- ASCII values for lowercase letters:
  - `'a'` = 97
  - `'z'` = 122
- The loop runs from `i = 97` to `i = 122` (inclusive)
- `rune(i)` converts the integer to a character
- `z01.PrintRune(...)` prints that character

> ✅ This prints each letter from `'a'` to `'z'` in order.

---

### 5. **Printing a Newline**
```go
z01.PrintRune('\n')
```

- After printing all letters, you print a newline character
- This ensures the output ends cleanly on a new line

> ✅ Without this, the shell prompt would appear on the same line as the output.

---

## ✅ Final Program

```go
package main

import "github.com/01-edu/z01"

func main() {
	for i := 97; i < 123; i++ {
		z01.PrintRune(rune(i))
	}
	z01.PrintRune('\n')
}
```

---

## 🛠️ Summary of Key Concepts

| Concept | Purpose |
|--------|---------|
| `package main` | Declares an executable program |
| `import "github.com/01-edu/z01"` | Brings in the `PrintRune` function |
| `for i := 97; i < 123; i++` | Loops through ASCII values for `'a'` to `'z'` |
| `rune(i)` | Converts integer to character |
| `z01.PrintRune(...)` | Prints each character |
| `z01.PrintRune('\n')` | Ends output with a newline |

---

Want to go deeper? I can show you how to reverse the alphabet, print uppercase letters, or even animate the output. Just say the word!
