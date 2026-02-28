# `json_parser` - [coding challenge #2] JSON Parser

## 🧐 About Repo

### In this repo, I build my own version of a JSON Parser (Uses a Lexer and Parser to try and Parse a given JSON file into something which is either valid or invalid)

Part of this Coding challenge: https://codingchallenges.fyi/challenges/challenge-json-parser

**Written in:** Go 🔵

## 👨‍🏫 Set-up & Build

1. Clone the repo:

   ```terminal
   $ git clone https://github.com/leekli/coding-challenges
   ```

2. Change directory:

   ```terminal
   $ cd json-parser
   ```

3. Ensure at least Go **1.22.5** is installed.

4. Build the Go Binary:

   ```terminal
   $ go build json_parser.go
   ```

5. There are various test files in the `testdata/` folder to test valid and invalid JSON files against - alternatively you can supply your own JSON file.

> [!NOTE]
> There are numerous test files available and can be ran using `go test ./... -run Test -v`, run this from the root of the json-parser/ directory

## 💻 How to run

Simply execute the program and give it a JSON file, it will print a message at the end either stating that the given JSON file was valid or invalid JSON. You are free to supply your own JSON file.

```terminal
$ ./json_parser testdata/step4/valid.json
```

<br></br>
