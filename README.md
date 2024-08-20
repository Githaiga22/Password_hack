# Password Hack

  ## Overview

Password Hack is a Go program that extracts and calculates the sum of two-digit numbers from a given input string. The program is designed to parse through the input string, identify substrings containing digits, and extract the first and last digits of each substring to form a two-digit number. The sum of these two-digit numbers is then calculated and returned as the output.

##  How it Works
The program uses the strings package to split the input string into individual words. Then, it iterates over each word, checking each character to see if it's a digit. If a digit is found, it's converted to an integer and stored in either the first or last variable, depending on whether it's the first or last digit encountered in the word. If both first and last are non-zero, the program adds the two-digit number formed by these digits to a running sum.

**Code Explanation**
The pass function takes a string input and returns the sum of all two-digit numbers formed by the first and last digits of each word in the input string. The main function calls pass with a sample input string and prints the result.

***Here's a breakdown of the code:***

- The main function is the entry point of the program. It calls the pass function with a sample input string and prints the result.
- The pass function takes a string input and returns an integer value.
- The function uses the strings.Split function to split the input string into individual words.
- It then iterates over each word, checking each character to see if it's a digit using the condition r >= '0' && r <= '9'.
- If a digit is found, it's converted to an integer using the expression int(r - '0').
- The first and last digits of each word are stored in the first and last variables, respectively.
- If both first and last are non-zero, the program adds the two-digit number formed by these digits to a running sum using the expression sum += first*10 + last.
- Finally, the function returns the calculated sum.

## Getting Started

### Pre-requisites

- Go installed on your system.
- A working Go environment.

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/Githaiga22/Password_hack.git
2. Navigate to the project directory:
```bash
cd Password_hack
```
3. Run the Program
```bash
 go run main.go
```
### Example Output
Running the program with the sample input string will output the sum of all two-digit numbers formed by the first and last digits of each word in the input string.

input string;
```bash
"1abc2 pqr3stu8vwx a1b2c3d4e5f treb7uchet"
```
output;
```bash
142
```
### License
This program is licensed under the MIT License. See the LICENSE file for details.

### Contributing
Contributions are welcome! If you'd like to contribute to this project, please fork the repository and submit a pull request.

### Acknowledgments
Thanks to the Go programming language and the strings package for making this program possible!

## Authors

- allkamau


## Github Profiles

- GitHub Profile: [Githaiga22](https://github.com/Githaiga22)

