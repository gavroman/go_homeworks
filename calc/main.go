package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"strconv"
)

func isDigit(char int32) bool {
	return char >= '0' && char <= '9'
}

func validateBrackets(expr string) bool {
	counter := 0
	for _, char := range expr {
		switch char {
		case '(':
			counter++
		case ')':
			counter--
		}
		if counter < 0 {
			return false
		}
	}
	if counter != 0 {
		return false
	}
	return true
}

func calc(expr string) (float64, error) {
	resultArray := make([]string, 0, len(expr))
	stack := New()
	numbers := ""
	for _, char := range expr {
		if isDigit(char) {
			numbers += string(char)
		} else {
			if numbers != "" {
				resultArray = append(resultArray, numbers)
				numbers = ""
			}
			switch char {
			case '+', '-':
				for stack.Len() != 0 && *stack.Peek() != "(" {
					resultArray = append(resultArray, *stack.Pop())
				}
				stack.Push(string(char))

			case '*', '/':
				for stack.Len() != 0 && *stack.Peek() != "+" && *stack.Peek() != "-" {
					resultArray = append(resultArray, *stack.Pop())
				}
				stack.Push(string(char))

			case '(':
				stack.Push(string(char))

			case ')':
				for stack.Len() != 0 && *stack.Peek() != "(" {
					resultArray = append(resultArray, *stack.Pop())
				}
				stack.Pop()
			default:
				return 0, errors.New("wrong character in expression")
			}
		}
	}
	if numbers != "" {
		resultArray = append(resultArray, numbers)
	}
	for stack.Len() != 0 {
		resultArray = append(resultArray, *stack.Pop())
	}

	for i := range resultArray {
		_, err := strconv.Atoi(resultArray[i])
		if err == nil {
			stack.Push(resultArray[i])
		} else {
			right, err := strconv.ParseFloat(*stack.Pop(), 64)
			left, err := strconv.ParseFloat(*stack.Pop(), 64)
			if err != nil {
				return 0, err
			}
			switch resultArray[i] {
			case "+":
				stack.Push(fmt.Sprintf("%f", left+right))
			case "-":
				stack.Push(fmt.Sprintf("%f", left-right))
			case "*":
				stack.Push(fmt.Sprintf("%f", left*right))
			case "/":
				stack.Push(fmt.Sprintf("%f", left/right))
			default:
				fmt.Println(resultArray[i])
				return 0, errors.New("wrong operation character")
			}
		}
	}
	result, err := strconv.ParseFloat(*stack.Pop(), 64)
	return result, err
}

func main() {
	flag.Parse()
	args := flag.Args()

	expressionString := ""
	for i := range args {
		expressionString += args[i]
	}

	result, err := calc(expressionString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
