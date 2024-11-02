package circular_sentence

import "strings"

func isCircularSentence(sentence string) bool {
	arr := strings.Split(sentence, " ")
	if len(arr) == 1 {
		if sentence[0] != sentence[len(sentence)-1] {
			return false
		}
		return true
	}
	firstLetter := sentence[0]
	for i := 1; i < len(arr); i++ {
		if arr[i-1][len(arr[i-1])-1] != arr[i][0] {
			return false
		}
	}
	if arr[len(arr)-1][len(arr[len(arr)-1])-1] != firstLetter {
		return false
	}
	return true
}

func Solve(s string) bool {
	return isCircularSentence(s)
}
