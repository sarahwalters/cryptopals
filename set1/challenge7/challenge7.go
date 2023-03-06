package challenge7

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
)

type Block struct {
  // This block is from [Start, End) in the original data.
  Start, End int
  // The block of data.
  Bytes []byte
}

// Splits the data into blocks of size blockSize.
// The last block will be smaller than blockSize if len(data) is not divisible
// by blockSize -- this function does not pad the last block.
func Blocks(data []byte, blockSize int) []Block {
  start, end := 0, blockSize
  var result []Block
  for {
    if end >= len(data) {
      result = append(result, Block{
        Start: start,
        End: end,
        Bytes: data[start:],
      })
      break
    }
    
    result = append(result, Block{
        Start: start,
        End: end,
        Bytes: data[start:end],
      })

    start += blockSize
    end += blockSize
  }
  return result
}

// https://stackoverflow.com/questions/24072026/golang-aes-ecb-encryption
// https://en.wikipedia.org/wiki/Block_cipher_mode_of_operation#Electronic_Codebook_.28ECB.29
func DecryptAESWithECB(ciphertext, key []byte) ([]byte, error) {
	// The crypto/aes package will choose AES-128 if the key is 16 bytes long,
	// AES-192 if the key is 24 bytes long, or AES-256 if the key is 32 bytes
	// long. See https://pkg.go.dev/crypto/aes#NewCipher.
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	
	decrypted := make([]byte, len(ciphertext))

	// QUESTION: Does ECB block size have to be the length of the key?
	for _, b := range Blocks(ciphertext, len(key)) {
		cipher.Decrypt(decrypted[b.Start:b.End], ciphertext[b.Start:b.End])
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
