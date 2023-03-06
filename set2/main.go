package main

import (
	"fmt"

	"cryptopals/set2/challenge10"
	"cryptopals/set2/challenge9"
)

func runChallenge(runFn func(), challengeNumber int) {
	fmt.Printf("Challenge %d:\n", challengeNumber)
	runFn()
	fmt.Println()
}

func main() {
	runChallenge(challenge9.Run, 9)
	runChallenge(challenge10.Run, 10)
}
