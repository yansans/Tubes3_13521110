package features

import (
	"fmt"
	"sort"
)

func generateLPS(pattern string) []int {
	lps := make([]int, len(pattern))
	i := 0
	j := 1

	for j < len(pattern) {
		if pattern[j:j+1] == pattern[i:i+1] {
			lps[j] = i + 1
			i++
		} else {
			i = 0
		}
		j++
	}
	return lps
}

func kmp(pattern string, text string) []int {
	// Main KMP Algorithm
	lps := generateLPS(pattern)
	i := 0
	j := -1

	for i < len(text) {
		if pattern[j+1] == text[i] {
			i++
			j++
		} else {
			if j == -1 {
				i++
			} else {
				j = lps[j] - 1
			}
		}

		if j == len(pattern)-1 {
			break
		}
	}

	return generateMatchedIdx(pattern, text, i-1)
}

func generateBMT(pattern string) map[byte]int {
	bmt := make(map[byte]int)
	var star byte = '*'

	for i := 0; i < len(pattern); i++ {
		if len(pattern)-i-1 > 1 {
			bmt[pattern[i]] = len(pattern) - i - 1
		} else {
			bmt[pattern[i]] = 1
		}
		bmt[star] = len(pattern)
	}

	return bmt
}

func bmMatchPattern(pattern string, text string, k *int) bool {
	j := len(pattern) - 1
	for j > 0 {
		*k--
		j--
		if text[*k] != pattern[j] {
			return false
		}
	}
	return true
}

func boyerMoore(pattern string, text string) []int {
	// Main Boyer-Moore Algorithm
	bmt := generateBMT(pattern)
	i := len(pattern) - 1
	j := len(pattern) - 1
	k := -1
	var star byte = '*'

	for i < len(text) {
		k = i
		if text[k] == pattern[j] {
			if bmMatchPattern(pattern, text, &k) {
				break
			}
		}
		if _, ok := bmt[text[k]]; ok {
			i += bmt[text[k]]
		} else {
			i += bmt[star]
		}
	}
	return generateMatchedIdx(pattern, text, i)
}

func generateMatchedIdx(pattern string, text string, idx int) []int {
	// Generate indexes of the matching pattern in the text
	matchIndexes := make([]int, 0)
	for k := len(pattern); k > 0; k-- {
		matchIndexes = append(matchIndexes, idx)
		idx--
	}
	sort.Ints(matchIndexes)

	return matchIndexes
}

// Driver
func stringMatchingDriver() {
	var hello string = "hello, World!"
	lps := make([]int, len(hello))
	fmt.Println(hello[12:13])
	fmt.Println(len(hello))
	fmt.Println(lps)
	fmt.Println()

	fmt.Println("LPS-test:")
	fmt.Println(generateLPS("abcdabeabf"))
	fmt.Println(generateLPS("abcdeabfabc"))
	fmt.Println(generateLPS("aabcadaabe"))
	fmt.Println(generateLPS("aaaabaacd"))
	fmt.Println()

	fmt.Println("KMP-test:")
	fmt.Println(kmp("ababd", "ababcabcabababd"))
	fmt.Println(kmp("aaab", "aaaaaaab"))
	fmt.Println(kmp("BAB", "ABABABAC"))
	fmt.Println(kmp("TEST", "THIS IS A TEST"))
	fmt.Println()

	fmt.Println("BM-test:")
	fmt.Println(boyerMoore("ababd", "ababcabcabababd"))
	fmt.Println(boyerMoore("aaab", "aaaaaaab"))
	fmt.Println(boyerMoore("BAB", "ABABABAC"))
	fmt.Println(boyerMoore("TEST", "THIS IS A TEST"))
	fmt.Println()
}
