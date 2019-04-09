package gorae

import (
	"net/http"
)

// Rae object to access the REST api
type rae struct {
	apiHTTP    string
	authHeader string
	client     *http.Client
	req        *http.Request
}

// RaeWord actual word, unique definition ID
type RaeWord struct {
	Word, ID string
}

// RaeSearchResult array of words matching the search criteria
type RaeSearchResult struct {
	Approx int
	Res    []RaeWord
}

// RaeFunc Specify which function to use from the RAE API
type RaeFunc int8

const (
	nothing RaeFunc = iota
	wordDay
	searchword
	fetchDefByID
	words
)

const raeRestAPI = "https://dle.rae.es/data/"
const raeRestAuthHeader = "Basic cDY4MkpnaFMzOmFHZlVkQ2lFNDM0"
