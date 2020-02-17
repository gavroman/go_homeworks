package main

import (
    "errors"
    "flag"
    "fmt"
    Stack "github.com/golang-collections/collections/stack"
    "strconv"
)

func isDigit(char int32) bool {
    return char >= '0' && char <= '9'

}

func reversePolishNotation(expr string) ([]string, error) {
    resultArray := make([]string, 0, len(expr))
    stack := Stack.New()
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
                for stack.Len() != 0 && stack.Peek() != "(" {
                    got := fmt.Sprintf("%v", stack.Pop())
                    resultArray = append(resultArray, got)
                }
                stack.Push(string(char))

            case '*', '/':
                if stack.Peek() == "*" || stack.Peek() == "/" {
                    for stack.Len() != 0 && stack.Peek() != "+" && stack.Peek() != "-" {
                        got := fmt.Sprintf("%v", stack.Pop())
                        resultArray = append(resultArray, got)
                    }
                }
                stack.Push(string(char))

            case '(':
                stack.Push(string(char))

            case ')':
                for stack.Len() != 0 && stack.Peek() != "(" {
                    got := fmt.Sprintf("%v", stack.Pop())
                    resultArray = append(resultArray, got)
                }
                stack.Pop()
            default:
                return []string{}, errors.New("wrong character in expression")
            }
        }
    }
    if numbers != "" {
        resultArray = append(resultArray, numbers)
    }
    for stack.Len() != 0 {
        got := fmt.Sprintf("%v", stack.Pop())
        resultArray = append(resultArray, got)
    }
    return resultArray, nil
}

func calcFromPolishNotation(exprArray []string) (float64, error) {
    stack := Stack.New()
    for i := range exprArray {
        number, err := strconv.Atoi(exprArray[i])
        if err == nil {
            stack.Push(number)
        } else {
            right, err := strconv.ParseFloat(fmt.Sprintf("%v", stack.Pop()), 64)
            left, err := strconv.ParseFloat(fmt.Sprintf("%v", stack.Pop()), 64)
            if err != nil {
                return 0, err
            }
            switch exprArray[i] {
            case "+":
                stack.Push(left + right)
            case "-":
                stack.Push(left - right)
            case "*":
                stack.Push(left * right)
            case "/":
                stack.Push(left / right)
            default:
                return 0, errors.New("wrong operation character")
            }
        }
    }
    result, err := strconv.ParseFloat(fmt.Sprintf("%v", stack.Pop()), 64)
    return result, err
}

func main() {
    flag.Parse()
    args := flag.Args()

    expressionString := ""
    for i := range args {
        expressionString += args[i]
    }

    outputArray, err := reversePolishNotation(expressionString)
    if err != nil {
        return
    }
    fmt.Println(outputArray)
    result, err := calcFromPolishNotation(outputArray)
    if err == nil {
        fmt.Println(result)
    }
}
