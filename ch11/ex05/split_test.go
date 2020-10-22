package ex05

import (
	"reflect"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	var tests = []struct {
		s    string
		sep  string
		want []string
	}{
		{"a:b:c", ":", []string{"a", "b", "c"}},
	}

	for _, test := range tests {
		got := strings.Split(test.s, test.sep)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Split(%q, %q) returned %v, want %v",
				test.s, test.sep, got, test.want)
		}
	}
}
