/*

https://github.com/GrenderG/uDRAE-sdk


Header authentication
Authorization: Basic cDY4MkpnaFMzOmFHZlVkQ2lFNDM0

Word of the day
https://dle.rae.es/data/wotd?callback=json

Key query
https://dle.rae.es/data/keys?q=hola&callback=jsonp123

Search word
https://dle.rae.es/data/search?w=hola

This returns an id. This id is going to be used in the fetch endpoint.
Fetch word
https://dle.rae.es/data/fetch?id=KYtLWBc

You will need to parse response, because the response contains html tags.
Example

<?php
$handler = curl_init(\"https://dle.rae.es/data/search?w=hola\");
curl_setopt($handler, CURLOPT_HTTPHEADER, array(\"Authorization: Basic cDY4MkpnaFMzOmFHZlVkQ2lFNDM0\"));
curl_setopt($handler, CURLOPT_VERBOSE, false);
curl_setopt($handler, CURLOPT_SSL_VERIFYPEER, false);
curl_setopt($handler, CURLOPT_SSL_VERIFYHOST, false);
$response = curl_exec ($handler);
curl_close($handler);
echo $response
>
*/

package gorae

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	tgb "github.com/go-telegram-bot-api/telegram-bot-api"
)

// newRae initialize a object to call the RESTful API
func newRae(fun RaeFunc, key string) (r rae) {
	r.apiHTTP = raeRestAPI
	r.authHeader = raeRestAuthHeader

	var remoteFunction string
	switch fun {
	case wordDay:
		remoteFunction = "wotd?callback=json"
	case searchword:
		remoteFunction = "search?w=" + key + "&m=30"
	case fetchDefByID:
		remoteFunction = "fetch?id=" + key
	case words:
		remoteFunction = "search?w=" + key
	default:
		log.Fatal("unknown remote function")
	}

	r.req, _ = http.NewRequest("GET", r.apiHTTP+remoteFunction, nil)
	//req.Header.Add("User-Agent", "Diccionario/2 CFNetwork/808.2.16 Darwin/16.3.0")
	//req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.req.Header.Add("Authorization", r.authHeader)

	//fmt.Println(r.apiHTTP + remoteFunction)

	r.client = &http.Client{}

	return
}

//WordOfTheDay ID of today's word
func WordOfTheDay() (word RaeWord) {
	fmt.Println("word of the day")
	r := newRae(wordDay, "")

	resp, _ := r.client.Do(r.req)
	jsonstr := getJSONFromBody(resp)
	fmt.Println("got:" + jsonstr)

	dec := json.NewDecoder(strings.NewReader(jsonstr))

	//json.Unmarshal([]byte(jsonstr), &w)

	if err := dec.Decode(&word); err == io.EOF {
		return RaeWord{"", ""}
	} else if err != nil {
		log.Fatal(err)
	}

	return word
}

//FetchDefinition from the word's ID, get the definition
func FetchDefinition(ID string) (definition string) {
	fmt.Println("fetch definition of " + ID)

	r := newRae(fetchDefByID, ID)

	resp, _ := r.client.Do(r.req)
	data, _ := ioutil.ReadAll(resp.Body)
	definition = string(data)

	return htmlToText(definition)
}

// return ID of exact word
func searchExactWord(word string) (definition string) {
	var res RaeSearchResult
	r := newRae(searchword, word)

	resp, _ := r.client.Do(r.req)
	data, _ := ioutil.ReadAll(resp.Body)
	jsonstr := string(data)

	json.Unmarshal([]byte(jsonstr), &res)

	fmt.Println(jsonstr)
	fmt.Println(res)

	if len(res.Res) == 0 {
		return ""
	}
	return res.Res[0].ID
}

//SearchWords get words similar to "word"'s
func SearchWords(word string) (res RaeSearchResult, opts tgb.InlineKeyboardMarkup) {
	r := newRae(words, word)

	resp, _ := r.client.Do(r.req)
	data, _ := ioutil.ReadAll(resp.Body)
	jsonstr := string(data)

	fmt.Println("json de searchwords", jsonstr)
	json.Unmarshal([]byte(jsonstr), &res)

	if len(res.Res) > 1 {
		replacer := strings.NewReplacer("<sup>", "", "</sup>", "")
		var rows []tgb.InlineKeyboardButton
		for k, palabra := range res.Res {
			if k > 3 {
				break
			}
			pa := tgb.InlineKeyboardButton{}
			pa.Text = replacer.Replace(palabra.Header)
			pa.CallbackData = &res.Res[k].ID
			rows = append(rows, pa)
		}
		fmt.Println(rows)
		opts = tgb.NewInlineKeyboardMarkup(tgb.NewInlineKeyboardRow(rows...))
	}

	return
}
