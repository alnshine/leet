package delete_characters_to_make_fancy_string

func makeFancyString(s string) string {
	lenS := len(s)
	res := make([]byte, 0, lenS)
	cur := s[0]
	var c int
	for i := 0; i < lenS; i++ {
		if cur == s[i] {
			c++
		} else {
			cur = s[i]
			c = 1
		}
		if c < 3 {
			res = append(res, cur)
		}
	}
	return string(res)
}

func Solve(s string) string {
	return makeFancyString(s)
}
