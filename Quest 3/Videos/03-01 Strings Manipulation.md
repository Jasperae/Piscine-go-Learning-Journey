Absolutely brilliant, Jasper — and now let’s make this checkpoint-proof. Since you’re working under strict constraints (no `fmt`, no `strings`, no `unicode`, no `utf8`), I’ve adapted your masterclass to use only what’s allowed: `package main`, `z01.PrintRune`, and pure Go logic.

---

# 🧠 Masterclass: String Manipulation in Go — Mutability with Rune Slices

## 🧩 Part 1: Overview — Strings in Go Are Immutable

| Concept     | Behavior in Go                          | Why It Matters |
|-------------|------------------------------------------|----------------|
| String      | Immutable sequence of bytes              | Cannot modify in-place |
| Rune Slice  | Mutable representation of Unicode strings| Allows safe edits |

Go strings cannot be changed directly. To modify them, convert to a slice of runes, make your edits, and convert back.

---

## 🧩 Part 2: Code Setup — Minimal Program with `z01.PrintRune`

### 🧪 Example
```go
package main

import "github.com/01-edu/z01"

func main() {
	str := "hello"
	for _, r := range str {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
```

✅ Prints: `hello`

---

## 🧩 Part 3: Mutability Strategy — Convert, Modify, Rebuild

### 🔍 Step-by-Step Breakdown

| Step | Code | Result |
|------|------|--------|
| Convert to rune slice | `r := []rune(str)` | `['h','e','l','l','o']` |
| Modify index 0        | `r[0] = 'A'`        | `['A','e','l','l','o']` |
| Convert back          | `str = string(r)`   | `"Aello"` |

### 🧪 Full Example
```go
package main

import "github.com/01-edu/z01"

func main() {
	str := "hello"
	r := []rune(str)
	r[0] = 'A'
	for _, ch := range r {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
```

✅ Prints: `Aello`

---

## 🧩 Part 4: Why Use Runes?

- Runes handle Unicode safely (e.g., `'é'`, `'😊'`)
- Byte slices break multi-byte characters — not allowed in piscine
- Rune slices let you edit any character by index

---

## 🧩 Part 5: Best Practices & Pitfalls

### ✅ Best Practices
- Use `[]rune` for all string edits
- Always loop with `range` to print characters
- Use `z01.PrintRune` for output

### ❌ Common Pitfalls
| Mistake | Why It Fails | Fix |
|--------|---------------|-----|
| Editing string directly | Strings are immutable | Convert to `[]rune` |
| Using byte slice on Unicode | Breaks characters | Use rune slice instead |
| Using `fmt.Println` | Not allowed in checkpoint | Use `z01.PrintRune` |

---

## 🧩 Part 6: Mini Exercises

### 🧪 Exercise 1: Replace first letter of `"world"` with `'M'`
```go
str := "world"
r := []rune(str)
r[0] = 'M'
for _, ch := range r {
	z01.PrintRune(ch)
}
z01.PrintRune('\n') // prints "Morld"
```

### 🧪 Exercise 2: Replace second character of `"héllo"` with `'a'`
```go
str := "héllo"
r := []rune(str)
r[1] = 'a'
for _, ch := range r {
	z01.PrintRune(ch)
}
z01.PrintRune('\n') // prints "hallo"
```

### 🧪 Exercise 3: Build a helper function
```go
func replaceChar(s string, index int, newChar rune) string {
	r := []rune(s)
	r[index] = newChar
	return string(r)
}
```

Use it like:
```go
str := replaceChar("hello", 0, 'Y')
for _, ch := range str {
	z01.PrintRune(ch)
}
z01.PrintRune('\n') // prints "Yello"
```

---

## 🧩 Summary Table

| Step | Action | Result |
|------|--------|--------|
| 1    | Declare string | `"hello"` |
| 2    | Convert to rune slice | `[]rune` |
| 3    | Modify index     | Change character |
| 4    | Loop to print    | Use `z01.PrintRune` |
| 5    | End with newline | `\n` |

---

This is how we build mastery, Jasper — not just by changing strings, but by understanding Go’s type system and working within the constraints of the checkpoint. Want to try building a `ReplaceAll` or `ReverseString` next? I’ve got your back.
