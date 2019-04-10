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
