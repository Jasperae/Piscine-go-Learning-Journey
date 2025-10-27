Let’s break this down so you understand *why* every part of the `PrintNbr` function works — not just what it does. This will help you build similar logic from scratch and confidently handle edge cases in Go.

---

## 🎯 Goal Recap

Write a function:
```go
func PrintNbr(n int)
```
That prints any integer — positive, negative, or zero — using only:
- The `int` type (no `int64`)
- The `z01.PrintRune()` function to print characters

---

## 🧠 Step-by-Step Theory

### 1. **Handling Negative Numbers**
```go
if n < 0 {
	z01.PrintRune('-')
	...
}
```

- If the number is negative, print a minus sign first.
- Then convert it to positive so you can process its digits.

> ✅ This ensures correct formatting for negative numbers.

---

### 2. **Handling the Edge Case: Minimum `int` Value**
```go
if n == -9223372036854775808 {
	PrintNbr(922337203685477580)
	z01.PrintRune('8')
	return
}
```

- In Go, the smallest `int` value is `-9223372036854775808` on 64-bit systems.
- You **cannot** negate this value directly — it causes overflow.
- So you manually print all but the last digit (`922337203685477580`) and then append `'8'`.

> ✅ This avoids overflow and handles the edge case safely.

---

### 3. **Converting to Positive**
```go
n = -n
```

- After printing the minus sign and checking for the edge case, you convert `n` to positive.
- This simplifies digit extraction.

---

### 4. **Recursive Digit Printing**
```go
if n >= 10 {
	PrintNbr(n / 10)
}
```

- This recursively prints all digits **except the last one**.
- For example, `123` becomes:
  - `PrintNbr(12)` → `PrintNbr(1)` → prints `'1'`
  - Then prints `'2'`
  - Then prints `'3'`

> ✅ This ensures digits are printed in the correct order.

---

### 5. **Printing the Last Digit**
```go
z01.PrintRune('0' + rune(n%10))
```

- `n % 10` gives the last digit
- `'0' + rune(n%10)` converts it to the correct character
  - `'0'` is ASCII 48
  - If `n % 10 == 3`, this becomes `'3'`

> ✅ This prints the digit as a character.

---

## ✅ Final Function Recap

```go
package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		if n == -9223372036854775808 {
			PrintNbr(922337203685477580)
			z01.PrintRune('8')
			return
		}
		n = -n
	}

	if n >= 10 {
		PrintNbr(n / 10)
	}
	z01.PrintRune('0' + rune(n%10))
}
```

---

## 🛠️ Summary of Key Concepts

| Concept | Purpose |
|--------|---------|
| `n < 0` | Detects negative numbers |
| `z01.PrintRune('-')` | Prints minus sign |
| Edge case check | Handles minimum `int` value safely |
| `n = -n` | Converts to positive for digit processing |
| `PrintNbr(n / 10)` | Recursively prints higher digits |
| `'0' + rune(n % 10)` | Converts digit to character |

---

Want to go deeper? I can show you how to format numbers with commas, print floating-point values, or even build your own version of `itoa`. Just say the word!
