package main

import (
	"fmt"
	"strings"
	"unicode"
)

func wordcount(text string) map[string]int {
	clean := strings.Map(func(r rune) rune {
		if unicode.IsPunct(r) {
			return -1
		}
		return r
	}, text)

	word := strings.Fields(strings.ToLower(clean))

	freq := make(map[string]int)
	for _, w := range word {
		freq[w]++

	}
	return freq
}

// palindrom problem
func palindrome(s string) bool {
	var filtered strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			filtered.WriteRune(unicode.ToLower(r))
		}
	}
	str := filtered.String()
	left, right := 0, len(str)-1
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--

	}
	return true
}

func main() {
	fmt.Println(wordcount("racecar drivers are talented."))
	fmt.Println(wordcount("buses are long buses as weird"))

	fmt.Println(palindrome("racecar"))
	fmt.Println(palindrome("bus"))
}
