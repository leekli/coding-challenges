# `wc` - [coding challenge #1] word count

## 🧐 About Repo

### In this repo, I build my own version of the Unix command line tool: `wc`

Part of this Coding challenge: https://codingchallenges.fyi/challenges/challenge-wc

**Written in:** Go 🔵

## 👨‍🏫 Set-up & Build

1. Clone the repo:

   ```terminal
   $ git clone https://github.com/leekli/coding-challenges 
   ```

2. Change directory:

   ```terminal
   $ cd wc
   ```

3. Ensure at least Go **1.22.5** is installed.


4. Build the Go Binary:

   ```terminal
   $ go build wc.go
   ```

5. A test file `test.txt` is supplied to use as an example.

## 💻 Commands & Options available

- **Help:** Shows the help message, explanation of the application, and options available.

  ```terminal
  $ ./wc -h
  ```

> [!NOTE]
> All commands below can be used either by giving a file name, or piping through standard input (stdin).

- **-c [file_name | stdin]:** Displays the number of bytes in the input file / standard input given

  - With a file name:

  ```terminal
  $ ./wc -c test.txt
  ```

  - With standard input:

  ```terminal
  $ cat test.txt | ./wc -c
  ```

<br></br>

- **-l [file_name | stdin]:** Displays the number of lines in the input file / standard input given

  - With a file name:

  ```terminal
  $ ./wc -l test.txt
  ```

  - With standard input:

  ```terminal
  $ cat test.txt | ./wc -l
  ```

  <br></br>

- **-w [file_name | stdin]:** Displays the number of words in the input file / standard input given

  - With a file name:

  ```terminal
  $ ./wc -w test.txt
  ```

  - With standard input:

  ```terminal
  $ cat test.txt | ./wc -w
  ```

  <br></br>

- **-m [file_name | stdin]:** Displays the number of characters in the input file / standard input given

  - With a file name:

  ```terminal
  $ ./wc -m test.txt
  ```

  - With standard input:

  ```terminal
  $ cat test.txt | ./wc -m
  ```

<br></br>

- **[file_name | stdin]:** No flags given, only a file name. Displays the number of lines, bytes and characters in the input file given, or piping/using standard input.

  - With a file name:

  ```terminal
  $ ./wc test.txt
  ```

  - Using standard input and no flags or file name given:

  ```terminal
  $ cat test.txt | ./wc
  ```