package wordy

import (
	"strconv"
	"strings"
)

const testVersion = 1

// Answer parses a word problem to return the mathematical result
func Answer(input string) (int, bool) {
	// q holds the list of operands and operators
	q := []string{}

	// break apart word problem into operands and operators
	q, valid := parseWordProblem(input)
	if len(q) == 0 || !valid {
		return 0, false
	}

	// apply operators to operands to get the result
	result := applyOperators(q)
	return result, true
}

// helper function to parse a word problem and put operands/operators into an array
func parseWordProblem(input string) ([]string, bool) {
	q := []string{}
	words := strings.Split(input, " ")
	reachedFirstOperand := false
	for i, word := range words {
		// handle end of sentence question mark
		if i == len(words)-1 {
			word = word[0 : len(word)-1]
		}
		_, err := strconv.Atoi(word)
		if err == nil {
			// this is an operand
			if !reachedFirstOperand {
				// we are past the "What is..." part, set flag to true
				reachedFirstOperand = true
			}
			// add operand to q
			q = append(q, word)
		} else if err != nil && reachedFirstOperand {
			// this should be an operator, add to q
			// check for valid operators
			if word != "divided" && word != "multiplied" && word != "plus" && word != "minus" && word != "by" {
				return q, false
			}
			// handle multiple words - e.g. divided by, multiplied by
			if word != "by" {
				if (word == "divided" || word == "multiplied") && (len(words)-1 > i) && (words[i+1] == "by") {
					word += " by"
				}
				q = append(q, word)
			}
		}
	}
	return q, true
}

// given a list of operands and operators, apply operators to operands in a left to right fashion
func applyOperators(q []string) int {
	operand := 0
	operator := ""
	operandCount := 0
	result := 0
	for _, word := range q {
		i, err := strconv.Atoi(word)
		if err == nil {
			operandCount++
			if operandCount >= 2 {
				// we have two numbers, apply last operator
				result = applyOperator(operand, i, operator)
			}
			if result != 0 {
				// operand becomes the result of the last operation
				operand = result
			} else {
				// operand becomes previous operand
				operand = i
			}
		} else {
			operator = word
		}
	}
	return result
}

// helper function to apply a string operator to two operands
func applyOperator(operand1 int, operand2 int, operator string) int {
	var result int
	switch operator {
	case "plus":
		result = operand1 + operand2
	case "minus":
		result = operand1 - operand2
	case "divided by":
		result = operand1 / operand2
	case "multiplied by":
		result = operand1 * operand2
	}
	return result
}
