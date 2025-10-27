Absolutely! Let’s break down this `sortparams` program line by line and uncover the **why** behind every part. By the end, you’ll understand how it works from the inside out—and be able to rebuild or adapt it confidently.

---

## 🧠 Goal of the Program

We want to write a Go program that:

- Accepts command-line arguments
- Sorts them in **ASCII order**
- Prints each argument on a **new line**
- Uses `z01.PrintRune` to print characters one by one (as required in 01-edu projects)

---

## ✅ Full Code

```go
package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	sortStringArr(args)
	printStringArr(args)
}
```

---

## 🔍 Step-by-Step Breakdown

### 1. `os.Args[1:]`

- `os.Args` is a slice of strings:
  - `os.Args[0]` is the program name (e.g., `./sortparams`)
  - `os.Args[1:]` are the actual user-provided arguments
- We slice off the first element to get only the arguments to sort and print.

---

### 2. `sortStringArr(args)`

This function sorts the arguments in-place using a **simple bubble sort**:

```go
func sortStringArr(args []string) {
	for i := 0; i < len(args); i++ {
		for j := i + 1; j < len(args); j++ {
			if compare(args[j], args[i]) {
				args[i], args[j] = args[j], args[i]
			}
		}
	}
}
```

- It compares each pair of elements.
- If `args[j]` is less than `args[i]` (in ASCII order), it swaps them.
- This continues until the entire slice is sorted.

---

### 3. `compare(a, b string) bool`

This function compares two strings **character by character**:

```go
func compare(a, b string) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}
```

- It loops through both strings until it finds a character that differs.
- Returns `true` if `a` should come before `b` in ASCII order.
- If all characters are equal up to the length of the shorter string, the shorter string comes first.

### Why ASCII order?

- ASCII values determine the sort order:
  - Digits (`0–9`) come before uppercase letters (`A–Z`)
  - Uppercase letters come before lowercase letters (`a–z`)
- So `"1" < "A" < "a"` in ASCII

---

### 4. `printStringWz01(s string)`

```go
func printStringWz01(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}
```

- Loops through each character in the string and prints it using `z01.PrintRune`.

---

### 5. `printStringArr(args []string)`

```go
func printStringArr(args []string) {
	for _, arg := range args {
		printStringWz01(arg)
		z01.PrintRune('\n')
	}
}
```

- Loops through each sorted argument
- Prints the argument followed by a newline

---

## 🧪 Example Execution

```bash
$ go run . 1 a 2 A 3 b 4 C
1
2
3
4
A
C
a
b
```

---

## ✅ Summary

| Concept | Purpose |
|--------|---------|
| `os.Args[1:]` | Accesses command-line arguments |
| Bubble sort | Sorts strings in ASCII order |
| `compare(a, b)` | Determines ASCII order between two strings |
| `z01.PrintRune` | Prints each character individually |
| `'\n'` | Adds a line break after each word |

---

Want to explore how to sort in reverse order, or how to ignore case while sorting? I’d love to walk you through that next.
