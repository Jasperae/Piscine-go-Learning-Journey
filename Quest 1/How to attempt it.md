https://www.youtube.com/watch?list=PLHyAJ_GrRtf8MgH469gNRWFV-wPW_gLue&v=A0Mqc215igw

`curl` - **Client Web Address** is a command-line tool used for transferring data between a client and a server. It supports multiple protocols like HTTP, HTTPS, FTP, and more. The tool is often used for downloading or uploading files, interacting with web services (APIs), and performing network debugging.


**FIRST CREATE A REPO**
My repository
1. Under Settings, click on Applications on the left panel.
2. Select important Features to be included

1. I created a file on the VS code application
2. the three dots up, I created a new terminal
3. in the terminal, i typed - git clone https://acad.learn2earn.ng/git/jaigwe/piscine-go.git
4. Then I filled in my details accordingly - so with that I was able to clone into an empty repository
5. Still at that 3 dots dropdown, click on open and you will see you saved file there.

```
git add hello.sh
git commit -m "My very first commit"
git push
```
```
$ git config --global credential.helper store
$ git clone https://acad.learn2earn.ng/git/(Your UserName)/piscine-go.git
```
```
$ rm -rf ~/piscine-go
$ ls piscine-go -a
.  ..  .git
```
```
$ touch hello.sh
$ code hello.sh
$ echo "Hello Your UserName!"
Hello jaigwe!
$ bash hello.sh
$ cat hello.sh
```
```
$ git add [hello.sh](http://hello.sh/)

$ git commit -m "My very first commit"
[main (root-commit) 0dde791] My very first commit
1 file changed, 0 insertions(+), 0 deletions(-)
create mode 100644 [hello.sh](http://hello.sh/)
```

```
$ git push
Enumerating objects: 3, done.
Counting objects: 100% (3/3), done.
Writing objects: 100% (3/3), 217 bytes | 217.00 KiB/s, done.
Total 3 (delta 0), reused 0 (delta 0), pack-reused 0
remote: fatal: ambiguous argument '0000000000000000000000000000000000000000 0dde791d9292069df40875b712a3788a10bc8a61 refs/heads/main..': unknown revision or path not in the working tree.
remote: Use '--' to separate paths from revisions, like this:
remote: 'git <command> [<revision>...] -- [<file>...]'
remote: . Processing 1 references
remote: Processed 1 references in total
$
```

## My Workflow

```
cd ~
git clone https://acad.learn2earn.ng/git/jaigwe/piscine-go.git
cd piscine-go
code .
```

`git config --global credential.helper store`


## Fourth Video in the Quest 01 Playlist

Find is used to find files, and grep is used to find patterns in files

mkdir test

ls

cd test

ls

touch {1..100}

find 10

find -name ‘1*’  starting with 1

find -name ‘*1’   Ending with 1

find -name ‘*1*’      Contains 1

echo hello > 101  — Adds creates file 101 and adds String hello into it

cat 1 — Checks for the content of file 1

grep hello  1 — to call out what file contains string hello

grep hello * — to call out all the files that contains string hello0


## Descriptive Summary of the Video: Basic Unix Text-Processing Commands
**Intro**

This descriptive summary distills a short tutorial on three core terminal tools for string manipulation: cut, tr, and sed. The speaker demonstrates practical, bite-sized examples using an inline sentence, highlighting how each tool handles delimitation, character replacement, and pattern-based substitutions. The focus is on clarity, accessibility, and *gradual familiarity* with command syntax.

$ echo 'Hello! How are you today?' | cut -d " " -f5
today?
$ echo 'Hello! How are you today?' | cut -d " " -f2

How

$ echo 'Hello! How are you today?' | cut -d "e" -f2
llo! How ar
$ echo 'Hello! How are you today?' | cut -d "e" -f1
H
$ echo 'Hello! How are you today?' | cut -d "e" -f3
you today?
$ echo 'Hello! How are you today?' | cut -d "e" -f2

In this walkthrough, the narrator opens a terminal, echoes a sentence, and explores each tool in turn. The example sentence is simple: "Hello! How are you today?" This line is then manipulated to illustrate different outcomes without touching any files—only standard input is required.

**Descriptive sequence**

- **Cut**: The command requires a delimiter (for instance, a space). It returns:
    - Section one: the portion before the first delimiter (e.g., "Hello!").
    - Section two: the portion after the first delimiter and before the second delimiter, when relevant.
    - The tool can also split by a single character, but it cannot split using multi-character patterns.
- **Tr**: This tool performs direct character replacements:
    - Replacing all small o with uppercase O demonstrates a straightforward mapping.
    - Replacing both o and w with uppercase letters (O and A, respectively) is also possible.
    - Note: <em>tr</em> is best for single-character substitutions and is not intended for complex patterns.
- **Sed**: The most flexible among the trio, it supports pattern-based substitutions:
    - A longer sentence variant is used to illustrate making all spaces uppercase letters when combined with a global flag.
    - It can replace occurrences of words like “you” with “I,” showing pattern matching and global scope.
    - Important: a proper syntax is required, using a pair of delimiters to specify the replacement target and the substitute.

**Table: Commands at a glance**

| Command | Primary purpose | Basic example used in the demo |
| --- | --- | --- |
| cut | Split text by a delimiter | Section extraction around spaces or a chosen character |
| tr | Translate or map single characters | Uppercase conversion for chosen letters |
| sed | Stream editor for substitutions | Global and pattern-based replacements |
- The speaker emphasizes consulting the manual pages for each tool to unlock broader capabilities.
- The takeaway is that there isn’t a single method for any task; multiple approaches exist, and experience grows by trying variations and reading documentation.

**Center**

- *Practice tips*: run commands directly in the terminal, experiment with different delimiters, and extend patterns with regular expressions when using sed.
- *Encouragement*: there are many valid methods to achieve the same result, so explore and practice.

**Outro**

Thank you for listening. This concise primer introduces three foundational string-manipulation tools and demonstrates practical usage with minimal setup. With curiosity and persistence, you can master these commands and discover myriad ways to tailor text processing to your needs.

## Descriptive Summary

Hello! In this concise, feature-rich summary we unpack a brief tutorial that introduces two handy command-line tools for developers: jq for JSON manipulation and wc for word counting (WC — Word Counter). The dialogue is practical, example-driven, and designed to extend routine shell workflows with light but effective techniques. Below you’ll find an organized, easy-to-skim guide that captures the core ideas, steps, and outcomes.

**Intro**

- The speaker presents two extra commands to enhance everyday terminal tasks: **jq** (a JSON-specific filter and extractor) and **wc** (a word, line, and byte counter).
- The focus is practical usage, not exhaustive coverage, with emphasis on quick demonstrations using a real API example.

<center>

### Practical walkthrough (centered)

- The opening example uses a curl request to fetch data from a GitHub API, yielding a list of users (in this case, 46 objects are returned, not all users are shown).
- The goal is to demonstrate filtering to a single object by a condition, rather than manually scanning IDs.

</center>

- The core technique with jq relies on placing a filter expression between apostrophes. The presenter emphasizes that jq’s syntax must be exact and recommends online resources for precision.
- A representative jq workflow is shown:
    - Start with the object list.
    - Filter to a specific user by a condition (e.g., an ID corresponding to a username like "kevwil"), aiming to retrieve the object with ID 35.
    - The demonstration shows how jq returns a filtered, concise data snippet that focuses only on the chosen entry.
- A quick follow-up demonstrates how jq can isolate nested or specific keys, such as the avatar URL. This illustrates jq’s power to drill into objects and extract just the desired fields.
- After showcasing jq, the presenter clears the screen to emphasize clarity and then transitions to the next tool.
- The second tool, **wc**, is introduced as a companion for counting elements in the output from the previous command:
    - Piping the output of a curl/jq command into wc reveals three numbers: bytes, words, and lines.
    - To extract a single metric, options are used; for example, the number of lines is mapped to the short flag **l**.
    - In the shown example, the number of lines is 20.
- The speaker reinforces that these two tools are “the very basics” of their capabilities: jq for JSON manipulation and wc for counting data elements, with a possible extension to isolate specific keys or values when needed.

**Tables and quick reference**

| Tool | Purpose | Basic usage | Key note |
| --- | --- | --- | --- |
| jq | JSON manipulation and filtering | jq '...expr...' | Requires precise syntax; ideal for selecting objects by condition |
| wc | Word/byte/line counting | wc -l (lines), wc -w (words), wc -c (bytes) | Useful for quantifying output length after filters |
- The takeaway is succinct: these utilities are foundational, easy to learn, and surprisingly effective for everyday scripting tasks.

**Center**

- As the session winds down, the speaker reiterates the availability of these tools for future exercises and encourages practice to gain fluency with their syntax and options.

**Outro**

- A brief sign-off invites viewers to return for more tips and confirms that the next session will cover additional capabilities or alternative workflows.
- The tone remains friendly and practical, inviting experimentation and reinforcing that mastering jq and wc can streamline JSON handling and data measurement in shell environments.

```
$ curl https://api.github.com/users | jq ' .[] | select(.id == 35) | .avatar_url'
% Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
Dload  Upload   Total   Spent    Left  Speed
100 30885  100 30885    0     0   3212      0  0:00:09  0:00:09 --:--:--  7440
"https://avatars.githubusercontent.com/u/35?v=4"
```

```
$ curl https://api.github.com/users | jq ' .[] | select(  .id == 35   )' | wc
% Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
Dload  Upload   Total   Spent    Left  Speed
100 30885  100 30885    0     0   3359      0  0:00:09  0:00:09 --:--:--  8673
21      40     967
```

The `-e` flag in the command

cat -e filename

is **not part of `cat`’s normal behavior**, it’s a **special option** that makes invisible characters **visible**, mainly for debugging or precise verification.

Let’s break it down 👇
