package main

import (
	"flag"
	"fmt"
	"strings"
)

var table = map[string]string{
	"a": "n", "b": "o", "c": "p", "d": "q",
	"e": "r", "f": "s", "g": "t", "h": "u",
	"i": "v", "j": "w", "k": "x", "l": "y",
	"m": "z", "n": "a", "o": "b", "p": "c",
	"q": "d", "r": "e", "s": "f", "t": "g",
	"u": "h", "v": "i", "w": "j", "x": "k",
	"y": "l", "z": "m", " ": " ", ".": ".",
	",": ",", "!": "!", "?": "?", ";": ";",
	":": ":", "\"": "\"", "'": "'", "(": "(",
	")": ")", "[": "[", "]": "]", "{": "{",
	"}": "}", "0": "5", "1": "6", "2": "7",
	"3": "8", "4": "9", "5": "0", "6": "1",
	"7": "2", "8": "3", "9": "4", "&": "&",
	"@": "@", "#": "#", "$": "$", "%": "%",
	"^": "^", "*": "*", "=": "=", "+": "+",
	"-": "-", "_": "_", "~": "~", "`": "`",
	"|": "|", "/": "/", "\\": "\\",
}

var msg = `
Hello, this is a simple implementation of the Caesar cipher algorithm.
only supports the lowercase alphabet and numbers.

Usage:
	$ go run main.go -str="Hello, this is a simple implementation of the Caesar cipher algorithm."
`

func main() {
	str := flag.String("str", "", "string to cipher")
	help := flag.Bool("h", false, "help")
	flag.Parse()

	if *help || *str == "" {
		fmt.Println(msg)
		return
	}

	fmt.Println(caesarCipher(*str))

}

func caesarCipher(str string) string {
	str = strings.ToLower(str)
	var result string
	for _, v := range str {
		result += table[string(v)]
	}
	return result
}
