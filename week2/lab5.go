package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// To binary

// https://stackoverflow.com/questions/37349071/golang-how-to-convert-string-to-binary-representation
func stringToBin(s string) (binString string) {
	for _, c := range s {
		binString = fmt.Sprintf("%s%b", binString, c)
	}
	return

}

// to hex
// https://www.admfactory.com/convert-string-to-hex-in-golang/
func stringToHex(s string) (hexString string) {
	hexString = hex.EncodeToString([]byte(s))
	return hexString
}

// to base64
// https://gobyexample.com/base64-encoding
func stringToBase64(s string) (hexString string) {
	b64String := base64.StdEncoding.EncodeToString([]byte(s))
	return b64String
}

func main() {
	var strings string

	fmt.Print("Enter your string: ")
	fmt.Scan(&strings)

	fmt.Println(stringToBin(strings))
	fmt.Println(stringToHex(strings))
	fmt.Println(stringToBase64(strings))
}
