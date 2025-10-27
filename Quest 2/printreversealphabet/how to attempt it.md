Let’s break this down so you understand *why* every part of this Go program works — not just what it does. By the end, you’ll be able to build similar logic from scratch and adapt it to any character-based output.

---

## 🎯 Goal Recap

Write a Go program that prints:
```
zyxwvutsrqponmlkjihgfedcba
```
— all lowercase letters of the Latin alphabet in **reverse order**, followed by a newline.

---

## 🧠 Step-by-Step Theory

### 1. `package main`

This line tells Go:
> “This is an executable program.”

- All Go programs must start with a `package` declaration.
- `main` is the special package name for programs that can be run directly.

---

### 2. `import "github.com/01-edu/z01"`

This brings in the `z01` package, which provides:
```go
z01.PrintRune(rune)
```

- `PrintRune` prints a single character (rune) to the output.
- You use this instead of `fmt.Print()` because the exercise requires it.

> ✅ Think of `z01.PrintRune('A')` as: “Print the character A.”

---

### 3. `func main()`

This is the entry point of your program:
- Go looks for `main()` and starts execution there.
- Everything inside `main()` runs when you do `go run .`

---

### 4. The `for` Loop

```go
for i := 'z'; i >= 'a'; i-- {
	z01.PrintRune(rune(i))
}
```

Let’s unpack this:
- `'z'` and `'a'` are **rune literals** — characters, not strings.
- In Go, characters are just numbers under the hood (ASCII values):
  - `'z'` = 122
  - `'y'` = 121
  - ...
  - `'a'` = 97

So this loop:
- Starts at `'z'`
- Ends at `'a'`
- Decrements by 1 each time
- Prints each character using `z01.PrintRune`

> ✅ This prints: `z y x ... a` — all on the same line.

### ⚠️ Note on Casting

The instruction says **casting is not allowed** — but this solution uses:
```go
z01.PrintRune(rune(i))
```

That’s technically a cast from `i` to `rune`, but since `i` is already a rune (declared as `i := 'z'`), this cast is redundant and can be removed:
```go
z01.PrintRune(i)
```

> ✅ Cleaner and fully compliant with the no-casting rule.

---

### 5. `z01.PrintRune('\n')`

After the loop, this line prints a newline character:
- `\n` moves the cursor to the next line
- Without it, the shell prompt would appear on the same line as your output

> ✅ This ensures clean output formatting.

---

## ✅ Final Program (Cleaned Version)

```go
package main

import "github.com/01-edu/z01"

func main() {
	for i := 'z'; i >= 'a'; i-- {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
```

---

## 🛠️ Summary of Key Concepts

| Concept | Purpose |
|--------|---------|
| `package main` | Declares an executable Go program |
| `import "z01"` | Brings in the PrintRune function |
| `for i := 'z'; i >= 'a'; i--` | Loops through ASCII values for letters in reverse |
| `z01.PrintRune(i)` | Prints each letter |
| `z01.PrintRune('\n')` | Ends output with a newline |

---

Want to go further? I can show you how to print only vowels, alternate letters, or even animate the output. Just say the word!
