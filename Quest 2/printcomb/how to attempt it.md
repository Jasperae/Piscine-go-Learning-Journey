Let’s break this down so you understand *why* every part of the `PrintComb()` function works — not just what it does. By the end, you’ll be able to build this kind of logic from scratch and even extend it to more complex problems.

---

## 🎯 Goal Recap

Write a function that prints all **unique combinations** of three **different digits** (0–9) such that:
- The digits are in **strictly increasing order**: `a < b < c`
- Each combination is printed in ascending order
- Combinations are separated by `, ` — except the last one
- Output is on a **single line**, ending with a newline

---

## 🧠 Step-by-Step Theory

### 1. **Using Runes for Digits**
```go
for a := '0'; a <= '9'; a++
```

- `'0'` to `'9'` are rune literals (characters), not integers.
- Using runes directly allows you to print them with `z01.PrintRune()` without converting from int to rune.

> ✅ This simplifies printing and keeps the code readable.

---

### 2. **Nested Loops for Unique Combinations**
```go
for a := '0'; a <= '9'; a++ {
	for b := a + 1; b <= '9'; b++ {
		for c := b + 1; c <= '9'; c++ {
			// Print combination
		}
	}
}
```

- `b := a + 1` ensures `b > a`
- `c := b + 1` ensures `c > b`
- This guarantees `a < b < c` — the required condition

> ✅ This avoids duplicates and invalid combinations like `000`, `999`, or `987`.

---

### 3. **Printing Digits with `z01.PrintRune`**
```go
z01.PrintRune(a)
z01.PrintRune(b)
z01.PrintRune(c)
```

- Each digit is printed in order
- Since they’re runes, no conversion is needed

---

### 4. **Handling the Separator**
```go
if a != '7' || b != '8' || c != '9' {
	z01.PrintRune(',')
	z01.PrintRune(' ')
}
```

- The last valid combination is `789`
- This condition ensures that `, ` is printed **only** if it’s not the last combination

> ✅ Prevents trailing comma after the final output.

---

### 5. **Unnecessary Space Before First Combination**
```go
if a != '0' || b != '1' || c != '2' {
	z01.PrintRune(' ')
}
```

- This prints a space before every combination except the first (`012`)
- But it’s not necessary — the separator logic already handles spacing correctly

> ⚠️ This line can be removed for cleaner logic.

---

### 6. **Ending the Output**
```go
z01.PrintRune('\n')
```

- Prints a newline after all combinations
- Ensures the shell prompt appears on a new line

---

## ✅ Cleaned-Up Version

Here’s a cleaner version of the function with the unnecessary space logic removed:

```go
package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for a := '0'; a <= '9'; a++ {
		for b := a + 1; b <= '9'; b++ {
			for c := b + 1; c <= '9'; c++ {
				z01.PrintRune(a)
				z01.PrintRune(b)
				z01.PrintRune(c)
				if a != '7' || b != '8' || c != '9' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
```

---

## 🛠️ Summary of Key Concepts

| Concept | Purpose |
|--------|---------|
| `'0'` to `'9'` | Rune literals for digits |
| `a < b < c` | Ensures digits are increasing |
| `z01.PrintRune()` | Prints single characters |
| `if a != '7' || b != '8' || c != '9'` | Skips separator after last combo |
| `\n` | Ends output with newline |

---

Want to go deeper? I can show you how to generate combinations of letters, handle variable-length combos, or even animate the output. Just say the word!
