package challenge10

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"

	"cryptopals/set1/challenge7"
)

func encryptAESWithECB(plaintext, key []byte) ([]byte, error) {
	// The crypto/aes package will choose AES-128 if the key is 16 bytes long,
	// AES-192 if the key is 24 bytes long, or AES-256 if the key is 32 bytes
	// long. See https://pkg.go.dev/crypto/aes#NewCipher.
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	encrypted := make([]byte, len(plaintext))

	// QUESTION: Does ECB block size have to be the length of the key?
	for _, b := range challenge7.Blocks(plaintext, len(key)) {
		cipher.Encrypt(encrypted[b.Start:b.End], plaintext[b.Start:b.End])
	}

	return encrypted, nil
}

// Check that if we encrypt some plaintext with a key using AES with ECB,
// then decrypt it with the same key using AES With ECB, we get the plaintext
// back.
func checkAESWithECB(plaintext, key []byte) (bool, error) {
	ciphertext, err := encryptAESWithECB(plaintext, key)
	if err != nil {
		return false, err
	}

	decrypted, err := challenge7.DecryptAESWithECB(ciphertext, key)
	if err != nil {
		return false, err
	}

	return string(decrypted) == string(plaintext), nil
}

// XORs the input slices together byte by byte.
func xor(xs, ys []byte) ([]byte, error) {
	if len(xs) != len(ys) {
		return nil, fmt.Errorf("inputs to xor must be the same length")
	}

	result := make([]byte, 0, len(xs))
	for i, x := range xs {
		y := ys[i]
		result = append(result, x^y)
	}
	return result, nil
}

// Great diagrams of encryption and decryption here:
// https://en.wikipedia.org/wiki/Block_cipher_mode_of_operation#:~:text=citation%20needed%5D-,Cipher%20block%20chaining%20(CBC),-%5Bedit%5D

// First plaintext block:
// - xor initialization vector with plaintext block
// - encode with block cipher
// - get a ciphertext
//
// Remaining plaintext blocks:
// - xor last ciphertext with plaintext block
// - pass that to the block cipher
// - get a new ciphertext
func encryptAESWithCBC(plaintext, key, iv []byte) ([]byte, error) {
	// QUESTION: Does CBC block size have to be the length of the key?
	var ciphertext, lastCiphertextBlock []byte
	for _, b := range challenge7.Blocks(plaintext, len(key)) {
		var xorWith []byte
		if len(lastCiphertextBlock) == 0 {
			xorWith = iv
		} else {
			xorWith = lastCiphertextBlock
		}

		xorBlock, err := xor(b.Bytes, xorWith)
		if err != nil {
			return nil, err
		}

		ciphertextBlock, err := encryptAESWithECB(xorBlock, key)
		if err != nil {
			return nil, err
		}
		ciphertext = append(ciphertext, ciphertextBlock...)
		lastCiphertextBlock = ciphertextBlock
	}

	return ciphertext, nil
}

// First ciphertext block:
// - decrypt with block cipher
// - xor decrypted result with initialization vector
// - get a plaintext block
//
// Remaining ciphertext blocks:
// - decrypt with block cipher
// - xor decrypted result with last ciphertext block (not plaintext!)
// - get a plaintext block
func decryptAESWithCBC(ciphertext, key, iv []byte) ([]byte, error) {
	// QUESTION: Does CBC block size have to be the length of the key?
	var plaintext, lastCiphertextBlock []byte
	for _, b := range challenge7.Blocks(ciphertext, len(key)) {
		decryptedBlock, err := challenge7.DecryptAESWithECB(b.Bytes, key)
		if err != nil {
			return nil, err
		}

		var xorWith []byte
		if len(lastCiphertextBlock) == 0 {
			xorWith = iv
		} else {
			xorWith = lastCiphertextBlock
		}

		plaintextBlock, err := xor(decryptedBlock, xorWith)
		if err != nil {
			return nil, err
		}
		plaintext = append(plaintext, plaintextBlock...)

		lastCiphertextBlock = b.Bytes
	}

	return plaintext, nil
}

func decryptFile(path string, key, iv []byte) (string, error) {
	// The data in the file is base64-encoded, even though the challenge doesn't say
	// so -- can tell because decrypting these bytes directly gives nonsense while
	// decoding to base64 and decrypting the base64 gives good results.
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	bytes, err := base64.StdEncoding.DecodeString(string(raw))
	if err != nil {
		return "", err
	}

	plaintext, err := decryptAESWithCBC(bytes, key, iv)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

// Check that if we encrypt some plaintext with a key using AES with ECB,
// then decrypt it with the same key using AES With ECB, we get the plaintext
// back.
func checkAESWithCBC(plaintext, key, iv []byte) (bool, error) {
	ciphertext, err := encryptAESWithCBC(plaintext, key, iv)
	if err != nil {
		return false, err
	}

	decrypted, err := decryptAESWithCBC(ciphertext, key, iv)
	if err != nil {
		return false, err
	}

	return string(plaintext) == string(decrypted), nil
}

func Run() {
	key := []byte("YELLOW SUBMARINE")

	// QUESTION: Assuming ECB and CBC block size is the length of the key, and
	// initialization vector size needs to be the same as block size (since the
	// IV is XOR'd with a block) -- so does the initialization vector need to be
	// the length of the key?
	// Create the initialization vector for CBC -- challenge says to use all zeros.
	var iv []byte
	for i := 0; i < len(key); i++ {
		iv = append(iv, 0)
	}

	// Check that ECB implementation is working correctly.
	checkECBPlaintext := []byte("encrypt, decrypt, & check result") // Length 32 -- 2 blocks
	ok, err := checkAESWithECB(checkECBPlaintext, key)
	switch {
	case err != nil:
		log.Fatal(err)
	case ok:
		fmt.Println("checkAESWithECB passed")
	case !ok:
		log.Fatal(fmt.Errorf("checkAESWithECB failed"))
	}

	// Decrypt the file with CBC.
	plaintext, err := decryptFile("/home/swalters4925/cryptopals/set2/challenge10/data.txt", key, iv)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(plaintext)

	// Check that CBC implementation is working correctly.
	ok, err = checkAESWithCBC([]byte(plaintext), key, iv)
	switch {
	case err != nil:
		log.Fatal(err)
	case ok:
		fmt.Println("checkAESWithCBC passed")
	case !ok:
		log.Fatal(fmt.Errorf("checkAESWithCBC failed"))
	}
}
