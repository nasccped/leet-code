package main

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	mapping := map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	letters := []string{}
	for _, r := range digits {
		letters = append(letters, mapping[r])
	}
	result := []string{}
	var backtrack func(combination string, strList []string)
	backtrack = func(combination string, strList []string) {
		if len(strList) == 0 {
			result = append(result, combination)
			return
		}
		for i := 0; i < len(strList[0]); i++ {
			combination += strList[0][i : i+1]
			backtrack(combination, strList[1:])
			combination = combination[:len(combination)-1]
		}
	}
	backtrack("", letters)
	return result
}
