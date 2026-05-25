package stringcalculator

import "testing"

// Rule 1 — start here. The empty-string case already passes with the stub;
// the two-numbers case is your first real RED. Make it green, then keep going
// through the rules in string_calculator.go, adding a test for each.

func TestAdd_emptyStringReturnsZero(t *testing.T) {
	got, err := Add("")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 0 {
		t.Errorf(`Add("") = %d, want 0`, got)
	}
}

func TestAdd_twoNumbers(t *testing.T) {
	got, err := Add("1,2")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 3 {
		t.Errorf(`Add("1,2") = %d, want 3`, got)
	}
}
