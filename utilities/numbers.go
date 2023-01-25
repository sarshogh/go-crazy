package utilities

func IsEqual(numbers1, numbers2 []int) bool {
	if len(numbers1) != len(numbers2) {
		return false
	}

	for i, v := range numbers1 {
		if v != numbers2[i] {
			return false
		}
	}
	return true
}
