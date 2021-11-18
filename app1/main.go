package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"fmt"
)

type request struct {
	URL     string      `json:"url"`
	Method  string      `json:"method"`
	Headers http.Header `json:"headers"`
	Body    []byte      `json:"body"`
}

func handle(rw http.ResponseWriter, r *http.Request) {
	var err error
	rr := &request{}
	rr.Method = r.Method
	rr.Headers = r.Header
	rr.URL = r.URL.String()
	rr.Body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rrb, err := json.Marshal(rr)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(rrb)
	rw.Header().Set("Content-Type", "application/json")
	// rw.Write(rrb)
	rw.Write([]byte("New data with pipeline"))
}

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":80", nil)
}