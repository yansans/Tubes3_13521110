package features

import (
	"fmt"
	"strconv"
	"strings"
)

func parseDate(date string) [3]int {
	var tempString strings.Builder
	intDate := [3]int{5, 5, 2023}
	// Parse day
	j := 0
	for i := 0; i < len(date); i++ {
		if date[i] == '/' {
			num, err := strconv.Atoi(tempString.String())
			if err != nil {
				return intDate
			}
			intDate[j] = num
			j++
			tempString.Reset()
		}
		tempString.WriteByte(date[i])
	}

	return intDate
}

func toJulian(date string) int {
	intDate := parseDate(date)
	a := 0

	if intDate[1] == 1 || intDate[1] == 2 {
		a = 1
	}

	y := intDate[2] + 4800 - a
	m := intDate[1] + (12 * a) - 3

	return intDate[0] + (((153 * m) + 2) / 5) + (365 * y) + (y / 4) - (y / 100) + (y / 400) - 32045
}

func createDays() map[int]string {
	days := make(map[int]string)

	days[0] = "Senin"
	days[1] = "Selasa"
	days[2] = "Rabu"
	days[3] = "Kamis"
	days[4] = "Jumat"
	days[5] = "Sabtu"
	days[6] = "Minggu"

	return days
}

func createReverseDays() map[int]string {
	days := make(map[int]string)

	days[0] = "Senin"
	days[1] = "Minggu"
	days[2] = "Sabtu"
	days[3] = "Jumat"
	days[4] = "Kamis"
	days[5] = "Rabu"
	days[6] = "Selasa"

	return days
}

func CalculateDay(date string) string {
	targetDate := toJulian(date)
	referenceDate := toJulian("08/05/2023")
	days := createDays()
	noOfDays := 0

	if targetDate >= referenceDate {
		noOfDays = targetDate - referenceDate
	} else {
		days = createReverseDays()
		noOfDays = referenceDate - targetDate
	}
	return days[noOfDays%7]
}

func DayDriver() {
	fmt.Println(CalculateDay("25/08/2023")) // Jumat
	fmt.Println(CalculateDay("08/03/2025")) // Sabtu
	fmt.Println(CalculateDay("19/02/2020")) // Rabu
	fmt.Println(CalculateDay("04/09/476"))  // The day which the Roman Empire Fall
}
