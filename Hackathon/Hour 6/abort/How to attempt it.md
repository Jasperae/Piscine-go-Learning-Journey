**theory and logic** behind this `Abort` function step by step, like a true computer scientist 👇

---

## 🧠 Step 1: What does the problem ask?

> Write a function that returns the **median** of five integers.

Let’s unpack that.

### What is a **median**?

In a set of numbers, the **median** is the “middle” value **after sorting** them in order (ascending or descending).

Examples:

* For `[2, 3, 5, 7, 8]` → median = `5` (the third value).
* For `[9, 1, 4, 2, 7]` → sorted = `[1, 2, 4, 7, 9]`, median = `4`.

So our task is:

> Take 5 numbers → sort them → return the middle (3rd smallest) one.

---

## 🧩 Step 2: The naive idea

You could do:

1. Put all five numbers in an array.
2. Sort the array.
3. Return the 3rd element.

That’s simple, right?

In Go:

```go
nums := []int{a, b, c, d, e}
sort.Ints(nums)
return nums[2]
```

✅ That would work perfectly.

❌ But in this kind of challenge (like your piscine quests), you’re often **not allowed to use packages like `sort`** — so you must do it *manually* using logic and comparisons.

---

## ⚙️ Step 3: How sorting by comparisons works

If you can only compare two values at a time (`>` or `<`), you can rearrange variables so that smaller numbers move toward the front.

Example:
You can swap two variables if they’re in the wrong order:

```go
if a > b {
	a, b = b, a
}
```

This ensures that `a` is always ≤ `b`.

That’s the **building block** of sorting.

---

## 🪜 Step 4: The pattern in your code

Look at your function again:

```go
if a > b { a, b = b, a }
if a > c { a, c = c, a }
if a > d { a, d = d, a }
if a > e { a, e = e, a }
```

👉 After these 4 comparisons:

* `a` is guaranteed to be the **smallest** value among all five.
  Because you keep swapping whenever `a` is greater than any of the others.

So now:

```
a ≤ b, a ≤ c, a ≤ d, a ≤ e
```

✅ `a` is the **minimum**.

---

Next:

```go
if b > c { b, c = c, b }
if b > d { b, d = d, b }
if b > e { b, e = e, b }
```

Now:

* `b` is guaranteed to be the **second smallest** value.

---

Then:

```go
if c > d { c, d = d, c }
if c > e { c, e = e, c }
```

Now:

* `c` becomes the **third smallest**.

---

Finally:

```go
if d > e { d, e = e, d }
```

Now:

* The variables are in **sorted order**:

  ```
  a ≤ b ≤ c ≤ d ≤ e
  ```

---

## 🎯 Step 5: Why return `c`?

Because for **five numbers**, the median is the **3rd one** (the middle position).

Positions:

```
1st → a
2nd → b
3rd → c ← median
4th → d
5th → e
```

✅ So returning `c` gives you the **median value**.

---

## 🧩 Step 6: What’s really happening — in words

You can describe the algorithm like this:

1. Start with five integers.
2. Compare pairs and **swap** them so that smaller ones move left.
3. Do this systematically to “bubble” the smallest to the left and the largest to the right.
4. Once fully sorted, return the 3rd number — the middle one.

This is a **manual sorting algorithm**, very similar to **selection sort**, but using direct variable comparisons instead of loops or arrays.

---

## ⚡ Step 7: Why not loops?

Good question!

You *could* write a shorter version with arrays and loops, like:

```go
arr := []int{a, b, c, d, e}
for i := 0; i < len(arr); i++ {
    for j := i + 1; j < len(arr); j++ {
        if arr[i] > arr[j] {
            arr[i], arr[j] = arr[j], arr[i]
        }
    }
}
return arr[2]
```

But the piscine often **disallows slices, arrays, or external packages**,
so you must simulate sorting with only variables — like in your solution.

---

## 🧠 Summary of the Theory

| Concept                 | Explanation                                |
| ----------------------- | ------------------------------------------ |
| **Median**              | The middle value after sorting the numbers |
| **Comparison**          | Check two values to determine order        |
| **Swap**                | Exchange two values to correct their order |
| **Manual sorting**      | Sorting using only `if` and variable swaps |
| **Median for 5 values** | Always the 3rd smallest one                |

---

## 💡 Analogy

Imagine 5 people standing in random order by height.
You go through them, swapping whenever a taller one stands before a shorter one.
After a few swaps, everyone is arranged shortest → tallest.
Now, the one standing in the middle is the **median height** — that’s `c`.

---

Would you like me to show you a **visual diagram** (with arrows showing swaps) for how this sorting happens? It’s a fun way to actually *see* how the numbers move into order.
