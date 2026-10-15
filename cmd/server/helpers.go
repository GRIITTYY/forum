package main

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"net/http"
)

type jsonResponse map[string]any

func writeJson(w http.ResponseWriter, data jsonResponse, status int) error {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	return json.MarshalWrite(w, data, jsontext.WithIndent("\t"))
}
