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

func Kmp(pattern string, text string) []int {
	empty := make([]int, 0)
	if len(pattern) == 0 {
		return empty
	}
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
			return generateMatchedIdx(pattern, text, i-1)
		}
	}

	return empty
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

func BoyerMoore(pattern string, text string) []int {
	empty := make([]int, 0)
	if len(pattern) == 0 {
		return empty
	}
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
	if i == len(text) {
		return empty
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

func min(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0]
	for _, num := range nums {
		if num < res {
			res = num
		}
	}
	return res
}

func levenstheinDistance(pattern string, text string) int {
	m := len(pattern)
	n := len(text)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}
	for j := 1; j <= n; j++ {
		for i := 1; i <= m; i++ {
			if pattern[i-1] == text[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			}
		}
	}
	return dp[m][n]
}

func CalculateSimilarity(pattern string, text string) float64 {
	if len(pattern) >= len(text) {
		return (1 - float64(levenstheinDistance(pattern, text))/float64(len(pattern))) * 100
	} else {
		return (1 - float64(levenstheinDistance(pattern, text))/float64(len(text))) * 100
	}
}

// Driver
func StringMatchingDriver() {
	// var hello string = "hello, World!"
	// lps := make([]int, len(hello))
	// fmt.Println(hello[12:13])
	// fmt.Println(len(hello))
	// fmt.Println(lps)
	// fmt.Println()

	// fmt.Println("LPS-test:")
	// fmt.Println(generateLPS("abcdabeabf"))
	// fmt.Println(generateLPS("abcdeabfabc"))
	// fmt.Println(generateLPS("aabcadaabe"))
	// fmt.Println(generateLPS("aaaabaacd"))
	// fmt.Println()

	// fmt.Println("KMP-test:")
	// fmt.Println(Kmp("ababd", "ababcabcabababd"))
	// fmt.Println(Kmp("", "aaaaaaab"))
	// fmt.Println(Kmp("CCC", "ABABABAC"))
	// fmt.Println(Kmp("TEST", "THIS IS A TEST"))
	// fmt.Println(Kmp("ibu", "Apa ibukota Indonesia?"))
	// fmt.Println()

	// // fmt.Println("BM-test:")
	// fmt.Println(BoyerMoore("ababd", "ababcabcabababd"))
	// fmt.Println(BoyerMoore("", "aaaaaaab"))
	// fmt.Println(BoyerMoore("ACC", "ABABABAC"))
	// fmt.Println(BoyerMoore("TEST", "THIS IS A TEST"))
	// fmt.Println(BoyerMoore("something", "Apa ibukota Indonesia?"))
	// fmt.Println()

	fmt.Println("Levensthein Distance-test:")
	s1 := "a"
	t1 := "apa mata kuliah IF semester 4 yang paling seru?"
	dist1 := levenstheinDistance(s1, t1)
	fmt.Printf("Levenshtein distance between %q and %q is %d.\n", s1, t1, dist1)
	s2 := "book"
	t2 := "book1"
	dist2 := levenstheinDistance(s2, t2)
	fmt.Printf("Levenshtein distance between %q and %q is %d.\n", s2, t2, dist2)
	s3 := "book1"
	t3 := "book"
	dist3 := levenstheinDistance(s3, t3)
	fmt.Printf("Levenshtein distance between %q and %q is %d.\n", s3, t3, dist3)
	s4 := "Something that matters"
	t4 := "It doesn't matter"
	dist4 := levenstheinDistance(s4, t4)
	fmt.Printf("Levenshtein distance between %q and %q is %d.\n", s4, t4, dist4)
	fmt.Println()

	fmt.Println("Similarity-test:")
	simil1 := CalculateSimilarity(s1, t1)
	fmt.Printf("Similarity between %q and %q is %f.\n", s1, t1, simil1)
	simil2 := CalculateSimilarity(s2, t2)
	fmt.Printf("Similarity between %q and %q is %f.\n", s2, t2, simil2)
	simil3 := CalculateSimilarity(s3, t3)
	fmt.Printf("Similarity between %q and %q is %f.\n", s3, t3, simil3)
	simil4 := CalculateSimilarity(s4, t4)
	fmt.Printf("Similarity between %q and %q is %f.\n", s4, t4, simil4)
	fmt.Println()
}
