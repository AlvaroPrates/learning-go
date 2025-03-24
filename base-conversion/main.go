package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please, provide only the number to be converted.")
		os.Exit(1)
	}

	arg := os.Args[1]
	base := 10

	if strings.HasPrefix(arg, "0b") || strings.HasPrefix(arg, "0B") {
		base, arg = 2, arg[2:]
	} else if strings.HasPrefix(arg, "0o") || strings.HasPrefix(arg, "0O") {
		base, arg = 8, arg[2:]
	} else if strings.HasPrefix(arg, "0x") || strings.HasPrefix(arg, "0X") {
		base, arg = 16, arg[2:]
	}

	n, err := strconv.ParseInt(arg, base, 64)
	if err != nil {
		fmt.Printf("%q is not a valid number.\nUse '0b' prefix when entering a binary,'0o' for octal and '0x' for hexadecimal.\n", arg)
		os.Exit(1)
	}

	if base != 2 {
		fmt.Printf("Binary: %b\n", n)
	}
	if base != 8 {
		fmt.Printf("Octal: %o\n", n)
	}
	if base != 10 {
		fmt.Printf("Decimal: %d\n", n)
	}
	if base != 16 {
		fmt.Printf("Hexa: %X\n", n)
	}
}
