package main

import (
	"fmt"

	"cryptopals/set1/challenge1"
	"cryptopals/set1/challenge2"
	"cryptopals/set1/challenge3"
	"cryptopals/set1/challenge4"
	"cryptopals/set1/challenge5"
	"cryptopals/set1/challenge6"
	"cryptopals/set1/challenge7"
	"cryptopals/set1/challenge8"
)

func runChallenge(runFn func(), challengeNumber int) {
	fmt.Printf("Challenge %d:\n", challengeNumber)
	runFn()
	fmt.Println()
}

func main() {
	runChallenge(challenge1.Run, 1)
	runChallenge(challenge2.Run, 2)
	runChallenge(challenge3.Run, 3)
	runChallenge(challenge4.Run, 4)
	runChallenge(challenge5.Run, 5)
	runChallenge(challenge6.Run, 6)
	runChallenge(challenge7.Run, 7)
	runChallenge(challenge8.Run, 8)
}
