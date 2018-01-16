package main

import (
	"fmt"
	"math/rand"
)

func main() {
	alpha := []string{
		"a",
		"aa",
		"i",
		"ii",
		"u",
		"uu",
		"ru",
		"ye",
		"yee",
		"ai",
		"wo",
		"woo",
		"au",
		"am",
		"aha",
		"ka",
		"kha",
		"ga",
		"gha",
		"nya",
		"cha",
		"chha",
		"ja",
		"jha",
		"inya",
		"ta",
		"tha (back)",
		"da (back)",
		"dha (back)",
		"arha",
		"t'a",
		"tha (forward)",
		"da (forward)",
		"dha (forward)",
		"na",
		"pa",
		"fa",
		"ba",
		"bha",
		"ma",
		"ya",
		"ra",
		"la",
		"wa",
		"sha",
		"shha",
		"sa",
		"ha",
		"la(alt)",
	}

	list := rand.Perm(49)

	for i, v := range list {
		fmt.Printf("%s ", alpha[v])
		if i%5 == 4 {
			fmt.Print("\n")
		}
	}
}
