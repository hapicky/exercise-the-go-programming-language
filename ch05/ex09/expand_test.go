package expand

import (
	"testing"
)

func TestExpand(t *testing.T) {

	hello := func(s string) string {
		return "hello " + s + "."
	}

	var tests = []struct {
		s    string
		f    func(string) string
		want string
	}{
		{
			"$foo",
			hello,
			"hello foo.",
		},
		{
			"$foo $bar",
			hello,
			"hello foo. hello bar.",
		},
		{
			"noop",
			hello,
			"noop",
		},
	}

	for _, test := range tests {
		got := expand(test.s, test.f)
		if got != test.want {
			t.Errorf(`expected: %v, but %v`, test.want, got)
		}
	}
}
