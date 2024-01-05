package main

type Master struct {
	secret string
}

func NewMaster(secret string) *Master {
	return &Master{secret: secret}
}

func (m *Master) Guess(word string) int {
	return countMatches(word, m.secret)
}

func findSecretWord(wordlist []string, master *Master) {
	ln := len(wordlist)
	for i := 0; i < 10; i++ {
		nMatches := master.Guess(wordlist[0])
		if nMatches == 6 {
			return
		}
		ln = reorganise(wordlist[0], wordlist, ln, nMatches)
	}
}

func reorganise(matchWord string, wordList []string, ln, nMatches int) int {
	i := 0
	for i < ln {
		if countMatches(matchWord, wordList[i]) != nMatches {
			ln--
			wordList[i] = wordList[ln]
		} else {
			i++
		}
	}
	return ln
}

func countMatches(word1, word2 string) int {
	count := 0
	for i := range word1 {
		if word1[i] == word2[i] {
			count++
		}
	}

	return count
}

func main() {
	elements := []string{"acckzz", "ccbazz", "eiowzz", "abcczz"}
	findSecretWord(elements, NewMaster("acckzz"))
}
