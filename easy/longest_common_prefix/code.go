package longest_common_prefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		j := 0
		for ; j < len(strs[i]) && j < len(prefix) && strs[i][j] == prefix[j]; j++ {
		}
		if j == 0 {
			return ""
		}
		prefix = prefix[:j]
	}

	return prefix
}
