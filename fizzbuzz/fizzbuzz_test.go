package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		name string
		in   int
		want string
	}{
		{"returns the number when not divisible by 3 or 5", 1, "1"},
		{"Fizz on multiples of 3", 6, "Fizz"},
		{"Buzz on multiples of 5", 10, "Buzz"},
		{"FizzBuzz on multiples of 15", 30, "FizzBuzz"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FizzBuzz(tt.in); got != tt.want {
				t.Errorf("FizzBuzz(%d) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
