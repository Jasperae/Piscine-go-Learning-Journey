Absolutely brilliant, Jasper — and now let’s tailor this masterclass to your **piscine checkpoint reality**, where you're restricted to using only the `github.com/01-edu/z01` package. That constraint turns character handling into a true test of your understanding of Go’s type system, Unicode encoding, and low-level output control.

---

# 🧠 Piscine Masterclass: Character Handling in Go with `z01.PrintRune`

## 🧩 Part 1: Why `PrintRune` Is a Big Deal

| Concept     | Purpose                                      | Why It’s Important |
|-------------|----------------------------------------------|---------------------|
| `PrintRune` | Outputs a single Unicode code point (rune)   | Forces precision and encoding awareness |
| Rune        | Represents a Unicode character (`int32`)     | Enables exact character-level control |
| String      | Sequence of bytes or runes                   | Must be manually unpacked for output |

In the piscine, you can’t use `fmt.Println`, `strings`, or `strconv`. So `z01.PrintRune` becomes your only tool for output — and mastering it means mastering Go’s character system.

---

## 🧩 Part 2: Runes vs Strings — Know the Difference

### 🔍 What Is a Rune?
A rune is an alias for `int32` and represents a single Unicode character.

| Literal | Type   | Meaning                     |
|---------|--------|-----------------------------|
| `'A'`   | rune   | Unicode code point 65       |
| `"A"`   | string | Sequence of bytes           |

### 🧠 Piscine Tip
- Use single quotes (`'A'`) for runes — required by `PrintRune`
- Double quotes (`"A"`) are strings — not accepted by `PrintRune`
- You must loop through strings to print them character by character

---

## 🧩 Part 3: Correct Usage of `z01.PrintRune`

### ✅ Valid Examples
```go
z01.PrintRune('A')          // prints A
z01.PrintRune(rune(65))     // prints A
```

### ❌ Invalid Examples
```go
z01.PrintRune("A")          // ❌ string, not rune
z01.PrintRune(65)           // ❌ int, not rune
```

---

## 🧩 Part 4: Printing a Full String

Since `PrintRune` only prints one character at a time, you must loop through strings manually.

### ✅ Helper Function
```go
func PrintString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
```

### 🧪 Usage
```go
PrintString("You sef tall o!")
z01.PrintRune('\n')
```

---

## 🧩 Part 5: Indexing and Casting

### 🔍 Print First Character of a String
```go
s := "Hello"
z01.PrintRune(rune(s[0])) // prints H
```

- `s[0]` returns a byte (`uint8`)
- You must cast it to `rune` to use with `PrintRune`

---

## 🧩 Part 6: Looping Through a String

### 🔁 Byte-by-Byte (Manual)
```go
s := "Hello"
for i := 0; i < len(s); i++ {
	z01.PrintRune(rune(s[i]))
}
```

### 🔁 Unicode-Safe (Recommended)
```go
for _, r := range "Gó!" {
	z01.PrintRune(r)
}
```

- Handles multi-byte characters like `ó`, `ç`, or emojis
- Essential for international text or special symbols

---

## 🧩 Part 7: Common Pitfalls in the Piscine

| Mistake | Why It Fails | Fix |
|--------|---------------|-----|
| `"A"` instead of `'A'` | String, not rune | Use single quotes |
| `PrintRune(65)`        | Int, not rune    | Cast: `rune(65)` |
| `s[0]` without casting | Returns byte     | Use `rune(s[0])` |

---

## 🧩 Part 8: Mini Piscine Exercises

### 🧪 Exercise 1: Print 'G' using ASCII
```go
z01.PrintRune(rune(71)) // prints G
```

### 🧪 Exercise 2: Print each character of "Go!"
```go
s := "Go!"
for _, r := range s {
	z01.PrintRune(r)
}
z01.PrintRune('\n')
```

### 🧪 Exercise 3: Print a Unicode string
```go
PrintString("Gó rocks!")
z01.PrintRune('\n')
```

---

## 🧩 Part 9: Advanced Notes for Piscine Survivors

- Use `[]rune(s)` to convert a string into a rune slice for safe indexing
- You can build your own `PrintInt`, `PrintBool`, and `PrintSlice` using `PrintRune`
- Runes are UTF-8 aware — they support emojis, accented letters, and symbols

---

## 🧩 Summary Table

| Task                        | Code Example                     | Output |
|-----------------------------|----------------------------------|--------|
| Print rune literal          | `z01.PrintRune('A')`             | A      |
| Print from ASCII code       | `z01.PrintRune(rune(66))`        | B      |
| Print first char of string  | `z01.PrintRune(rune(s[0]))`      | H      |
| Print full string (manual)  | loop with `rune(s[i])`           | Hello  |
| Print full string (Unicode) | `for _, r := range s`            | Gó!    |

---

This is how we build piscine mastery, Jasper — not just by printing characters, but by understanding the encoding, the types, and the mechanics behind every output. Ready to build your own `PrintHex`, `PrintChar`, or even a custom formatter next? Let’s keep sharpening your toolkit.
