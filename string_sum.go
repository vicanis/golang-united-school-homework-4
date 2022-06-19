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
	var bs, firstNumber, secondNumber []byte

	for _, c := range input {
		if !is(byte(c), "whitespace") {
			bs = append(bs, byte(c))
		}
	}

	if len(bs) == 0 {
		return "", errorEmptyInput
	}

	var i int

	if bs[0] == '-' {
		firstNumber = append(firstNumber, '-')
		i++
	}

	for {
		if is(bs[i], "digit") {
			firstNumber = append(firstNumber, bs[i])
		} else {
			break
		}

		if i >= len(bs)-2 {
			return "", errorNotTwoOperands
		}

		i++
	}

	if !is(bs[i], "operator") {
		return "", fmt.Errorf("no operator ('%c')", bs[i])
	}

	operator := bs[i]

	i++

	for {
		if is(bs[i], "digit") {
			secondNumber = append(secondNumber, bs[i])
		} else {
			break
		}

		if i == len(bs)-1 {
			break
		}

		i++
	}

	firstInt, err := strconv.Atoi(string(firstNumber))
	if err != nil {
		return "", err
	}

	secondInt, err := strconv.Atoi(string(secondNumber))
	if err != nil {
		return "", err
	}

	var resultInt int

	switch operator {
	case '-':
		resultInt = firstInt - secondInt
	case '+':
		resultInt = firstInt + secondInt
	}

	return fmt.Sprintf("%d", resultInt), nil
}
