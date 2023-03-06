package challenge2

import (
	"encoding/hex"
	"errors"
	"fmt"
	"log"
)

func fixedXOR(hex1, hex2 string) (string, error) {
	bytes1, err := hex.DecodeString(hex1)
	if err != nil {
		return "", err
	}
	bytes2, err := hex.DecodeString(hex2)
	if err != nil {
		return "", err
	}

	if len(bytes1) != len(bytes2) {
		return "", errors.New("inputs must be the same number of bytes")
	}

	var byte2 byte
	xor := make([]byte, 0, len(bytes1))
	for i, byte1 := range bytes1 {
		byte2 = bytes2[i]
		xor = append(xor, byte1^byte2)
	}

	fmt.Printf("bytes1: %q\n", bytes1) // "\x1c\x01\x11\x00\x1f\x01\x01\x00\x06\x1a\x02KSSP\t\x18\x1c"
	fmt.Printf("bytes2: %q\n", bytes2) // "hit the bull's eye"
	fmt.Printf("xor: %q\n", xor)       // "the kid don't play"

	return hex.EncodeToString(xor), nil
}

func Run() {
	inputHex1 := "1c0111001f010100061a024b53535009181c"
	inputHex2 := "686974207468652062756c6c277320657965"
	wantHex := "746865206b696420646f6e277420706c6179"

	gotHex, err := fixedXOR(inputHex1, inputHex2)
	switch {
	case err != nil:
		log.Fatal(err)
	case gotHex == wantHex:
		fmt.Printf("got expected result: %q\n", gotHex)
	default:
		fmt.Printf("got unexpected result: %q\n", gotHex)
	}
}
