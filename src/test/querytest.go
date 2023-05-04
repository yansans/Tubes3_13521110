package main

import "github.com/yansans/Tubes3_13521110/src/controllers"

func main() {
	listOfQuestion := controllers.GetQuestionList()
	questionMap := controllers.GetQuestionMap()

	println("List of question: ")
	for _, question := range listOfQuestion {
		println(question)
	}
	println()
	println("Question map: ")
	for key, value := range questionMap {
		print(
			key, " : ", value, "\n",
		)
	}

}
