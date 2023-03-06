package challenge6

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"sort"

	"cryptopals/set1/challenge3"
)

// Converts a string to a bit string.
// https://stackoverflow.com/questions/37349071/golang-how-to-convert-string-to-binary-representation
func bitString(s string) string {
	res := ""
	for _, c := range s {
		res = fmt.Sprintf("%s%.8b", res, c)
	}
	return res
}

// Returns the number of differing bits between the two input strings.
// Assumes the strings are the same length.
func hammingDistance(a, b string) int {
	aBits, bBits := bitString(a), bitString(b)
	var distance int
	for i := 0; i < len(aBits); i++ {
		aBit, bBit := aBits[i], bBits[i]
		if aBit != bBit {
			distance++
		}
	}

	return distance
}

type scoredKeySize struct {
	keySize int
	// Average of pairwise Hamming distances between the first four sets of
	// keySize bytes of the data, divided by keySize to normalize across
	// different key sizes. A lower score indicates that the key size is a
	// good match for the data. (TODO -- think / read about why this is)
	score float32
}

// Scores potential key sizes for decoding the data. Returns a list of key
// sizes from 2 through 40, in ascending order by score. See the scoredKeySize
// struct for how the score is calculated.
func scoreKeySizes(data []byte) ([]scoredKeySize, error) {
	var result []scoredKeySize

	for keySize := 2; keySize <= 40; keySize++ {
		firstKeySizeBytes := string(data[0:keySize])
		secondKeySizeBytes := string(data[keySize : 2*keySize])
		thirdKeySizeBytes := string(data[2*keySize : 3*keySize])
		fourthKeySizeBytes := string(data[3*keySize : 4*keySize])

		hd12 := hammingDistance(firstKeySizeBytes, secondKeySizeBytes)
		hd13 := hammingDistance(firstKeySizeBytes, thirdKeySizeBytes)
		hd14 := hammingDistance(firstKeySizeBytes, fourthKeySizeBytes)
		hd23 := hammingDistance(secondKeySizeBytes, thirdKeySizeBytes)
		hd24 := hammingDistance(secondKeySizeBytes, fourthKeySizeBytes)
		hd34 := hammingDistance(thirdKeySizeBytes, fourthKeySizeBytes)
		hdAvg := float32(hd12+hd13+hd14+hd23+hd24+hd34) / 6

		result = append(result, scoredKeySize{
			keySize: keySize,
			score:   hdAvg / float32(keySize),
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].score < result[j].score
	})
	return result, nil
}

func breakIntoBlocks(data []byte, blockSize int) [][]byte {
	var blocks [][]byte
	var currentStart int

	// len(data) isn't necessarily divisible by blockSize.
	// Handle blocks that are definitely complete first.
	for currentStart < len(data)-blockSize {
		blocks = append(blocks, data[currentStart:currentStart+blockSize])
		currentStart = currentStart + blockSize
	}
	// Handle the last, possibly incomplete block.
	blocks = append(blocks, data[currentStart:])

	return blocks
}

func transposeBlocks(blocks [][]byte) [][]byte {
	if len(blocks) == 0 {
		return nil
	}

	// Use the length of the first input block to initialize slices that
	// represent the transposed blocks.
	transposed := make([][]byte, len(blocks[0]))

	// Add the values for the blocks to the initialized slices.
	for _, block := range blocks {
		for i, value := range block {
			transposed[i] = append(transposed[i], value)
		}
	}

	return transposed
}

func decodeVigenere(bytes []byte) (string, string, error) {
	scoredKeySizes, err := scoreKeySizes(bytes)
	if err != nil {
		return "", "", err
	}
	if len(scoredKeySizes) == 0 {
		return "", "", fmt.Errorf("len(scoredKeySizes) must be greater than zero")
	}

	blocks := breakIntoBlocks(bytes, scoredKeySizes[0].keySize)
	transposed := transposeBlocks(blocks)

	var key []rune
	var decodedBlocks [][]rune
	for _, t := range transposed {
		decodedBlock, encodingRune, _ := challenge3.Decode(t)
		decodedBlocks = append(decodedBlocks, []rune(decodedBlock))
		key = append(key, encodingRune)
	}

	if len(decodedBlocks) == 0 {
		return "", "", fmt.Errorf("len(decodedBlocks) must be greater than zero")
	}
	var decoded []rune
	for i := range decodedBlocks[0] {
		for _, decodedBlock := range decodedBlocks {
			if i < len(decodedBlock) {
				decoded = append(decoded, decodedBlock[i])
			}
		}
	}

	return string(decoded), string(key), nil
}

func Run() {
	raw, err := ioutil.ReadFile("/home/swalters4925/cryptopals/set1/challenge6/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	bytes, err := base64.StdEncoding.DecodeString(string(raw))
	if err != nil {
		log.Fatal(err)
	}

	decoded, key, err := decodeVigenere(bytes)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("decoded: %q\n", decoded)
	fmt.Printf("key: %q\n", key)
}
