Let’s break this down so you understand not just *what* this script does, but *why* every part of it works — and how you could build it yourself from scratch.

---

## 🎯 Goal Recap

Create a script `who-are-you.sh` that:
- Fetches superhero data from a JSON file online
- Filters for the superhero with `id: 70`
- Prints only their name
- Keeps the output exactly formatted like:
  ```
  "name"$
  $
  ```

---

## 🧠 Step-by-Step Theory

### 1. **Shebang: `#!/bin/bash`**
This line tells your system:
> “Use the Bash shell to interpret this script.”

It ensures consistent behavior across systems and shells.

---

### 2. **Fetching JSON with `curl`**
```bash
curl -s https://acad.learn2earn.ng/assets/superhero/all.json
```

- `curl` is a command-line tool to fetch data from URLs
- `-s` means “silent” — it suppresses progress and error messages
- This downloads the JSON array of superhero objects

> ✅ You now have the raw JSON data in your terminal

---

### 3. **Parsing JSON with `jq`**
```bash
jq '.[] | select(.id == 70) | .name'
```

Let’s break this down:
- `.[]`: iterates over each object in the array
- `select(.id == 70)`: filters for the object with `id` equal to `70`
- `.name`: extracts the `name` field from that object

> ✅ This gives you the name of the superhero, still wrapped in quotes (as required)

---

### 4. **Putting It All Together**

Here’s the full script:

```bash
#!/bin/bash
curl -s https://acad.learn2earn.ng/assets/superhero/all.json | jq '.[] | select(.id == 70) | .name'
```

- This fetches the data
- Filters for the superhero with ID 70
- Extracts and prints the name

---

### 5. **Verifying the Output Format**

When you run:
```bash
./who-are-you.sh | cat -e
```

You should see:
```
"name"$
$
```

- `cat -e` shows `$` at the end of each line
- The second `$` is from the empty line after the output

> ✅ This confirms the output is exactly as required

---

## 🛠️ Summary of Key Concepts

| Component | Purpose |
|----------|---------|
| `#!/bin/bash` | Use Bash to run the script |
| `curl -s` | Fetch JSON data silently |
| `jq` | Parse and filter JSON |
| `.[]` | Iterate over array items |
| `select(.id == 70)` | Filter for superhero with ID 70 |
| `.name` | Extract the name field |
| `cat -e` | Show end-of-line characters for verification |

---

Want to go deeper? I can show you how to extract multiple fields, handle missing IDs, or format the output differently. Just say the word!

Shell pipelines are a powerful feature in Unix-like systems (like Linux and macOS) that let you **chain multiple commands together**, passing the output of one command directly as the input to the next — like an assembly line for data.

---

## 🧪 Basic Syntax

```bash
command1 | command2 | command3
```

- The `|` symbol is called a **pipe**.
- It connects the **standard output** (`stdout`) of `command1` to the **standard input** (`stdin`) of `command2`.

---

## 🧠 Real-World Example

```bash
cat file.txt | grep "error" | sort | uniq -c
```

### What this does:
1. `cat file.txt` — prints the contents of `file.txt`
2. `grep "error"` — filters lines that contain the word "error"
3. `sort` — sorts those lines alphabetically
4. `uniq -c` — counts how many times each unique line appears

➡️ You get a frequency count of each unique error line in the file.

---

## ✅ Why Use Pipelines?

- **Efficiency**: Avoids creating temporary files
- **Modularity**: Each command does one job well
- **Power**: You can build complex data processing flows with simple tools

---

## 🔧 Common Pipeline Tools

| Command   | What it does                     |
|-----------|----------------------------------|
| `grep`    | Filters lines by pattern         |
| `awk`     | Pattern scanning and processing  |
| `sed`     | Stream editing                   |
| `cut`     | Extracts columns or fields       |
| `sort`    | Sorts lines                      |
| `uniq`    | Removes duplicates               |
| `wc`      | Counts lines, words, characters  |
| `jq`      | Parses and filters JSON          |

---

Let me know if you want to build a pipeline for a specific task — like filtering logs, processing JSON, or analyzing CSV files. I’d love to help you craft one.
