package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)


var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")

	errorIncorrectInput = errors.New("incorrect input")
)

func StringSum(input string) (output string, err error) {
	// if input is empty, has number of operands not equal to two or equals "+" and "-"

	if input == "" || input == " " {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}

	input = strings.TrimSpace(input)

	var input1 = strings.Fields(input)
	if len(input1) != 2 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}

	//summation

	x, myerr1 := strconv.Atoi(input1[0])
	if myerr1 != nil {
		return "", fmt.Errorf("the error at the first operand:%w", myerr1)

	}
	y, myerr2 := strconv.Atoi(input1[1])
	if myerr2 != nil {
		return "", fmt.Errorf("the error at the second operand:%w", myerr2)
	}
	output = strconv.Itoa(x + y)

	//if input has characters

	for _, a := range input {
		if !strings.Contains("0123456789+- ", string(a)) {
			return "", fmt.Errorf("sorry, but: %w", errorIncorrectInput)
		}
	}

	return output, nil
}

func NegNums(input string) []string {
	result := make([]string, 0)

	len := ""
	for _, i := range input {
		if string(input[i]) == "-" || string(input[i]) == "+" {
			len = len + "." + string(input[i])
		} else {
			len = len + string(input[i])
		}
	}
	return result
}
