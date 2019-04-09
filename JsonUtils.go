package gorae

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

/*
The RAE returns a JSON string that begins with "json(" and ends with ")"
*/
func getJSONFromBody(response *http.Response) (jsonstr string) {
	var begin int

	data, _ := ioutil.ReadAll(response.Body)
	// the data is in binary, we need to convert it to string
	text := string(data)
	fmt.Println(text)

	if strings.Index(text, "json(") >= 0 {
		begin = 5
	} else {
		log.Fatal("not a json string: ", text)
	}

	jsonstr = text[begin : len(text)-1]

	return
}
