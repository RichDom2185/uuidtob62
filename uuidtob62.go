package main

import (
	"fmt"
	"io"
	"math/big"
	"os"
	"strings"

	termutil "github.com/andrew-d/go-termutil"
)

func parse(uuid string) {
	b16, ok := new(big.Int).SetString(uuid, 16)
	if !ok {
		fmt.Println("Error parsing UUID: Invalid format")
		return
	}
	fmt.Println(b16.Text(62))
}

func readFromFile() []byte {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		panic(err)
	}
	return data
}

func readFromSTDIN() []byte {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Error reading input:", err)
		panic(err)
	}
	return data
}

func main() {
	isSTDINAvailable := !termutil.Isatty(os.Stdin.Fd())
	isArgumentAvailable := len(os.Args) == 2

	if !isSTDINAvailable && !isArgumentAvailable {
		fmt.Println("Incorrect usage: Either pass in a UUID string, or provide a file name.")
		return
	}

	var data []byte
	if isSTDINAvailable {
		data = readFromSTDIN()
	}
	if isArgumentAvailable {
		data = readFromFile()
	}

	original := strings.TrimSpace(string(data))
	original = strings.ReplaceAll(original, "-", "")
	parse(original)
}
