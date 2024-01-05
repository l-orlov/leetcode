package reverse_prefix_of_word

import "fmt"

func iterate1(s string) {
	for i := 0; i < len(s); i++ {
		i1 := len(s)
		i2 := len(s)
		i3 := len(s)
		_ = i1
		_ = i2
		_ = i3
	}
}
func iterate2(s string) {
	n := len(s)
	for i := 0; i < n; i++ {
		i1 := n
		i2 := n
		i3 := n
		_ = i1
		_ = i2
		_ = i3
	}
}

func reverseStr(input string) string {
	res := []rune(input)
	n := len(res)
	for i := 0; i < n/2; i++ {
		res[i], res[n-1-i] = res[n-1-i], res[i]
	}
	return string(res)
}

func ReverseString() {
	input := "The quick brown 狐 jumped over the lazy 犬"
	fmt.Println(reverseStr(input))
}

func reversePrefix(word string, ch byte) string {
	res := []byte(word)
	for i := 0; i < len(res); i++ {
		if res[i] == ch {
			n := i + 1
			for j := 0; j < n/2; j++ {
				res[j], res[n-1-j] = res[n-1-j], res[j]
			}
			break
		}
	}
	return string(res)
}
