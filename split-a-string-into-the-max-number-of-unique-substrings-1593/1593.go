package split_a_string_into_the_max_number_of_unique_substrings_1593

func maxUniqueSplit(s string) int {
	//Здесь создается карта m, в которой будут храниться уникальные подстроки,
	//найденные в процессе поиска.
	m := make(map[string]struct{})

	var backtrack func(start int) int
	backtrack = func(start int) int {
		if start == len(s) {
			return 0
		}

		maxCount := 0
		//Перебираются все возможные подстроки, начиная с позиции start и
		//заканчивая i. substr — это текущая подстрока.
		for i := start + 1; i <= len(s); i++ {
			substr := s[start:i]

			//Если подстрока substr отсутствует в карте m, это означает,
			//что она уникальна и может быть добавлена.
			if _, ok := m[substr]; !ok {

				//Уникальная подстрока добавляется в карту m,
				//и функция backtrack вызывается рекурсивно от следующей позиции i с
				//увеличением счетчика count на 1.
				m[substr] = struct{}{}
				count := 1 + backtrack(i)
				if count > maxCount {
					maxCount = count
				}

				//После возврата из рекурсивного вызова подстрока удаляется из карты,
				//чтобы позволить поиск других разбиений.
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
