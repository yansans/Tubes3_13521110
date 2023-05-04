package features

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func precedence(operator byte) int {
	switch operator {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	case '^':
		return 3
	default:
		return -1
	}
}

func isOperator(char byte) bool {
	return char == '+' || char == '-' || char == '*' || char == '/' || char == '^'
}

func isOperatorStr(str string) bool {
	return str == "+" || str == "-" || str == "*" || str == "/" || str == "^"
}

func writeToPostfix(postfix *[]string, currentString *strings.Builder) {
	if currentString.String() != "" {
		*postfix = append(*postfix, currentString.String())
		currentString.Reset()
	}
}

func isBeginEndOperator(infix string) bool {
	// Checks if the beginning or end of the infix expression is an operator
	startIdx := 0
	endIdx := len(infix) - 1
	for infix[startIdx] == ' ' || infix[startIdx] == '\t' {
		startIdx++
	}
	for infix[endIdx] == ' ' || infix[endIdx] == '\t' {
		endIdx--
	}
	if isOperator(infix[startIdx]) || isOperator(infix[endIdx]) {
		return true
	} else {
		return false
	}
}

func infixToPostfix(infix string) []string {
	stack := Stack{}
	postfix := make([]string, 0)
	operatorCount := 0 // Count of the operator after an operand or a parenthesis
	var currentString strings.Builder

	if isBeginEndOperator(infix) {
		postfix = postfix[:0] // Clear postfix
		return postfix        // Signifies invalid expression
	}

	// Main infix to postfix conversion algorithm
	for i := 0; i < len(infix); i++ {
		if infix[i] == ' ' || infix[i] == '\t' {
			continue
		} else if infix[i] == '(' {
			if i+1 != len(infix) && isOperator(infix[i+1]) { // Checks if there is an operator after (
				postfix = postfix[:0] // Clear postfix
				return postfix        // Signifies invalid expression
			}
			operatorCount = 0
			writeToPostfix(&postfix, &currentString)

			stack.Push(infix[i])
		} else if infix[i] == ')' {
			if i != 0 && isOperator(infix[i-1]) { // Checks if there is an operator before )
				postfix = postfix[:0] // Clear postfix
				return postfix        // Signifies invalid expression
			}
			operatorCount = 0
			writeToPostfix(&postfix, &currentString)

			for !stack.IsEmpty() && stack.Peek().(byte) != '(' {
				// Pop and write to postfix
				currentString.WriteByte(stack.Pop().(byte))
				writeToPostfix(&postfix, &currentString)
			}
			if (!stack.IsEmpty() && stack.Peek().(byte) != '(') || stack.IsEmpty() {
				postfix = postfix[:0] // Clear postfix
				return postfix        // Signifies invalid expression
			}
			stack.Pop()
		} else if isOperator(infix[i]) {
			operatorCount++
			writeToPostfix(&postfix, &currentString)

			for !stack.IsEmpty() && precedence(stack.Peek().(byte)) >= precedence(infix[i]) && stack.Peek().(byte) != '(' {
				// Pop and write to postfix
				currentString.WriteByte(stack.Pop().(byte))
				writeToPostfix(&postfix, &currentString)
			}
			stack.Push(infix[i])
		} else {
			operatorCount = 0
			currentString.WriteByte(infix[i])
		}

		if operatorCount > 1 {
			postfix = postfix[:0] // Clear postfix
			return postfix        // Signifies invalid expression
		}
	}
	writeToPostfix(&postfix, &currentString)
	// Pop all remaining elements from the stack
	for !stack.IsEmpty() {
		if stack.Peek().(byte) == '(' {
			postfix = postfix[:0] // Clear postfix
			return postfix        // Signifies invalid expression
		}
		// Pop and write to postfix
		currentString.WriteByte(stack.Pop().(byte))
		writeToPostfix(&postfix, &currentString)
	}

	return postfix
}

func operate(op1 float64, op2 float64, operator string) float64 {
	switch operator {
	case "+":
		return op1 + op2
	case "-":
		return op1 - op2
	case "*":
		return op1 * op2
	case "/":
		return op1 / op2
	case "^":
		return math.Pow(op1, op2)
	default:
		return 0
	}
}

func Calculator(expression string) string {
	postfix := infixToPostfix(expression)
	stack := Stack{}
	var op1, op2, temp float64

	if len(postfix) == 0 {
		return "Sintaks persamaan tidak sesuai"
	} else {
		for i := 0; i < len(postfix); i++ {
			if isOperatorStr(postfix[i]) {
				op2 = stack.Pop().(float64)
				op1 = stack.Pop().(float64)
				temp = operate(op1, op2, postfix[i])
				stack.Push(temp)
			} else {
				f, err := strconv.ParseFloat(postfix[i], 64)
				if err != nil {
					return "Sintaks persamaan tidak sesuai"
				}
				stack.Push(f)
			}
		}
		return strconv.FormatFloat(stack.Peek().(float64), 'f', -1, 64)
	}
}

func CalculatorDriver() {
	fmt.Println("Infix to Postfix-Test:")
	infix := "a+b*(c^d-e)^(f+g*h)-i"
	postfix := infixToPostfix(infix)
	var test strings.Builder

	for i := 0; i < len(postfix); i++ {
		test.WriteString(postfix[i])
	}
	fmt.Println(postfix)
	fmt.Println(test.String())
	fmt.Println()

	fmt.Println("Calculator-Test:")
	expression1 := "(4 + 6) * 2 / 5 - 3"                                   // 1
	expression2 := "3.2 *(4+ 5.010) -12/ 3"                                // 24.832
	expression3 := "4 + 5 * ( 2 - 7 ))"                                    // Invalid
	expression4 := "+ 3.2 *(4+ 5.010) -12/ 3"                              // Invalid
	expression5 := "a+b*(c^d-e)^(f+g*h)-i"                                 // Invalid (Should be numbers)
	expression6 := "4 + 5 * ( 2 - 7..021 ))"                               // Invalid
	expression7 := "(((12 + 7) / 3) * (8 - 4)) + (10 ^ 3) - ((5 * 2) / 4)" // 1022.8333...
	result1 := Calculator(expression1)
	result2 := Calculator(expression2)
	result3 := Calculator(expression3)
	result4 := Calculator(expression4)
	result5 := Calculator(expression5)
	result6 := Calculator(expression6)
	result7 := Calculator(expression7)
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	fmt.Println(result6)
	fmt.Println(result7)
	fmt.Println()
}
