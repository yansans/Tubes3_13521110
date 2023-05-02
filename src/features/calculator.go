package main

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

// func infixToPostfix(infix string) {
// 	j := 0
// 	stack := Stack{}
// 	postfix := make([]int, 0)
// 	top := -1

// 	for i := 0; i < len(infix); i++ {
// 		if infix[i] == '(' {
// 			continue
// 		} else if infix[i] == ')' {
// 			continue
// 		} else if isOperator(infix[i]) {
// 			continue
// 		} else {

// 		}
// 	}

// }
