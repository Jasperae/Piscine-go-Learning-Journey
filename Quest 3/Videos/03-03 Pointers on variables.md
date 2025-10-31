Absolutely brilliant, Jasper — and now let’s checkpoint-proof this masterclass. Since piscine constraints don’t allow fmt, we’ll rewrite everything using only:

- ✅ package main
- ✅ github.com/01-edu/z01 for output
- ❌ No fmt, strings, or other standard packages

---

🧠 Masterclass: Pointers in Go — Memory, Mutation & Mastery (Checkpoint Edition)

🧩 Part 1: What Is a Pointer?

| Concept     | Meaning |
|-------------|--------|
| Value       | Actual data stored in a variable |
| Address     | Memory location of that value |
| Pointer     | Variable that stores an address |
| Dereference | Access or modify value at that address |

Pointers let you change a value without copying it — perfect for in-place mutation.

---

🧩 Part 2: Declare and Use Pointers

✅ Example
`go
package main

import "github.com/01-edu/z01"

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

func main() {
	a := 10
	p := &a
	PrintInt(*p)      // prints 10
	z01.PrintRune('\n')
	*p = 42
	PrintInt(*p)      // prints 42
	z01.PrintRune('\n')
}
`

---

🧩 Part 3: Value vs Reference

✅ Passing by Value
`go
func addTen(x int) int {
	return x + 10
}
`
➡ Copies the value — original stays unchanged.

✅ Passing by Pointer
`go
func addTen(x *int) {
	*x += 10
}
`
➡ Modifies the original value directly.

---

🧩 Part 4: Mutation in Action

`go
func addTen(x *int) {
	*x += 10
}

func main() {
	a := 10
	addTen(&a)
	PrintInt(a) // prints 20
	z01.PrintRune('\n')
}
`

---

🧩 Part 5: Common Pitfalls

| Mistake                     | Fix                        |
|----------------------------|----------------------------|
| Using *int without &val| Always assign with &a    |
| Forgetting  to modify   | Use p = newValue        |
| Mixing int and *int    | Match types in functions   |

---

🧩 Part 6: Mini Exercises

🔹 Print address and value (simplified)
`go
func main() {
	x := 42
	p := &x
	PrintInt(*p) // prints 42
	z01.PrintRune('\n')
}
`

🔹 Modify via pointer
`go
func main() {
	x := 5
	p := &x
	*p = 100
	PrintInt(x) // prints 100
	z01.PrintRune('\n')
}
`

🔹 Double a value
`go
func double(n *int) {
	n = 2
}
`

---

🧩 Summary Table

| Concept     | Code Example     | Result        |
|-------------|------------------|---------------|
| Declare int | a := 10        | Value: 10     |
| Pointer     | p := &a        | Address of a|
| Dereference | *p             | Value at p  |
| Modify      | *p += 5        | a becomes 15|
| Function    | func f(x *int) | Mutates value |

---

This is how you master pointers, Jasper — not just by using them, but by understanding how they shape memory and control flow. Want to build a pointer-based swap function next? Let’s keep going.