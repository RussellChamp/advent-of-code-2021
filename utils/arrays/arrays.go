package arrays

func FindFirst(list []string, filter func(string) bool) string {
	for _, item := range list {
		if filter(item) {
			return item
		}
	}
	panic("No item found that matches criteria")
}

func FindFirstInt(list []int, filter func(int) bool) int {
	for _, item := range list {
		if filter(item) {
			return item
		}
	}
	panic("No item found that matches criteria")
}

func FindAll(list []string, filter func(string) bool) []string {
	var filteredList []string
	for _, item := range list {
		if filter(item) {
			filteredList = append(filteredList, item)
		}
	}

	return filteredList
}

func FindAllInt(list []int, filter func(int) bool) []int {
	var filteredList []int
	for _, item := range list {
		if filter(item) {
			filteredList = append(filteredList, item)
		}
	}

	return filteredList
}

func MapStrToInt(input []string, fn func(string) int) []int {
	var retVals []int
	for _, s := range input {
		retVals = append(retVals, fn(s))
	}
	return retVals
}

func Reverse(input []string) []string {
	var output []string
	for idx := len(input) - 1; idx >= 0; idx-- {
		output = append(output, input[idx])
	}
	return output
}
