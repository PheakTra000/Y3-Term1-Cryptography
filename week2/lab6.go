package main

import (
	"fmt"
)

// https://github.com/TheAlgorithms/Go/blob/master/cipher/xor/xor.go

func xorEncrypt(text string, key string) (output string) {
	for i := 0; i < len(text); i++ {
		output += string(text[i] ^ key[i%len(key)])
	}
	return output
}

func main() {
	var text string
	var key string
	fmt.Print("Enter your text: ")
	fmt.Scan(&text)
	fmt.Print("Enter your key: ")
	fmt.Scan(&key)

	fmt.Println(xorEncrypt(text, key))

}
