Absolutely phenomenal, Jasper — and now let’s make this piscine-proof. Since you’re working under the constraint of using only `github.com/01-edu/z01`, we’ll adapt this masterclass to reflect the real-world limitations of your checkpoint environment. No `fmt`, no `strconv`, no `unicode` — just raw Go, runes, and `z01.PrintRune`.

---

# 🧠 Piscine Masterclass: Go Variable Declarations, ASCII Encoding & `z01.PrintRune`

## 🧩 Part 1: Why ASCII and Runes Matter in the Piscine

| Concept     | Purpose                                  | Relevance in Checkpoints |
|-------------|------------------------------------------|---------------------------|
| ASCII       | Maps characters to numeric codes         | Helps you print letters using numbers |
| Rune        | Represents a Unicode code point (`int32`) | Required by `z01.PrintRune` |
| `PrintRune` | Outputs a single character               | Your only allowed output tool |

In the piscine, you’re not just learning Go — you’re learning how to think like a compiler. Every character you print must be intentional, type-safe, and rune-correct.

---

## 🧩 Part 2: ASCII — The Numeric Language of Characters

### 🔍 What Is ASCII?

| Character | Decimal Code |
|-----------|--------------|
| A         | 65           |
| B         | 66           |
| a         | 97           |
| 0         | 48           |
| Space     | 32           |

### 🧠 Piscine Tip
- You can print any ASCII character using its code:
  ```go
  z01.PrintRune(rune(65)) // prints A
  ```

---

## 🧩 Part 3: Go Variable Declarations — Short vs Long Form

| Form   | Syntax               | Type Inference |
|--------|----------------------|----------------|
| Short  | `a := 10`            | Go infers `int` |
| Long   | `var b rune = 'A'`   | Explicit rune   |

### 🧠 Piscine Tip
- Use `:=` for quick assignments.
- Use `var` when you want to be explicit or declare zero values.

---

## 🧩 Part 4: Runes — Go’s Character Type

### 🔍 What Is a Rune?
A rune is an alias for `int32` — it represents a single Unicode character.

| Literal | Type | Meaning             |
|---------|------|---------------------|
| `'A'`   | rune | Unicode code point 65 |
| `"A"`   | string | Not allowed in `PrintRune` |

### 🧠 Piscine Tip
- Always use single quotes for characters: `'A'`
- Strings like `"A"` must be looped through rune-by-rune:
  ```go
  for _, r := range "Hello" {
      z01.PrintRune(r)
  }
  ```

---

## 🧩 Part 5: `z01.PrintRune` — Your Output Lifeline

### ✅ Correct Usage
```go
z01.PrintRune('A')        // ✅ prints A
z01.PrintRune(rune(66))   // ✅ prints B
```

### ❌ Incorrect Usage
```go
z01.PrintRune("A")        // ❌ string, not rune
z01.PrintRune(65)         // ❌ int, not rune
```

### 🧠 Piscine Tip
- Always cast integers to `rune` if you’re using ASCII codes.
- Use `'\n'` to print a newline.

---

## 🧩 Part 6: Common Pitfalls in the Piscine

| Mistake                  | Why It Fails              | Fix                     |
|--------------------------|---------------------------|--------------------------|
| `z01.PrintRune("A")`     | `"A"` is a string         | Use `'A'` instead        |
| `z01.PrintRune(65)`      | `int` not accepted        | Use `rune(65)`           |
| `s[0]` without casting   | Returns byte (`uint8`)    | Use `rune(s[0])`         |

---

## 🧩 Part 7: Mini Piscine Exercises

### 🧪 Exercise 1: Declare and print rune A
```go
a := rune(65)
z01.PrintRune(a)
```

### 🧪 Exercise 2: Use rune literal
```go
z01.PrintRune('B')
```

### 🧪 Exercise 3: Convert int to rune and print
```go
var x int = 67
z01.PrintRune(rune(x))  // prints C
```

### 🧪 Exercise 4: Print a full string
```go
for _, r := range "Piscine Go!" {
	z01.PrintRune(r)
}
z01.PrintRune('\n')
```

---

## 🧩 Part 8: Advanced Notes for Piscine Survivors

- You can’t use `unicode.IsLetter`, so build your own logic:
  ```go
  if r >= 'A' && r <= 'Z' {
      // it's an uppercase letter
  }
  ```
- To print digits from an `int`, use math:
  ```go
  n := 5
  z01.PrintRune(rune(n + '0')) // prints '5'
  ```

---

## 🧩 Summary Table

| Concept         | Tool            | Example                  |
|------------------|------------------|---------------------------|
| ASCII            | Encoding         | `rune(65)` → `'A'`        |
| Variable         | Declaration      | `a := rune(66)`           |
| Rune             | Character type   | `'A'`, `rune(65)`         |
| PrintRune        | Output rune      | `z01.PrintRune('A')`      |
| Print string     | Loop with `range`| `for _, r := range s`     |

---

This is how we build piscine mastery, Jasper — not just by printing characters, but by understanding the encoding, the types, and the compiler’s expectations. Ready to build your own `PrintInt`, `PrintBool`, or even `PrintHex` next? Let’s keep sharpening your toolkit.
