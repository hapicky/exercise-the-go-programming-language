package palindrome

import (
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	l := s.Len()
	if l == 0 {
		return false
	}

	j := l - 1
	for i := 0; i <= j; i++ {
		if !(!s.Less(i, j) && !s.Less(j, i)) {
			return false
		}
		j--
	}

	return true
}

type StringSlice []string

func (s StringSlice) Len() int {
	return len(s)
}

func (s StringSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s StringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
