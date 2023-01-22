package main

import (
	"fmt"
)

// Runtime: 4 ms. Beats: 64.88% Memory: 2.4 MB. Beats:12.50%
func getHint(secret string, guess string) string {
	a := 0
	b := 0
	count := make([]int, 10)
	for i := range secret {
		count[byte(secret[i])-48]++
	}
	done := make([]bool, len(secret))
	for i := range guess {
		if guess[i] == secret[i] {
			count[byte(secret[i])-48]--
			done[i] = true
			a++
		}
	}
	for i := range guess {
		if count[byte(guess[i])-48] > 0 && !done[i] {
			b++
			count[byte(guess[i])-48]--
		}
	}
	return fmt.Sprint(a, "A", b, "B")
}

func main() {

}
