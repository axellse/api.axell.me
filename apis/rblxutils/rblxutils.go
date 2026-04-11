package rblxutils

import (
	_ "embed"
	"net/http"
)

//go:embed docs.html
var DocsFile []byte

//GetDocs returns the documentation file
func GetDocs() []byte {
	return DocsFile
}

//InitApi initalizes the api and returns a function that can handle requests
func InitApi() func(w http.ResponseWriter, r *http.Request) {
	return HandleEndpoints
}