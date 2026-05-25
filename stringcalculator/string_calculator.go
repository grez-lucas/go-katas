package stringcalculator

import (
	"strconv"
	"strings"
)

// Add takes a string of separated numbers and returns their sum.
//
// This is a TDD exercise — implement the rules below one at a time. For each
// rule, add/uncomment a test, watch it fail (red), make it pass (green), then
// refactor. Don't jump ahead; let the tests drive the code.
//
//  1. ""  -> 0,  "1" -> 1,  "1,2" -> 3
//  2. Handle an arbitrary amount of numbers.
//  3. Allow newlines as separators too:  "1\n2,3" -> 6
//  4. Support a custom delimiter:  "//;\n1;2" -> 3
//  5. Negative numbers return an error naming the negatives.
//
// The stub below makes the package compile so the first test can run and fail.
func Add(numbers string) (int, error) {
	switch numbers {
	case "":
		return 0, nil
	default:
		// Check if we have a custom delimeter
		var customDelimeter string
		if strings.HasPrefix(numbers, "//") {
			// Split the custom delimeter from the header
			before, after, _ := strings.Cut(numbers, "\n")
			customDelimeter = strings.TrimPrefix(before, "//")
			numbers = after
		}

		sum := 0
		delimeters := "\n" + "," + customDelimeter
		parts := strings.FieldsFunc(numbers, func(r rune) bool {
			return strings.ContainsRune(delimeters, r)
		})
		for _, part := range parts {
			val, err := strconv.Atoi(part)
			if err != nil {
				return 0, err
			}
			sum += val
		}

		return sum, nil
	}
}
