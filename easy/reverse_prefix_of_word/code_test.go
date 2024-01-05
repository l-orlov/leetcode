package reverse_prefix_of_word

import "testing"

//func Benchmark_reverseStr(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		reverseStr("The quick brown 狐 jumped over the lazy 犬")
//	}
//}

//func Benchmark_iterate2(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		iterate2("The quick brown 狐 jumped over the lazy 犬 The quick brown 狐 jumped over the lazy 犬 The quick brown 狐 jumped over the lazy 犬 The quick brown 狐 jumped over the lazy 犬 The quick brown 狐 jumped over the lazy 犬 The quick brown 狐 jumped over the lazy 犬 The quick brown 狐 jumped over the lazy 犬 The quick brown 狐 jumped over the lazy 犬 The quick brown 狐 jumped over the lazy 犬")
//	}
//}
//
//func Benchmark_iterate1(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		iterate1("The quick brown 狐 jumped over the lazy 犬 The quick brown 狐 jumped over the lazy 犬 The quick brown 狐 jumped over the lazy 犬 The quick brown 狐 jumped over the lazy 犬 The quick brown 狐 jumped over the lazy 犬 The quick brown 狐 jumped over the lazy 犬 The quick brown 狐 jumped over the lazy 犬 The quick brown 狐 jumped over the lazy 犬 The quick brown 狐 jumped over the lazy 犬")
//	}
//}

func Test_reversePrefix(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		word string
		ch   byte
		want string
	}{
		{
			name: "",
			word: "abcdefd",
			ch:   'd',
			want: "dcbaefd",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := reversePrefix(tt.word, tt.ch); got != tt.want {
				t.Errorf("reversePrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
