package features

import (
	"fmt"
	"regexp"
	"strings"
)

var pattern2a string = "^hitung\\s*(.*)"

// pattern2b := "\\s*([\\(\\)\\d]+)\\s*([\\+\\-\\*/^]\\s*[\\(\\)\\d]+\\s*)*"
var pattern3 string = "^(?:hari apa )?((0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[0-2])/\\d{1,})$"

// var pattern3 string = "(\\d{0,2}/\\d{0,2}/\\d{1,})"
var pattern4 string = ".*tambahkan\\s+pertanyaan\\s+(.+)\\s+dengan\\s+jawaban\\s+(.+)$"
var pattern5 string = ".*hapus\\s+pertanyaan\\s+(.+)$"

var MultipleQueries string = "^.*\\s+(dan|,|\\?)+?\\s+.*(?:dan|,|\\s+|\\?)*?.*$"

func WhichFeature(question string) int {
	regex2a := regexp.MustCompile(pattern2a)
	// regex2b := regexp.MustCompile(pattern2b)
	regex3 := regexp.MustCompile(pattern3)
	regex4 := regexp.MustCompile(pattern4)
	regex5 := regexp.MustCompile(pattern5)

	if regex4.MatchString(question) {
		return 4
	} else if regex5.MatchString(question) {
		return 5
	} else if regex2a.MatchString(question) {
		return 2
	} else if regex3.MatchString(question) {
		return 3
		// else if regex2b.MatchString(question) {return 2}
	} else {
		return 1
	}
}

func GetQuery(question string) (string, string) {
	regex := regexp.MustCompile(MultipleQueries)
	match := regex.FindStringSubmatch(question)
	if len(match) > 1 {
		// for i := 0; i < len(match); i++ {
		// 	fmt.Println("match ", i, match[i])
		// }
		return match[1], match[2]
	} else {
		return question, ""
	}
}

func GetQueries(questions string) []string {
	regex := regexp.MustCompile(MultipleQueries)
	match := regex.FindStringSubmatch(questions)
	if len(match) > 1 {
		var sep string
		if match[1] == "dan" {
			sep = "dan"
		} else if match[1] == "," {
			sep = ","
		} else if match[1] == "?" {
			sep = "?"
		}
		listOfQuestions := strings.Split(questions, sep)
		return listOfQuestions
	}
	return []string{questions}
}

func ExtractExpressionFour(question string) [2]string {
	extracted := [2]string{"Could not extract text", "Could not extract text"}
	regex4 := regexp.MustCompile(pattern4)

	match := regex4.FindStringSubmatch(question)
	if len(match) > 1 {
		extracted[0] = strings.TrimSpace(match[1])
		extracted[1] = strings.TrimSpace(match[2])
	}
	return extracted
}

func ExtractExpression(question string) string {
	regex2a := regexp.MustCompile(pattern2a)
	// regex2b := regexp.MustCompile(pattern2b)
	regex3 := regexp.MustCompile(pattern3)
	regex5 := regexp.MustCompile(pattern5)

	feature := WhichFeature(question)

	if feature == 2 {
		match := regex2a.FindStringSubmatch(question)
		if len(match) > 1 {
			return strings.TrimSpace(match[1])
		} else {
			return "Could not extract text"
		}
	} else if feature == 3 {
		match := regex3.FindStringSubmatch(question)
		if len(match) > 1 {
			return strings.TrimSpace(match[1])
		} else {
			return "Could not extract text"
		}
	} else if feature == 5 {
		match := regex5.FindStringSubmatch(question)
		if len(match) > 1 {
			return strings.TrimSpace(match[1])
		} else {
			return "Could not extract text"
		}
	} else {
		return strings.TrimSpace(question)
	}
}

func RegexDriver() {
	fmt.Println("Which Feature-test:")
	fmt.Println(WhichFeature("tambahkan pertanyaan 69? dengan jawaban nice!"))
	fmt.Println(WhichFeature("25/08/2023?"))
	fmt.Println(WhichFeature("hapus pertanyaan say my name"))
	fmt.Println(WhichFeature("hitung (4 + 6) * 2 / 5 - 3"))
	fmt.Println(WhichFeature("apa mata kuliah terbaik?"))
	fmt.Println()

	fmt.Println("Extract Expression-test:")
	fmt.Println(ExtractExpressionFour("tambahkan pertanyaan 69? dengan jawaban nice!"))
	fmt.Println(ExtractExpression("25/08/2023?"))
	fmt.Println(ExtractExpression("hapus pertanyaan say my name"))
	fmt.Println(ExtractExpression("hitung (4 + 6) * 2 / 5 - 3"))
	fmt.Println(ExtractExpression("apa mata kuliah terbaik?"))
	fmt.Println()
}
