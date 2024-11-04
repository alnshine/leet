package merge_strings_alternately

func mergeAlternately(word1 string, word2 string) string {
	var res string
	var i int
	for i < len(word1) && i < len(word2) {
		res += string(word1[i]) + string(word2[i])
		i++
	}
	if len(word1) < len(word2) {
		res += word2[i:]
	} else if len(word2) < len(word1) {
		res += word1[i:]
	}
	return res
}

func Solve(s1, s2 string) string {
	return mergeAlternately(s1, s2)
}
