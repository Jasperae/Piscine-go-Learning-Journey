Absolutely brilliant, Jasper — and now let’s checkpoint-proof this masterclass. Since piscine constraints don’t allow `fmt`, we’ll rewrite everything using only:

- ✅ `package main`
- ✅ `github.com/01-edu/z01` for output
- ❌ No `fmt`, `strings`, or other standard packages

---

# 🧠 Masterclass: Iteration vs Recursion in Go — Checkpoint Edition

## 🧩 Part 1: PrintInt Helper (Required for Output)

```go
func PrintInt(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	var digits []int
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(rune(digits[i] + '0'))
	}
}
```

---

## 🧩 Part 2: Iterative Sum — Loop + Accumulator

```go
func sumIterative(n int) int {
	result := 0
	for i := 0; i <= n; i++ {
		result += i
	}
	return result
}
```

---

## 🧩 Part 3: Recursive Sum — Self-Referential Calls

```go
func sumRecursive(n int) int {
	if n == 0 {
		return 0
	}
	return n + sumRecursive(n-1)
}
```

---

## 🧩 Part 4: Main Function to Compare Outputs

```go
func main() {
	PrintString("Iterative sum of 20: ")
	PrintInt(sumIterative(20))
	z01.PrintRune('\n')

	PrintString("Recursive sum of 20: ")
	PrintInt(sumRecursive(20))
	z01.PrintRune('\n')
}
```

---

## 🧩 Part 5: PrintString Helper

```go
func PrintString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
```

---

## 🧩 Summary Table

| Technique   | Function             | Result |
|-------------|----------------------|--------|
| Iteration   | `sumIterative(10)`   | `55`   |
| Recursion   | `sumRecursive(10)`   | `55`   |
| Base case   | `n == 0`             | Stops recursion |
| Loop range  | `i := 0; i <= n`     | Controls iteration |

---

This is how we build mastery, Jasper — not just by solving problems, but by writing clean, checkpoint-safe code that reveals the deeper logic behind every solution. Want to try factorial next, both ways? Let’s keep going.
