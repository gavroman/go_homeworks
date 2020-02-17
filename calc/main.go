package main

import (
	"flag"
	"fmt"
	Stack "github.com/golang-collections/collections/stack"
)

func isDigit(char int32) bool {
	return char >= '0' && char <= '9'

}

func main() {
	flag.Parse()
	args := flag.Args()

	expressionString := ""
	for i := range args {
		expressionString += args[i]
	}

	outputArray := make([]string, 0, len(expressionString))
	operationStack := Stack.New()
	numberString := ""
	for _, char := range expressionString {
		if isDigit(char) {
			numberString += string(char)
		} else {
			if numberString != "" {
				outputArray = append(outputArray, numberString)
				numberString = ""
			}
			switch char {
			case '+', '-':
				for operationStack.Len() != 0 {
					if operationStack.Peek() == "(" {
						break
					}
					got := fmt.Sprintf("%v", operationStack.Pop())
					outputArray = append(outputArray, got)
				}
				operationStack.Push(string(char))

			case '*', '/', '(':
				operationStack.Push(string(char))

			case ')':
				for operationStack.Len() != 0 {
					if operationStack.Peek() == "(" {
						break
					}
					got := fmt.Sprintf("%v", operationStack.Pop())
					outputArray = append(outputArray, got)
				}
				operationStack.Pop()
			}

		}
	}
	outputArray = append(outputArray, numberString)
	for operationStack.Len() != 0 {
		got := fmt.Sprintf("%v", operationStack.Pop())
		outputArray = append(outputArray, got)
	}

	// fmt.Println(outputArray)
	for i := range outputArray {
		fmt.Print(outputArray[i]," , ")
	}
}
