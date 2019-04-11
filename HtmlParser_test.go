package gorae

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func genericHtmlToText(filename string, t *testing.T) bool {
	htmbin, err := ioutil.ReadFile(filename + ".html")

	if err != nil {
		t.Error("Couldn't open the test file")
		t.Error(err)
		return false
	}

	text := HtmlToText(string(htmbin))

	// compute the md5 of the result and refference file
	textSum := md5.Sum([]byte(text))

	refFileBin, err := ioutil.ReadFile(filename + ".txt")
	if err != nil {
		t.Error("Didn't find refference file")
		t.Error(err)
		return false
	}

	refFileSum := md5.Sum(refFileBin)

	if textSum != refFileSum {
		t.Error("sum of the reference and resulting text doesn't match")
		t.Errorf("Got: %x  Expected: %x", textSum, refFileSum)
		fmt.Println(text)
		fmt.Printf("%x", text)
	}

	return true
}

func TestHtmlToText(t *testing.T) {
	if genericHtmlToText("./testdata/zas", t) == false {
		t.Error("The simple HTML to Text test failed")
	}

	if genericHtmlToText("./testdata/amor", t) == false {
		t.Error("The complex HTML to Text test failed")
	}

}

func TestParseHeader(t *testing.T) {
	HTMLheader := `<header class="f">el, la</header>`
	tokenizer := html.NewTokenizer(strings.NewReader(HTMLheader))
	tokenizer.Next()

	header := parseHeader(tokenizer)
	if header != "*el, la*" {
		t.Error("Header doesn't match: " + header)
	}
}
