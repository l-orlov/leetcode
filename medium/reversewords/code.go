package reversewords

import "strings"

func reverseWords(s string) string {
	res := strings.Builder{}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			if res.Len() > 0 {
				res.WriteByte(' ')
			}

			j := i - 1
			for ; j >= 0 && s[j] != ' '; j-- {
			}
			res.WriteString(s[j+1 : i+1])
			i = j
		}
	}

	return res.String()
}
