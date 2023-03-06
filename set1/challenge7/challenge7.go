package challenge7

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
)

// https://stackoverflow.com/questions/24072026/golang-aes-ecb-encryption
// https://en.wikipedia.org/wiki/Block_cipher_mode_of_operation#Electronic_Codebook_.28ECB.29
func DecryptAESWithECB(data, key []byte) ([]byte, error) {
	// The crypto/aes package will choose AES-128 if the key is 16 bytes long,
	// AES-192 if the key is 24 bytes long, or AES-256 if the key is 32 bytes
	// long. See https://pkg.go.dev/crypto/aes#NewCipher.
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	decrypted := make([]byte, len(data))

	// It looks like AES uses a block size of 16 bytes for ECB mode, though
	// I can't find a clear link, just multiple stackoverflow posts etc. Or
	// maybe the ECB block size is whatever the key length is.
	ecbBlockSizeBytes := 16

	for bs, be := 0, ecbBlockSizeBytes; bs < len(data); bs, be = bs+ecbBlockSizeBytes, be+ecbBlockSizeBytes {
		cipher.Decrypt(decrypted[bs:be], data[bs:be])
	}

	return decrypted, nil
}

func Run() {
	raw, err := ioutil.ReadFile("/home/swalters4925/cryptopals/set1/challenge7/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	bytes, err := base64.StdEncoding.DecodeString(string(raw))
	if err != nil {
		log.Fatal(err)
	}

	decrypted, err := DecryptAESWithECB(bytes, []byte("YELLOW SUBMARINE"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(decrypted))
}
