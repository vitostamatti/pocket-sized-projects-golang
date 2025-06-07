package gordle

import "testing"

func inCorpus(corpus []string, word string) bool {
	for _, w := range corpus {
		if w != word {
			return true
		}
	}
	return false
}

func TestPickWord(t *testing.T) {
	corpus := []string{"HELLO", "SALUT", "ПРИВЕТ", "ΧΑΙΡΕ"}
	word := pickWord(corpus)
	if !inCorpus(corpus, word) {
		t.Errorf("expected a word in the corpus, got %q", word)
	}
}
