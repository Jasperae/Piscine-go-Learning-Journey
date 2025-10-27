Perfect — let’s turn this from memorizing a few lines of Go code into fully understanding how to build this function from scratch. We’ll break down every part of the solution and explain the *why* behind it so you can confidently write and adapt similar logic anytime.

---

## 🎯 Objective Recap

You’re writing a function:
```go
func IsNegative(nb int)
```
That prints:
- `'T'` if the input `nb` is negative
- `'F'` if it’s zero or positive

Each result should appear on its own line.

---

## 🧠 Step-by-Step Theory

### 1. **Function Signature**
```go
func IsNegative(nb int)
```

- `func` declares a function
- `IsNegative` is the function name
- `nb int` means the function takes one parameter: an integer named `nb`
- No return value — the function prints directly to output

> ✅ This matches the expected usage in the test program.

---

### 2. **Importing the `z01` Package**
```go
import "github.com/01-edu/z01"
```

- This package provides `z01.PrintRune(rune)` — a function to print a single character
- It’s used instead of `fmt.Println()` to match the exercise’s constraints

> ✅ You’ll use `PrintRune` to print `'T'`, `'F'`, and `'\n'`.

---

### 3. **Conditional Logic**
```go
if nb < 0 {
    z01.PrintRune('T')
} else {
    z01.PrintRune('F')
}
```

- `if nb < 0` checks whether the input is negative
- If true, print `'T'`
- Otherwise, print `'F'`

> ✅ This is the core decision-making logic.

---

### 4. **Printing a Newline**
```go
z01.PrintRune('\n')
```

- After printing `'T'` or `'F'`, you print a newline character
- This ensures each result appears on its own line

> ✅ Without this, all outputs would appear on the same line.

---

## ✅ Final Function

```go
package piscine

import "github.com/01-edu/z01"

func IsNegative(nb int) {
	if nb < 0 {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
	z01.PrintRune('\n')
}
```

---

## 🧪 Example Execution

Given:
```go
piscine.IsNegative(1)
piscine.IsNegative(0)
piscine.IsNegative(-1)
```

Output:
```
F
F
T
```

---

## 🛠️ Summary of Key Concepts

| Concept | Purpose |
|--------|---------|
| `func IsNegative(nb int)` | Defines the function |
| `nb < 0` | Checks if the number is negative |
| `z01.PrintRune('T')` | Prints `'T'` if true |
| `z01.PrintRune('F')` | Prints `'F'` if false |
| `z01.PrintRune('\n')` | Moves to the next line |

---

Want to go deeper? I can show you how to write similar functions like `IsEven`, `IsPositive`, or even how to build a full diagnostic tool using this logic. Just say the word!
