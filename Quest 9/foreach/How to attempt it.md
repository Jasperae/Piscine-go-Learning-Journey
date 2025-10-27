Here’s the complete and correct solution for the `ForEach` exercise, along with a full breakdown of **why** it works and how to build it from scratch.

---

## ✅ Final Code

```go
package piscine

func ForEach(f func(int), a []int) {
	for _, v := range a {
		f(v)
	}
}
```

---

## 🧠 What’s the Goal?

You’re asked to write a function that:
- Takes a **function** `f` that operates on an `int`
- Takes a **slice of ints** `a`
- Applies `f` to **each element** of `a`

This is a classic **higher-order function** pattern — passing a function as an argument and applying it to a collection.

---

## 🧩 Step-by-Step Breakdown

### 1. **Function Signature**

```go
func ForEach(f func(int), a []int)
```

- `f func(int)`: a function that takes an `int` and returns nothing
- `a []int`: a slice of integers

### 2. **Loop Through the Slice**

```go
for _, v := range a {
	f(v)
}
```

- `range a` iterates over each element in the slice
- `v` is the value of each element
- `f(v)` applies the function to that value

---

## 🧪 Example Usage

```go
a := []int{1, 2, 3, 4, 5, 6}
piscine.ForEach(piscine.PrintNbr, a)
```

Assuming `PrintNbr` prints an integer without a newline, the output will be:

```
123456
```

---

## ✅ Summary of Concepts

| Concept             | Why It Matters                              |
|---------------------|----------------------------------------------|
| Function as argument| Enables flexible behavior                   |
| Slice iteration     | Applies logic to each element               |
| `range` loop        | Clean and idiomatic Go                      |

---

Would you like to extend this to support functions that return values, or apply it to other types like `[]string`? I can show you how to generalize it.
