# Base Conversion

This is a simple number converter written in Go. The program takes a number in one base (binary, octal, decimal, or hexadecimal) and converts it to the other bases.

## How to Use

To use the converter, provide the number with the appropriate prefix as a command-line argument:

- **Decimal:** No prefix
- **Binary:** Start with `0b` or `0B` (e.g., `0b1101`)
- **Octal:** Start with `0o` or `0O` (e.g., `0o15`)
- **Hexadecimal:** Start with `0x` or `0X` (e.g., `0xD`)

### Example:

```sh
go run main.go 0b1101
```

This will display the following output:

```sh
Octal: 15
Decimal: 13
Hexa: D
```

# Learning Go Projects

This project is a part of the [Learning Go Projects](https://github.com/AlvaroPrates/learning-go) series.

## Learning Go - An Idiomatic Approach to Real-World Go Programming

This project was written to graps concepts about integer literals taught in the 'Predeclared Types and Declarations' chapter in the Learning Go book.

### Integer literal

In Go an integer literal is a sequence of numbers that are base 10 by default. It is possible to represent other bases using prefixes.

```go
var decimal int = 11
var binary int = 0b1101 // 13 decimal | 0b or 0B as a prefix denotes a binary number
var octal int = 0o1101 // 15 octal | 0o or 0O as a prefix denotes an octal number
var hexa int = 0x1101 // D hexadecimal | 0x or 0X as a prefix denotes a hexadecimal number
```
