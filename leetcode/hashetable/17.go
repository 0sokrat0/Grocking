package hashetable

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	phoneMap := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var output []string
	var backtrack func(combo string, nextDigits string)

	backtrack = func(combo, nextDigits string) {

		if nextDigits == "" {
			output = append(output, combo)
			return
		} else {
			letter := phoneMap[nextDigits[0]-'2']
			for _, l := range letter {
				backtrack(combo+string(l), nextDigits[1:])
			}
		}
	}
	backtrack("", digits)
	return output
}
