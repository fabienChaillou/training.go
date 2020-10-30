package hangman

import "testing"

func TestLetterInWord(t *testing.T) {
	word := []string{"b", "o", "d"}
	guess := "b"
	hasLetter := letterInWord(guess, word)
	if !hasLetter {
		t.Errorf("Word %s contains letter %s. got=%v\n", word, guess, hasLetter)
	}
}

func TestLetterNotInWord(t *testing.T) {
	word := []string{"b", "o", "d"}
	guess := "c"
	hasLetter := letterInWord(guess, word)
	if hasLetter {
		t.Errorf("Word %s does not contains letter %s. got=%v\n", word, guess, hasLetter)
	}
}

func TestInvalidWord(t *testing.T) {
	_, err := New(3, "")
	if err == nil {
		t.Errorf("Error should be returned when using an invalid word=''")
	}
}

func TestGameWon(t *testing.T) {
	g, _ := New(3, "bob")
	g.MakeAGuess("b")
	g.MakeAGuess("o")
	g.MakeAGuess("b")
	validState(t, "won", g.State)
}

func TestGameLost(t *testing.T) {
	g, _ := New(3, "bob")
	g.MakeAGuess("a")
	g.MakeAGuess("z")
	g.MakeAGuess("u")
	validState(t, "lost", g.State)
}

func TestGameAlreadyGuess(t *testing.T) {
	g, _ := New(3, "bob")
	g.MakeAGuess("b")
	g.MakeAGuess("b")
	validState(t, "alreadyGuessed", g.State)
}

func TestGameGoodGuess(t *testing.T) {
	g, _ := New(3, "bob")
	g.MakeAGuess("b")
	validState(t, "goodGuess", g.State)
}

func validState(t *testing.T, expectedState, actualState string) bool {
	if expectedState != actualState {
		t.Errorf("State should be '%v'. got=%v", expectedState, actualState)
		return false
	}

	return true
}
