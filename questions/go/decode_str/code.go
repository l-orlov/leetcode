package decode_str

/*
Дана закодированная строка следующего формата: k[encoded_text]
Здесь k это число повторений строки encoded_text.
Строка гарантированно имеет корректный формат: нет лишних пробелов, скобки всегда правильные и тд.
Необходимо декодировать строку

Пример:

Input:  "3[a]2[bc]"
Output: "aaabcbc"

Input:  "3[a2[c]]"
Output: "accaccacc"

Input:  "2[abc]3[cd]ef"
Output: "abcabccdcdcdef"
*/

func decodeStr(s string) string {
	input := []rune(s)
	output := decodeRunes(input)
	return string(output)
}

func decodeRunes(runes []rune) []rune {
	resultRunes := make([]rune, 0, len(runes))
	var currNum int
	for i := 0; i < len(runes); {
		if runes[i] == ']' {
			return resultRunes
		} else if runes[i] >= '0' && runes[i] <= '9' {
			for runes[i] >= '0' && runes[i] <= '9' {
				currNum = currNum*10 + int(runes[i]-'0')
				i++
			}
			i++
			subRunes := decodeRunes(runes[i:])
			i += len(subRunes)
			for currNum != 0 {
				resultRunes = append(resultRunes, subRunes...)
				currNum--
			}
			if runes[i] == ']' {
				i++
			}
		} else {
			resultRunes = append(resultRunes, runes[i])
			i++
		}
	}
	return resultRunes
}
