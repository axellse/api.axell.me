package rblxutils

import (
	_ "embed"
	"net/http"
)

//go:embed packagemap.json
var PackageMap []byte

func HandleEndpoints(w http.ResponseWriter, r *http.Request) {
	if r.PathValue("path") == "v1/package-map" || r.PathValue("path") == "v1/package-map/" {
		w.Header().Add("content-type", "application/json")
		w.Write(PackageMap)
	}
}