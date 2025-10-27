Absolutely! Let’s break this down so you understand **every single part** of the solution—not just how it works, but **why** it works. By the end, you’ll be able to build it from scratch and even adapt it for other formatting tasks.

---

## 🧠 Goal of the Function

We want to write a function `Capitalize(s string) string` that:

- Capitalizes the **first letter of each word**
- Lowercases the **rest of the letters**
- A **word** is defined as a sequence of **alphanumeric characters** (letters and digits)
- Non-alphanumeric characters (spaces, punctuation, symbols) are **not part of words** and act as **word boundaries**

---

## ✅ Full Code

```go
func isAlphanumeric(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}

func Capitalize(s string) string {
	result := make([]byte, len(s))
	inWord := false

	for i := 0; i < len(s); i++ {
		if isAlphanumeric(s[i]) {
			if !inWord {
				// First character of a word
				if s[i] >= 'a' && s[i] <= 'z' {
					result[i] = s[i] - ('a' - 'A') // convert to uppercase
				} else {
					result[i] = s[i] // already uppercase or digit
				}
				inWord = true
			} else {
				// Inside a word
				if s[i] >= 'A' && s[i] <= 'Z' {
					result[i] = s[i] + ('a' - 'A') // convert to lowercase
				} else {
					result[i] = s[i] // already lowercase or digit
				}
			}
		} else {
			// Non-alphanumeric character
			result[i] = s[i]
			inWord = false
		}
	}
	return string(result)
}
```

---

## 🔍 Step-by-Step Breakdown

### 1. `isAlphanumeric(c byte) bool`

- Helper function to check if a character is **alphanumeric**.
- Returns `true` if the character is:
  - A–Z (uppercase)
  - a–z (lowercase)
  - 0–9 (digit)

### Why use this?

- It defines what counts as a word character.
- Everything else (like `!`, `+`, `?`, space) is a **word boundary**.

---

### 2. `result := make([]byte, len(s))`

- Creates a new byte slice to store the transformed characters.
- Same length as the input string.

### Why use a byte slice?

- Go strings are immutable, so we build a new one.
- Using a byte slice is efficient for ASCII-based transformations.

---

### 3. `inWord := false`

- Tracks whether we’re **inside a word**.
- Helps decide whether to capitalize or lowercase a character.

---

### 4. `for i := 0; i < len(s); i++ {`

- Iterates through each character in the string.

---

### 5. `if isAlphanumeric(s[i]) {`

- If the character is part of a word:
  - If `inWord == false`, it’s the **start of a word**
  - If `inWord == true`, it’s **inside a word**

---

### 6. Capitalization Logic

#### First character of a word:
```go
if !inWord {
	if s[i] >= 'a' && s[i] <= 'z' {
		result[i] = s[i] - ('a' - 'A') // convert to uppercase
} else {
		result[i] = s[i] // already uppercase or digit
}
inWord = true
```

#### Inside a word:
```go
} else {
	if s[i] >= 'A' && s[i] <= 'Z' {
		result[i] = s[i] + ('a' - 'A') // convert to lowercase
	} else {
		result[i] = s[i] // already lowercase or digit
	}
}
```

---

### 7. Non-alphanumeric characters:
```go
} else {
	result[i] = s[i]
	inWord = false
}
```

- Copy the character as-is.
- Reset `inWord` to `false` to mark the end of a word.

---

### 8. `return string(result)`

- Converts the byte slice back to a string and returns it.

---

## 🧪 Example: `"Hello! How are you? How+are+things+4you?"`

Breakdown:
- `"Hello"` → `H` is capitalized, `ello` is lowercased
- `"How"` → same
- `"are"` → same
- `"you"` → same
- `"How+are+things+4you"` → each word after `+` is capitalized

Output:
```
Hello! How Are You? How+Are+Things+4you?
```

---

## ✅ Summary

| Concept | Purpose |
|--------|---------|
| `isAlphanumeric` | Defines word boundaries |
| `inWord` flag | Tracks start vs inside of word |
| ASCII math | Efficient case conversion |
| Byte slice | Builds new string efficiently |
| Word logic | Capitalize first letter, lowercase rest |

---

Want to adapt this to work with **Unicode characters**, or apply different rules like title casing or sentence casing? I’d love to walk you through that next.
