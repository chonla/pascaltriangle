package pascaltriangle

func PascalLineRecursive(line int) []int {
	if line == 1 {
		return []int{1}
	}
	if line == 2 {
		return []int{1, 1}
	}

	prevLine := PascalLineRecursive(line - 1)
	result := []int{1}

	for i, n := 0, len(prevLine)-1; i < n; i++ {
		result = append(result, prevLine[i]+prevLine[i+1])
	}

	result = append(result, 1)

	return result
}
