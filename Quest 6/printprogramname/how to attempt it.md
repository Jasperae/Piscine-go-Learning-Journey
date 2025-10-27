Absolutely! Let’s break down this `printprogramname` program line by line and uncover the **why** behind every part. By the end, you’ll understand how it works from the inside out—and be able to rebuild or adapt it confidently.

---

## 🧠 Goal of the Program

We want to write a Go program that:

- Prints the **name of the executable file** that was run
- For example, if you run `./Nessy`, it should print `Nessy`
- If you run `./printprogramname`, it should print `printprogramname`

---

## ✅ Full Code

```go
package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	r := os.Args[0]
	var k int

	for i := 0; i < len(r); i++ {
		if r[len(r)-1-i] == '/' {
			k = len(r) - 1 - i
			break
		}
	}
	s := r[k+1:]
	for _, i := range s {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
```

---

## 🔍 Step-by-Step Breakdown

### 1. `package main`

- Every Go program starts with `package main` when it’s meant to be an executable.
- This tells the compiler: “This is the entry point.”

---

### 2. `import ( "os" "github.com/01-edu/z01" )`

- `os`: Gives access to `os.Args`, which holds command-line arguments.
- `z01`: A special package used in 01-edu projects to print characters using `z01.PrintRune`.

---

### 3. `r := os.Args[0]`

- `os.Args[0]` is the **full path** to the program that was executed.
- Example:
  - If you run `./Nessy`, `os.Args[0]` might be `"./Nessy"`
  - If you run `student/piscine/printprogramname/main`, it might be `"student/piscine/printprogramname/main"`

---

### 4. `var k int`

- We’ll use `k` to store the index of the last `/` character in the path.
- This helps us isolate the actual program name from the full path.

---

### 5. `for i := 0; i < len(r); i++ { ... }`

- Loops through the string `r` from left to right.
- But inside the loop, we check characters from **right to left** using `r[len(r)-1-i]`.

---

### 6. `if r[len(r)-1-i] == '/' { ... }`

- Checks if the character is a `/`, which separates directories in a path.
- When we find the **last** `/`, we know the program name starts right after it.

---

### 7. `k = len(r) - 1 - i`

- Stores the index of the last `/` found.
- We’ll use this to slice the string and extract the program name.

---

### 8. `s := r[k+1:]`

- Slices the string from the character **after** the last `/` to the end.
- This gives us just the program name.

---

### 9. `for _, i := range s { z01.PrintRune(i) }`

- Loops through each character in the program name.
- Prints each character using `z01.PrintRune`.

---

### 10. `z01.PrintRune('\n')`

- Prints a newline at the end for clean output.

---

## 🧪 Example Execution

```bash
$ go build -o Nessy
$ ./Nessy
Nessy

$ go build
$ ./printprogramname
printprogramname
```

---

## ✅ Summary

| Concept | Purpose |
|--------|---------|
| `os.Args[0]` | Gets full path of the executable |
| `r[len(r)-1-i]` | Scans path from end to find last `/` |
| `r[k+1:]` | Extracts program name |
| `z01.PrintRune` | Prints each character individually |
| `'\n'` | Adds a line break after the name |

---

Want to explore how to print the full path, or how to detect if the program was run from a symbolic link? I’d love to walk you through that next.
