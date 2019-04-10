package gorae

import (
	"fmt"
	"testing"
)

/*
Word are composed of a 7 char ~random ID: 002rZ9U
and can be as short as 1 character: "a"
*/
func isAValidWord(word RaeWord) bool {
	if len(word.ID) < 5 || len(word.Header) < 1 {
		return false
	}
	return true
}

func TestWordOfTheDay(t *testing.T) {
	wot := WordOfTheDay()

	if isAValidWord(wot) {
		fmt.Println("Today's word: " + wot.ID + ", " + wot.Header)
	} else {
		t.Error("Didn't get a valid key for the Word of the Day: " +
			wot.ID + " " + wot.Header)
	}

}

func TestSingleResultApproximateWord(t *testing.T) {
	words, _ := SearchWords("enfasi")
	gotTheWord := false

	if len(words.Res) != 1 {
		t.Error("Didn't get the single word back from the RAE")
		return
	}

	if words.Res[0].Header == "énfasis." {
		gotTheWord = true
	}

	if !gotTheWord {
		t.Error("Didn't the the word 'énfasis.' back from the RAE")
	}

}

func TestMultipleResultApproximateWord(t *testing.T) {
	words, _ := SearchWords("a") // This returs 4 results

	if len(words.Res) < 3 {
		t.Errorf("'a' only returned %d words", len(words.Res))
	}
}

func TestFailedApproximateWordSearch(t *testing.T) {
	words, _ := SearchWords("aoeui")
	if len(words.Res) != 0 {
		t.Error("Searching 'aoeui' shouldn't return any word o_O")
	}
}

func TestExactWordSearchFail(t *testing.T) {
	words := searchExactWord("enfasis") // accent

	if len(words.Res) != 0 {
		t.Errorf("'enfasis' returned %d words", len(words.Res))
	}
}

func TestExactWordSearch(t *testing.T) {
	words := searchExactWord("énfasis")

	if len(words.Res) == 0 {
		t.Errorf("'énfasis' returned %d words", len(words.Res))
	}
}
