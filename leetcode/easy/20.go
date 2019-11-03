package easy

func isValid(s string) bool {
	m := make(map[rune]rune)
	m['('] = ')'
	m['{'] = '}'
	m['['] = ']'
	l := make([]rune, 0)
	for _, v := range s {
		if len(l) == 0 {
			l = append(l, v)
			continue
		}

		if _, ok := m[l[len(l)-1]]; !ok {
			return false
		}

		//l = append(l,v)

		if m[l[len(l)-1]] == v {
			l = l[:len(l)-1]
		} else {
			l = append(l, v)
		}

	}

	return len(l) == 0

}
