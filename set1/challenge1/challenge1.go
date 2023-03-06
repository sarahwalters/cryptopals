package challenge1

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

func hexToBase64(hexStr string) (string, error) {
	// In hex, 2 characters represent 1 byte (8 bits)
	// Convert from hex to bytes
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", err
	}

	fmt.Printf("bytes: %q\n", bytes) // "I'm killing your brain like a poisonous mushroom"

	// In base64, 4 characters represent 3 bytes (each character represents 6 bits)
	// Mapping here: https://medium.com/swlh/powering-the-internet-with-base64-d823ec5df747
	// Convert from bits to base64
	return base64.StdEncoding.EncodeToString(bytes), nil
}

func Run() {
	inputHex := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	wantBase64 := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	gotBase64, err := hexToBase64(inputHex)
	switch {
	case err != nil:
		log.Fatal(err)
	case gotBase64 == wantBase64:
		fmt.Printf("got expected result: %q\n", gotBase64)
	default:
		fmt.Printf("got unexpected result: %q\n", gotBase64)
	}
}
