**To extract the superhero with `id: 170` and display their `name`, `power`, and `gender`, you’ll use `curl` to fetch the JSON and `jq` to parse it. The key is understanding how to filter and format JSON data using command-line tools.**

---

## 🧠 Step-by-Step Theory Behind the Solution

### 🎯 Goal Recap
Create a script `to-git-or-not-to-git.sh` that:
- Fetches superhero data from a JSON file online
- Filters for the superhero with `id: 170`
- Extracts and prints:
  - `name`
  - `power`
  - `gender`
- Each on a separate line, with no extra characters

---

### 1. **Fetching JSON with `curl`**
```bash
curl -s https://acad.learn2earn.ng/assets/superhero/all.json
```

- `curl` is used to download data from a URL
- `-s` means “silent” — it suppresses progress and error messages
- This fetches the JSON array of superhero objects

---

### 2. **Parsing JSON with `jq`**
```bash
jq -r '.[] | select(.id == 170) | .name, .power, .appearance.gender'
```

Let’s break this down:
- `.[]`: iterates over each object in the array
- `select(.id == 170)`: filters for the object with `id` equal to `170`
- `.name, .power, .appearance.gender`: extracts the three fields we want
- `-r`: outputs raw strings (removes quotes)

> ✅ This gives us exactly:
```
Chameleon
28
Male
```

---

### 3. **Putting It All Together**

Here’s the full script:

```bash
#!/bin/bash
curl -s https://acad.learn2earn.ng/assets/superhero/all.json \
  | jq -r '.[] | select(.id == 170) | .name, .power, .appearance.gender'
```

- This fetches the data
- Filters for the superhero with ID 170
- Extracts and prints the required fields

---

## 🧪 Example Output

Running:
```bash
bash to-git-or-not-to-git.sh
```

Will output:
```
Chameleon
28
Male
```

Exactly as required.

---

## 🛠️ Summary of Key Concepts

| Component | Purpose |
|----------|---------|
| `#!/bin/bash` | Use Bash to run the script |
| `curl -s` | Fetch JSON data silently |
| `jq -r` | Parse JSON and remove quotes |
| `.[]` | Iterate over array items |
| `select(.id == 170)` | Filter for superhero with ID 170 |
| `.name, .power, .appearance.gender` | Extract required fields |

---

Want to go deeper? I can show you how to format the output as JSON, handle missing fields, or even loop through multiple IDs. Just say the word!
