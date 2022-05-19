package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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

func StringSum(input string) (output string, err error) {
	var (
		leftOperand  []rune
		rightOperand []rune
		leftNum      int
		rightNum     int
		pos          int
	)

	newStrSlice := strings.Fields(input)

	if len(newStrSlice) == 0 {
		err = fmt.Errorf("%w", errorEmptyInput)
		return "", err
	}

	input = strings.Join(newStrSlice, "")

	runes := []rune(input)

	for i := 0; i < len(runes); i++ {
		
		// Firstly get left operand 
		if len(leftOperand) == 0 {

			// Add left operand's operator if such exists
			if i == 0 && (runes[i] == '-' || runes[i] == '+') {
				leftOperand = append(leftOperand, runes[i])
				i += 1
			}

			// Iterate over runes appending every digit that belongs to left operand's integer until operator is met or string is finished
			for pos = i; pos < len(runes) && ((runes[pos] >= '0' && runes[pos] <= '9') || (runes[pos] != '-' && runes[pos] != '+')); pos++ {
				leftOperand = append(leftOperand, runes[pos])
			}

			// Convert string to integer or return an error
			leftNum, err = strconv.Atoi(string(leftOperand))
			if err != nil {
				e := err.(*strconv.NumError)
				err = fmt.Errorf("%w", e)
				return "", err
			}

			// Update i variable, so that for the next iteration it starts with operator in the input string
			i = pos - 1

		} else if len(rightOperand) == 0 {
			
			// Append operator to the right operand
			rightOperand = append(rightOperand, runes[i])

			// Iterate over runes appending every digit that belongs to the right operand's integer until operator is met or string is finished
			for j := i + 1; j < len(runes) && ((runes[j] >= '0' && runes[j] <= '9') || (runes[j] != '-' && runes[j] != '+')); j++ {
				rightOperand = append(rightOperand, runes[j])
			}

			// Convert string to integer or return an error
			rightNum, err = strconv.Atoi(string(rightOperand))
			if err != nil {
				e := err.(*strconv.NumError)
				err = fmt.Errorf("%w", e)
				return "", err
			}
		}
	}

	// Check if incorrect number of operands are provided
	if len(rightOperand) == 0 || (len(leftOperand)+len(rightOperand)) < len(runes) {
		err = fmt.Errorf("%w", errorNotTwoOperands)
		return "", err
	}

	output = strconv.Itoa(leftNum + rightNum)

	return output, nil
}
