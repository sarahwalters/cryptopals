package challenge4

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"cryptopals/set1/challenge3"
)

func Run() {
	data, err := ioutil.ReadFile("/home/swalters4925/cryptopals/set1/challenge4/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var maxScore float32
	var bestDecoded string
	var bestEncodingRune rune

	for _, hexLine := range strings.Split(string(data), "\n") {
		bytes, err := hex.DecodeString(hexLine)
		if err != nil {
			log.Fatal(err)
		}

		decoded, encodingRune, score := challenge3.Decode(bytes)
		if score > maxScore {
			maxScore = score
			bestDecoded = decoded
			bestEncodingRune = encodingRune
		}
	}

	// "Now that the party is jumping\n"
	fmt.Printf("decoded: %q, encodingRune: %q\n", bestDecoded, bestEncodingRune)
}
