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
	Header string // actual word
	ID     string // key to the definition
}

// RaeSearchResult array of words matching the search criteria
type RaeSearchResult struct {
	Approx int
	Res    []RaeWord
}

// RaeFunc Specify which function to use from the RAE API
type RaeFunc int8

const (
	nothing      RaeFunc = iota
	wordDay              // Word of the Day
	searchword           // Exact word search
	fetchDefByID         // Get the definition
	words                // Aproximate search (default on rae.es site)
)

const raeRestAPI = "https://dle.rae.es/data/"
const raeRestAuthHeader = "Basic cDY4MkpnaFMzOmFHZlVkQ2lFNDM0"
