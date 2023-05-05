package features

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/yansans/Tubes3_13521110/src/models"
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

func stringMatchingLogic(pattern string, data map[string]string, stringMatchingAlgo string) ([]string, bool) {
	result := make([]string, 0)
	resultSimilarity := make([]int, 0)
	// Check if there is an exact match of the pattern in the database
	for k := range data {
		if checkExactPattern(pattern, k, stringMatchingAlgo) {
			fmt.Println("Masuk exact")
			result = append(result, k)
			resultSimilarity = append(resultSimilarity, int(CalculateSimilarity(pattern, k)))
		}
	}
	if len(result) > 0 {
		sortSimilarity(&result, &resultSimilarity)
		return result, false
	}
	// Check if there is a similiar pattern (similarity >= 90%) in the database
	for k := range data {
		if CalculateSimilarity(pattern, k) >= 90 {
			fmt.Println("Masuk 90")
			result = append(result, k)
			resultSimilarity = append(resultSimilarity, int(CalculateSimilarity(pattern, k)))
		}
	}
	if len(result) > 0 {
		sortSimilarity(&result, &resultSimilarity)
		return result, false
	}
	// If there is nothing else, pick 3 questions with the highest similarity
	fmt.Println("Masuk tidak ditemukan")
	for k := range data {
		result = append(result, k)
		resultSimilarity = append(resultSimilarity, int(CalculateSimilarity(pattern, k)))
	}
	sortSimilarity(&result, &resultSimilarity)
	return result, true
}

func GetAnswer(questions string, data map[string]string, stringMatchAlgo string) string {
	var answers strings.Builder
	question := GetQueries(questions)
	for i := 0; i < len(question); i++ {
		// println(i, question[i])
		answers.WriteString(ChatLogic(strings.TrimSpace(question[i]), data, stringMatchAlgo))
		answers.WriteString("\n")
	}
	return answers.String()
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
		matchedString, notFound := stringMatchingLogic(question, data, stringMatchingAlgo)
		if notFound {
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
		result := Calculator(extractedPattern[0])
		if result != "Sintaks persamaan tidak sesuai" {
			answer.WriteString("Hasilnya adalah ")
		}
		answer.WriteString(Calculator(extractedPattern[0]))
	} else if feature == 3 {
		answer.WriteString("Hari ")
		answer.WriteString(CalculateDay(extractedPattern[0]))
	} else if feature == 4 {
		found := false
		for k := range data {
			if checkExactPattern(extractedPattern[0], k, stringMatchingAlgo) {
				answer.WriteString("Pertanyaan ada di database")
				found = true
			}
			if found {
				answer.WriteString("Pertanyaan ditambah ke database")
				break
			}
		}
		if !found {
			answer.WriteString("Pertanyaan belum ada di database")
			answer.WriteString("Pertanyaan ditambah ke database")
		}
		var question models.Query
		question.Question = extractedPattern[0]
		question.Answer = extractedPattern[1]
		addQuestion(question)
	} else if feature == 5 {
		found := false
		for k := range data {
			if checkExactPattern(extractedPattern[0], k, stringMatchingAlgo) {
				answer.WriteString("Pertanyaan ada di database\n")
				answer.WriteString("Pertanyaan dihapus dari database")
				found = true
				break
			}
		}
		if found {
			var question models.Query
			question.Question = extractedPattern[0]
			deleteQuestion(question)
		} else {
			answer.WriteString("Pertanyaan tidak ada di database")
		}
	}
	return answer.String()
}

func addQuestion(add models.Query) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	host := "localhost:6969"
	url := "http://" + host + "/query"
	data, err := json.Marshal(add)
	if err != nil {
		fmt.Println(err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	if res.StatusCode != http.StatusCreated {
		fmt.Println(res.StatusCode)
		fmt.Println(string(body))
		return
	} else {
		fmt.Println("Pertanyaan berhasil ditambahkan")
	}
}

func deleteQuestion(question models.Query) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	host := "localhost:6969"
	url := "http://" + host + "/query"
	data, err := json.Marshal(question)
	if err != nil {
		println("marshal error")
		fmt.Println(err)
		return
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, bytes.NewBuffer(data))
	if err != nil {
		println("new request error")
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		println("client do error")
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		println("read body error")
		fmt.Println(err)
		return
	}
	if res.StatusCode != http.StatusOK {
		println("status not ok")
		fmt.Println(res.StatusCode)
		fmt.Println(string(body))
		return
	} else {
		fmt.Println("Pertanyaan berhasil dihapus")
	}
}
