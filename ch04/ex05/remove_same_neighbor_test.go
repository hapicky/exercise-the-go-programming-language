package remove_same_neighbor

import (
	"reflect"
	"testing"
)

func TestRemoveSameNeighbor(t *testing.T) {
	var tests = []struct {
		input []string
		want  []string
	}{
		{
			[]string{"abc", "abc", "def", "ghi", "ghi", "abc"},
			[]string{"abc", "def", "ghi", "abc"},
		},
	}

	for _, test := range tests {
		got := RemoveSameNeighbor(test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf(`expected: %v, but %v`, test.want, got)
		}
	}
}
