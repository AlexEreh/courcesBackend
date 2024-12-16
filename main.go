package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		data, _ := json.Marshal(struct{ Message string }{"Hello!"})
		_, _ = io.WriteString(w, string(data))
	})
	_ = http.ListenAndServe("localhost:8080", mux)
}
