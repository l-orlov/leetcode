package roman_to_integer

var symbolVals = map[uint8]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

// I can be placed before V (5) and X (10) to make 4 and 9.
// X can be placed before L (50) and C (100) to make 40 and 90.
// C can be placed before D (500) and M (1000) to make 400 and 900.

// Constraints:
// 1) 1 <= s.length <= 15
// 2) s contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').
// 3) It is guaranteed that s is a valid roman numeral in the range [1, 3999].
func romanToInt(s string) int {
	if len(s) == 0 {
		return 0
	}

	result := 0
	var i int
	for i = 0; i < len(s)-1; i++ {
		tmp := symbolVals[s[i]]
		if tmp < symbolVals[s[i+1]] {
			result -= tmp
		} else {
			result += tmp
		}
	}
	if i < len(s) {
		result += symbolVals[s[i]]
	}

	return result
}
