package string_sum

import (
	"errors"
	"fmt"
	"strconv"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func includes(value byte, arr []byte) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}

	return false
}

func is(c byte, t string) bool {
	switch t {
	case "digit":
		return includes(c, []byte("0123456789"))
	case "whitespace":
		return includes(c, []byte(" \r\n\t"))
	case "operator":
		return includes(c, []byte("-+"))
	}

	return false
}

func StringSum(input string) (string, error) {
	var inputBytes []byte

	for _, c := range input {
		if !is(byte(c), "whitespace") {
			inputBytes = append(inputBytes, byte(c))
		}
	}

	if len(inputBytes) == 0 {
		return "", fmt.Errorf("parse failed: %w", errorEmptyInput)
	}

	var i, j int

	for {
		if i < len(inputBytes) && (is(inputBytes[i], "digit") || i == 0 && inputBytes[i] == '-') {
			i++
		} else {
			break
		}
	}

	firstNumber, err := strconv.Atoi(string(inputBytes[:i]))
	if err != nil {
		return "", fmt.Errorf("first number parse failed: %w", err)
	}

	if i >= len(inputBytes)-1 {
		return "", fmt.Errorf("%w: no second operand", errorNotTwoOperands)
	}

	operator := inputBytes[i]

	j = i + 1

	for {
		if j < len(inputBytes) && is(inputBytes[j], "digit") {
			j++
		} else {
			break
		}
	}

	if j < len(inputBytes)-1 {
		return "", fmt.Errorf("%w: too many operands", errorNotTwoOperands)
	}

	secondNumber, err := strconv.Atoi(string(inputBytes[i+1 : j]))
	if err != nil {
		return "", fmt.Errorf("second number parse failed: %w", err)
	}

	switch operator {
	case '-':
		return fmt.Sprintf("%d", firstNumber-secondNumber), nil
	case '+':
		return fmt.Sprintf("%d", firstNumber+secondNumber), nil
	}

	return "", fmt.Errorf("%w: unknown operator", errorNotTwoOperands)
}
