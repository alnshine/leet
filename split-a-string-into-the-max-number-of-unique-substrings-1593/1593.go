package split_a_string_into_the_max_number_of_unique_substrings_1593

func maxUniqueSplit(s string) int {
	m := make(map[string]struct{})

	var backtrack func(start int) int
	backtrack = func(start int) int {
		if start == len(s) {
			return 0
		}

		maxCount := 0
		for i := start + 1; i <= len(s); i++ {
			substr := s[start:i]
			if _, ok := m[substr]; !ok {
				m[substr] = struct{}{}
				count := 1 + backtrack(i)
				if count > maxCount {
					maxCount = count
				}
				delete(m, substr)
			}
		}
		return maxCount
	}

	return backtrack(0)
}

func Solve(s string) int {
	return maxUniqueSplit(s)
}
