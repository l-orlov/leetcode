package ransom_note

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	letterCounts := make([]int, 26)
	for _, letter := range magazine {
		letterCounts[letter-'a']++
	}

	for _, letter := range ransomNote {
		index := letter - 'a'
		if letterCounts[index] == 0 {
			return false
		}
		letterCounts[index]--
	}

	return true
}
