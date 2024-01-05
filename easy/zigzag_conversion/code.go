package zigzag_conversion

func convert(s string, numRows int) string {
	rows := make([][]byte, numRows)

	idx := 0
	for idx < len(s) {
		for i := 0; i < numRows && idx < len(s); i++ {
			rows[i] = append(rows[i], s[idx])
			idx++
		}
		for j := numRows - 2; 0 < j && idx < len(s); j-- {
			rows[j] = append(rows[j], s[idx])
			idx++
		}
	}

	var zigzag []byte
	for _, r := range rows {
		zigzag = append(zigzag, r...)
	}

	return string(zigzag)
}
