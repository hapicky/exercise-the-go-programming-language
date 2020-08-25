package palindrome

import (
	"strings"
	"testing"
)

func TestPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"", false},
		{"dad", true},
		{"noon", true},
		{"banana", false},
		{"was it a car or a cat i saw", true},
	}

	for _, test := range tests {
		ss := StringSlice(strings.Split(strings.ReplaceAll(test.s, " ", ""), ""))
		got := IsPalindrome(ss)
		if got != test.want {
			t.Errorf("IsPalindrome(%q) got %v, want %v\n",
				test.s, got, test.want)
		}
	}
}
