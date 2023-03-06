package challenge8

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

// Counts the number of times each byte in the hex string (2 hex characters) appears,
// and takes the average of the counts.
func averageByteCount(hexStr string) float32 {
	byteCounts := make(map[string]int)
	for start, end := 0, 2; end <= len(hexStr); start, end = start+2, end+2 {
		byteCounts[hexStr[start:end]]++
	}

	var sumOfByteCounts int
	for _, count := range byteCounts {
		sumOfByteCounts += count
	}
	return float32(sumOfByteCounts) / float32(len(byteCounts))
}

type scoredHexStr struct {
	// Which line the hex string is on in the input file.
	inputIndex int
	// The hex string itself.
	hexStr string
	// Average of byte frequencies in the hex string.
	// From https://en.wikipedia.org/wiki/Block_cipher_mode_of_operation#Electronic_Codebook_.28ECB.29:
	// "Because ECB encrypts identical plaintext blocks into identical ciphertext blocks, it does not hide data patterns well."
	// Idea: every 2 hex characters represent a byte -> 1 character; split string into 2-hex-char pairs and count frequencies.
	// If there is a string that has an outlier number of repeated characters, that's likely the one that is ECB-encrypted.
	// So, a higher score indicates that the string is likelier to be ECB-encrypted.
	score float32
}

// Scores hex strings for how likely they are to be ECB-encrypted. Returns a
// slice of scored hex strings, in descending order by score.
func scoreHexStrs(hexStrs []string) []scoredHexStr {
	var result []scoredHexStr
	for i := 0; i < len(hexStrs); i++ {
		result = append(result, scoredHexStr{
			inputIndex: i,
			hexStr:     hexStrs[i],
			score:      averageByteCount(hexStrs[i]),
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].score > result[j].score
	})
	return result
}

func Run() {
	raw, err := ioutil.ReadFile("/home/swalters4925/cryptopals/set1/challenge8/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var hexStrs []string
	for _, hexStr := range strings.Split(string(raw), "\n") {
		hexStrs = append(hexStrs, hexStr)
	}

	scoredHexStrs := scoreHexStrs(hexStrs)

	fmt.Printf("Likeliest to be ECB-encoded: %+v\n", scoredHexStrs[0])

	var allScores []float32
	for _, scoredHexStr := range scoredHexStrs {
		allScores = append(allScores, scoredHexStr.score)
	}
	fmt.Printf("All scores in descending order: %+v\n", allScores)
}
