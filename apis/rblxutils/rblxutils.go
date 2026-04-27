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

//go:embed packagemap.json
var PackageMap []byte

//go:embed joinPage.html
var JoinPage []byte

//InitApi initalizes the api and returns a function that can handle requests
func InitApi() func(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()
	mux.HandleFunc("/rblxutils/v1/package-map/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		w.Write(PackageMap)
	})

	mux.HandleFunc("/rblxutils/join/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "text/html")
		w.Write(JoinPage)
	})

	return mux.ServeHTTP
}