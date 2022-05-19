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
)

func StringSum(input string) (output string, err error) {
	input = strings.ReplaceAll(input, " ", "")

	if len(input) == 0 {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}
	if Handling(input) != 2 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}

	input1 := strings.LastIndex(input, "+")
	if input1 == -1 {
		input1 = strings.LastIndex(input, "-")
	}

	x, err := strconv.Atoi(strings.TrimLeft(input[0:input1], "+"))
	if err != nil {
		return "", fmt.Errorf("incorrect input: %w", input[0:input1], err)
	}
	y, err := strconv.Atoi(strings.TrimLeft(input[input1:], "+"))
	if err != nil {
		return "", fmt.Errorf("incorrect input: %w", input[input1:], err)
	}

	return strconv.Itoa(x + y), nil

}

func Handling(input string) (count int) {
	input = strings.TrimLeft(input, "+")
	input = strings.TrimLeft(input, "-")

	for _, a := range strings.Split(input, "-") {
		count += len(strings.Split(a, "+"))
	}
	return
}
