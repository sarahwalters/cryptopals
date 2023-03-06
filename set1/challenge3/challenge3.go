package challenge3

import (
	"encoding/hex"
	"fmt"
	"log"
	"strings"
)

// Maps from uppercase letters to letter frequencies. Frequencies are from
// https://pi.math.cornell.edu/~mec/2003-2004/cryptography/subs/frequencies.html.
var frequencies = map[rune]float32{
	'E': 12.02,
	'T': 9.10,
	'A': 8.12,
	'O': 7.68,
	'I': 7.31,
	'N': 6.95,
	'S': 6.28,
	'R': 6.02,
	'H': 5.92,
	'D': 4.32,
	'L': 3.98,
	'U': 2.88,
	'C': 2.71,
	'M': 2.61,
	'F': 2.30,
	'Y': 2.11,
	'W': 2.09,
	'G': 2.03,
	'P': 1.82,
	'B': 1.49,
	'V': 1.11,
	'K': 0.69,
	'X': 0.17,
	'Q': 0.11,
	'J': 0.10,
	'Z': 0.07,
}

// Common non-letter runes that we might expect to appear in the decoded
// output.
// https://ee.hawaii.edu/~tep/EE160/Book/chap4/subsection2.1.1.1.html
var commonNonLetterRunes = map[rune]bool{
	' ':  true, // 32
	'!':  true, // 33
	'"':  true, // 34
	'\'': true, // 39
	'(':  true, // 41
	')':  true, // 42
	',':  true, // 44
	'.':  true, // 46
	':':  true, // 58
	';':  true, // 59
	'?':  true, // 63
}

// Chosen arbitrarily. Could adjust if it's not giving good decodings.
const uncommonRunePenalty = 10

// XORs the hex string with the given rune.
func xor(bytes []byte, r rune) string {
	xor := make([]byte, 0, len(bytes))
	for _, b := range bytes {
		xor = append(xor, b^byte(r))
	}

	return string(xor)
}

// Values letters proportionally to their frequency.
// Skips common non-letter runes that we might expect to appear in the
// decoded output, like spaces and quote marks.
// Penalizes all other runes.
func ComputeScore(plaintext string) float32 {
	var total float32

	for _, r := range strings.ToUpper(plaintext) {
		if frequency, ok := frequencies[r]; ok {
			total += frequency
		} else if !commonNonLetterRunes[r] {
			total = total - uncommonRunePenalty
		}
	}

	return total
}

// Decodes the input. Returns the decoded string, the rune
// the input was XOR'd with to encode it, and the score of
// the decoded string.
func Decode(bytes []byte) (string, rune, float32) {
	var maxScore float32
	var decoded string
	var encodingRune rune

	for i := 0; i <= 127; i++ {
		r := rune(i)
		plaintext := xor(bytes, r)
		score := ComputeScore(plaintext)

		if score > maxScore {
			maxScore = score
			decoded = plaintext
			encodingRune = r
		}
	}

	return decoded, encodingRune, maxScore
}

func Run() {
	inputHex := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	bytes, err := hex.DecodeString(inputHex)
	if err != nil {
		log.Fatal(err)
	}

	decoded, encodingRune, _ := Decode(bytes)

	// "Cooking MC's like a pound of bacon"
	fmt.Printf("decoded: %q; encoding rune: %q\n", decoded, encodingRune)
}
