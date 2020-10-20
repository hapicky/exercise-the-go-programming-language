package mapintset

import (
	"sort"
	"strconv"
	"strings"
)

type MapIntSet struct {
	vals map[int]bool
}

func (s *MapIntSet) Has(x int) bool {
	return s.vals[x]
}

func (s *MapIntSet) Add(x int) {
	if s.vals == nil {
		s.vals = map[int]bool{}
	}
	s.vals[x] = true
}

func (s *MapIntSet) String() string {
	var vals []int
	for i, _ := range s.vals {
		vals = append(vals, i)
	}
	sort.Ints(vals)

	var builder strings.Builder
	builder.WriteString("{")
	for i, v := range vals {
		if i > 0 {
			builder.WriteString(" ")
		}
		builder.WriteString(strconv.Itoa(v))
	}
	builder.WriteString("}")
	return builder.String()
}

func (s *MapIntSet) UnionWith(t *MapIntSet) {
	for i, _ := range t.vals {
		s.Add(i)
	}
}
