package features

import (
	"fmt"
	"strconv"
	"strings"
)

func sortSimilarity(result *[]string, resultSimilarity *[]int) {
	var temp1 int
	var temp2 string
	for i := 0; i < len(*resultSimilarity)-1; i++ {
		for j := i; j < len(*resultSimilarity)-1; j++ {
			if (*resultSimilarity)[j] >= (*resultSimilarity)[i] {
				// Swap both the result and the resultSimilarity
				temp1 = (*resultSimilarity)[i]
				(*resultSimilarity)[i] = (*resultSimilarity)[j]
				(*resultSimilarity)[j] = temp1

				temp2 = (*result)[i]
				(*result)[i] = (*result)[j]
				(*result)[j] = temp2
			}
		}
	}
}

func checkExactPattern(pattern string, text string, stringMatchingAlgo string) bool {
	if stringMatchingAlgo == "bm" {
		fmt.Println("Masuk bm")
		if len(BoyerMoore(pattern, text)) > 0 {
			return true
		} else {
			return false
		}
	} else {
		if len(Kmp(pattern, text)) > 0 {
			return true
		} else {
			return false
		}
	}
}

func stringMatchingLogic(pattern string, data map[string]string, stringMatchingAlgo string) []string {
	result := make([]string, 0)
	resultSimilarity := make([]int, 0)
	for k := range data { // Check if there is an exact match of the pattern in the database
		if checkExactPattern(pattern, k, stringMatchingAlgo) {
			fmt.Println("Masuk exact")
			result = append(result, k)
			return result
		}
	}
	for k := range data { // Check if there is a similiar pattern (similarity >= 90%) in the database
		if CalculateSimilarity(pattern, k) >= 90 {
			fmt.Println("Masuk 90")
			result = append(result, k)
			return result
		}
	}
	// If there is nothing else, pick 3 questions with the highest similarity
	fmt.Println("Masuk tidak ditemukan")
	for k := range data {
		result = append(result, k)
		resultSimilarity = append(resultSimilarity, int(CalculateSimilarity(pattern, k)))
	}
	sortSimilarity(&result, &resultSimilarity)
	return result
}

func ChatLogic(question string, data map[string]string, stringMatchingAlgo string) string {
	var answer strings.Builder
	extractedPattern := make([]string, 0)
	feature := WhichFeature(question)

	if feature == 4 {
		extractedPattern = append(extractedPattern, ExtractExpressionFour(question)[0])
		extractedPattern = append(extractedPattern, ExtractExpressionFour(question)[1])
	} else {
		extractedPattern = append(extractedPattern, ExtractExpression(question))
	}

	if feature == 1 {
		matchedString := stringMatchingLogic(question, data, stringMatchingAlgo)
		if len(matchedString) > 1 {
			answer.WriteString("Pertanyaan tidak ditemukan di database. Apakah maksud Anda:\n")
			for i := 0; i < len(matchedString); i++ {
				if i == 3 {
					break
				}
				strNum := strconv.Itoa(i + 1)
				answer.WriteString(strNum)
				answer.WriteString(". ")
				answer.WriteString(matchedString[i])
				answer.WriteString("\n")
			}
		} else {
			return data[matchedString[0]]
		}
	} else if feature == 2 {
		answer.WriteString("Hasilnya adalah ")
		answer.WriteString(Calculator(extractedPattern[0]))
	} else if feature == 3 {
		answer.WriteString("Hari ")
		answer.WriteString(CalculateDay(extractedPattern[0]))
	} else if feature == 4 {

	} else {

	}
	return answer.String()
}
