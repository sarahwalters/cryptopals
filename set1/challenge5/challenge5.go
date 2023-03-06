package challenge5

import (
	"encoding/hex"
	"fmt"
)

// XORs the two bytes.
func xor(a, b byte) byte {
	return a ^ b
}

// XORs each byte in the input with a byte in the key. The first
// byte in the input is XOR'd with the first byte in the key, the
// second byte in the input with the next byte in the key, and so
// on, wrapping back around at the end of the key.
func repeatingKeyXOR(input, key string) string {
	var keyIndex int
	var result []byte

	for _, b := range []byte(input) {
		result = append(result, xor(b, key[keyIndex]))

		keyIndex++
		if keyIndex == len(key) {
			keyIndex = 0
		}
	}

	return hex.EncodeToString(result)
}

func Run() {
	input := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`
	key := "ICE"

	want := `0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f`

	got := repeatingKeyXOR(input, key)
	switch {
	case got == want:
		fmt.Printf("Got expected result: %q\n", got)
	default:
		fmt.Printf("Got unexpected result: %q\n", got)
	}
}
