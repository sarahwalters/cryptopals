package challenge9

import (
	"fmt"
	"log"
)

// https://www.rfc-editor.org/rfc/rfc2315#:~:text=Some%20content%2Dencryption%20algorithms%20assume
func padPKCS7(block string, wantBlockSize int) (string, error) {
	padding := wantBlockSize - len(block)
	if padding < 0 {
		return "", fmt.Errorf("len(block) must be smaller than wantBlockSize")
	}

	result := []byte(block)
	for i := 0; i < padding; i++ {
		result = append(result, byte(padding))
	}

	return string(result), nil
}

func Run() {
	input := "YELLOW SUBMARINE"
	blockSize := 20
	want := "YELLOW SUBMARINE\x04\x04\x04\x04"

	got, err := padPKCS7(input, blockSize)
	switch {
	case err != nil:
		log.Fatal(err)
	case got == want:
		fmt.Printf("got expected result: %q\n", got)
	default:
		fmt.Printf("got unexpected result: %q\n", got)
	}
}
