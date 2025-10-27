Absolutely! Let’s break this down so you understand **every decision and every line** of the `FirstRune` function—not just how it works, but **why** it works. By the end, you’ll be able to build it from scratch and even adapt it for more advanced string handling.

---

## 🧠 Goal of the Function

We want to write a function `FirstRune(s string) rune` that:

- Takes a string `s`
- Returns the **first rune** (Unicode character) of that string

### Example:
```go
FirstRune("Hello!") → 'H'
FirstRune("Salut!") → 'S'
FirstRune("Ola!") → 'O'
```

---

## ✅ Full Code

```go
func FirstRune(s string) rune {
	r := []rune(s)
	return r[0]
}
```

---

## 🔍 Step-by-Step Breakdown

### 1. `func FirstRune(s string) rune {`

- This defines a function named `FirstRune`.
- It takes a string `s` and returns a single `rune`.

### What is a `rune`?

- A `rune` in Go is an alias for `int32`.
- It represents a **Unicode code point**—a single character, which may be more than one byte.
- Go strings are made of **bytes**, but `rune` lets us work with full characters, including non-ASCII ones like `é`, `ç`, `你`.

---

### 2. `r := []rune(s)`

- Converts the string `s` into a slice of runes.
- Why do this?
  - Go strings are byte slices (`[]byte`), so indexing directly (`s[0]`) gives you the first byte—not necessarily the first character.
  - For example, `"é"` is two bytes in UTF-8, but one rune.
  - Converting to `[]rune` ensures we treat each character properly.

### Example:
```go
s := "Hello!"
r := []rune(s) // ['H', 'e', 'l', 'l', 'o', '!']
```

---

### 3. `return r[0]`

- Returns the first rune from the slice.
- This is the first character of the string, correctly interpreted even if it’s multi-byte.

---

## 🧪 Example Outputs

```go
FirstRune("Hello!") → 'H'
FirstRune("Salut!") → 'S'
FirstRune("Ola!") → 'O'
FirstRune("你好") → '你' (correctly handles Unicode)
```

---

## 🧼 Optional Enhancements

### 1. Handle empty strings safely:

If you want to avoid panics when `s` is empty:

```go
func FirstRune(s string) rune {
	r := []rune(s)
	if len(r) == 0 {
		return 0 // or return a default rune
	}
	return r[0]
}
```

---

## ✅ Summary

| Concept | Purpose |
|--------|---------|
| `rune` | Represents a full Unicode character |
| `[]rune(s)` | Converts string to character-aware slice |
| `r[0]` | Safely accesses the first character |
| Unicode-safe | Works with multi-byte characters like emojis or accented letters |

---

Want to explore how to get the **last rune**, or how to iterate through each rune in a string with proper formatting? I’d love to walk you through that next.
