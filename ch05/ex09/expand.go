package expand

import (
	"regexp"
)

func expand(s string, f func(string) string) string {
	re := regexp.MustCompile(`\$\w+`)
	indexes := re.FindAllStringIndex(s, -1)
	result := s
	for i := len(indexes) - 1; i >= 0; i-- {
		from := indexes[i][0]
		to := indexes[i][1]
		before := result[:from]
		after := result[to:]
		result = before + f(s[from+1:to]) + after
	}

	return result
}
