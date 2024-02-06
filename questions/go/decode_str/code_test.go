package decode_str

import "testing"

func Test_decodeStr(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "1",
			s:    "abcd",
			want: "abcd",
		},
		{
			name: "2",
			s:    "3[a]2[bc]",
			want: "aaabcbc",
		},
		{
			name: "3",
			s:    "2[abc]3[cd]ef",
			want: "abcabccdcdcdef",
		},
		{
			name: "3",
			s:    "11[a]",
			want: "aaaaaaaaaaa",
		},
		{
			name: "3",
			s:    "1[a2[b3[c]]]",
			want: "abcccbccc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeStr(tt.s); got != tt.want {
				t.Errorf("decodeStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
