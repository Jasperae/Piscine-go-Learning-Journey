Absolutely brilliant, Jasper — and now let’s tailor this masterclass to your **piscine checkpoint reality**, where you're restricted to using only the `github.com/01-edu/z01` package. That constraint transforms even basic iteration into a deeper exercise in precision, creativity, and low-level control.

---

# 🧠 Piscine Masterclass: Iteration in Go — The Power of the `for` Loop with `z01.PrintRune`

## 🧩 Part 1: Overview — One Loop to Rule Them All

Go simplifies iteration with a single loop construct: `for`. Whether you're counting, looping forever, or iterating over collections — it’s all done with `for`.

| Language | Loop Types         | Go’s Approach |
|----------|--------------------|----------------|
| C, Java  | `for`, `while`, `do-while` | Unified `for` loop |
| Go       | `for` only         | Handles all loop patterns |

In your piscine checkpoints, this simplicity is paired with a constraint: **you must use `z01.PrintRune` for all output**, which means printing numbers, strings, and even newlines must be done manually.

---

## 🧩 Part 2: Anatomy of a Go `for` Loop

### 🔍 Syntax Breakdown
```go
for initializer; condition; post {
    // loop body
}
```

| Part           | Role                              | Example        |
|----------------|-----------------------------------|----------------|
| Initializer    | Sets up loop variable             | `i := 0`       |
| Condition      | Controls loop continuation        | `i < 10`       |
| Post statement | Updates loop variable             | `i++`          |

---

## 🧪 Piscine Example: Print 0 to 9 Using `z01.PrintRune`

```go
package main

import "github.com/01-edu/z01"

func PrintInt(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
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

func main() {
	for i := 0; i < 10; i++ {
		PrintInt(i)
		z01.PrintRune('\n')
	}
}
```

### ✅ Output
```
0
1
2
...
9
```

---

## 🧩 Part 3: Loop Variants You Can Use

### 🔁 Loop Without Initializer
```go
i := 0
for ; i < 5; i++ {
	PrintInt(i)
	z01.PrintRune('\n')
}
```

### 🔁 Loop Without Post Statement
```go
for i := 0; i < 5; {
	PrintInt(i)
	z01.PrintRune('\n')
	i++
}
```

### 🔁 Infinite Loop (Use with Caution)
```go
for {
	PrintString("Looping...\n")
	break // or use return to exit
}
```

---

## 🧩 Part 4: Best Practices in the Piscine

| Tip | Why It Matters |
|-----|----------------|
| Use helper functions like `PrintInt` and `PrintString` | Reduces repetition and keeps code clean |
| Always end lines with `z01.PrintRune('\n')` | Ensures readable output |
| Avoid `fmt`, `strconv`, or `strings` | Not allowed in checkpoints |
| Use `range` for slices and strings | Idiomatic and efficient |

---

## 🧩 Part 5: Mini Piscine Exercises

### 🧪 Print Even Numbers from 0 to 20
```go
for i := 0; i <= 20; i += 2 {
	PrintInt(i)
	z01.PrintRune('\n')
}
```

### 🧪 Countdown from 10 to 0
```go
for i := 10; i >= 0; i-- {
	PrintInt(i)
	z01.PrintRune('\n')
}
```

### 🧪 Print “Go rocks!” 5 Times
```go
func PrintString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

for i := 0; i < 5; i++ {
	PrintString("Go rocks!\n")
}
```

---

## 🧩 Part 6: Advanced Notes for Piscine Survivors

- You must build your own formatting logic — even converting numbers to characters requires rune math:
  ```go
  z01.PrintRune(rune(score/10 + '0')) // crude way to print digits
  ```
- You can’t use `strconv.Itoa` — build your own number-to-string converter
- You’ll learn to appreciate every character you print — because you’ll write them all manually

---

## 🧩 Summary Table

| Component      | Role                        | Example        |
|----------------|-----------------------------|----------------|
| Initializer    | Start loop variable          | `i := 0`       |
| Condition      | Continue while true          | `i < 10`       |
| Post statement | Update after each iteration  | `i++`          |
| Output         | Manual printing              | `PrintInt(i)` + `PrintRune('\n')` |

---

This is how we build piscine fluency, Jasper — not just by looping, but by mastering control flow under constraint. You’re not just writing Go — you’re engineering it from the ground up. Ready to explore nested loops, labeled breaks, or slice iteration next? Let’s keep building your bulletproof guide.
