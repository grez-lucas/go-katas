package stringcalculator

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
	return 0, nil
}
