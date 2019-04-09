
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


// RaeFunc Specify which function to use from the RAE API
type RaeFunc int8

const (
	nothing RaeFunc = iota
	wordDay
	searchword
	fetchDefByID
	words
)

const RAE_REST_API = "https://dle.rae.es/data/"
const RAE_REST_Auth_Header = "Basic cDY4MkpnaFMzOmFHZlVkQ2lFNDM0"
