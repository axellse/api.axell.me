package main

import (
	_ "embed"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/axellse/api.axell.me/apis/rblxutils"
)

var RootUrl = "https://api.axell.me/"

//go:embed docs.html
var DocsTemplate []byte

//go:embed docs.css
var DocsCss []byte

var apis = map[string]func(w http.ResponseWriter, r *http.Request) {
	"rblxutils": rblxutils.InitApi(),
}

//go:embed index.html
var IndexRaw []byte
var Index = ComposeIndex(IndexRaw)

func ComposeDocs(docs []byte, name string) []byte {
	doc := strings.ReplaceAll(string(DocsTemplate), "&&NAME&&", name)
	doc = strings.ReplaceAll(doc, "&&URL&&", RootUrl + strings.ToLower(name) + "/")
	doc = strings.ReplaceAll(doc, "&&DOCS&&", string(docs))
	return []byte(doc)
}

func ComposeIndex(index []byte) []byte {
	apiList := ""
	for api, _ := range apis {
		apiList += "<li><a href=\"" + "/" + api + "/docs" +"\">" + api + "</a></li>\n"
	}
	return []byte(strings.ReplaceAll(string(index), "&&APIS&&", apiList))
}

var docs = map[string][]byte {
	"rblxutils": ComposeDocs(rblxutils.GetDocs(), "Rblxutils"),
}

func main() {
	http.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "text/html")
		w.Write(Index)
	})
	http.HandleFunc("/{api}/{path...}", func(w http.ResponseWriter, r *http.Request) {
		if r.PathValue("api") == "docs.css" {
			w.Header().Add("content-type", "text/css")
			w.Write(DocsCss)
			return
		}

		api, ok := apis[r.PathValue("api")]
		if !ok {
			w.WriteHeader(400)
			return
		}

		if r.PathValue("path") == "docs" {
			w.Header().Add("content-type", "text/html")
			w.Write(docs[r.PathValue("api")])
			return
		}

		api(w, r)
	})

	port := os.Getenv("RAY_PORT")
	if port == "" {
		port = "3001"
	}
	fmt.Println("starting server on", port)
	http.ListenAndServe("127.0.0.1:" + port, nil)
}