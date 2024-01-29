package recover_password

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"testing"
)

var alphabet = []rune{'a', 'b', 'c', 'd', '1', '2', '3'}

func RecoverPassword(h []byte) string {
	var step int

	for ; ; step++ {
		guess := genPassword(step)
		if bytes.Equal(hashPassword(guess), h) {
			return guess
		}
	}
}

func genPassword(step int) (res string) {
	for {
		res = res + string(alphabet[step%len(alphabet)])
		step = step/len(alphabet) - 1
		if step < 0 {
			break
		}
	}

	return
}

func TestRecoverPassword(t *testing.T) {
	for _, exp := range []string{
		"a",
		"12",
		"abc333d",
	} {
		t.Run(exp, func(t *testing.T) {
			act := RecoverPassword(hashPassword(exp))
			if act != exp {
				t.Error("recovered:", act, "expected:", exp)
			}
		})
	}
}

func Test_genPassword(t *testing.T) {
	for _, step := range []int{
		0,
		1,
		2,
		7,
		8,
		10,
		11,
		12,
	} {
		act := genPassword(step)
		fmt.Println(act)
	}
}

func hashPassword(in string) []byte {
	h := md5.Sum([]byte(in))
	return h[:]
}
