package main

import (
	"fmt"
	"math/big"
	"os"
	"strings"
)

func parse(uuid string) {
	b16, ok := new(big.Int).SetString(uuid, 16)
	if !ok {
		fmt.Println("Error parsing UUID: Invalid format")
		return
	}
	fmt.Println(b16.Text(62))
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing argument: Please provide a file name")
		return
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		panic(err)
	}
	original := strings.TrimSpace(string(data))
	original = strings.ReplaceAll(original, "-", "")
	parse(original)
}
