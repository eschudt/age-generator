package main

import (
	"encoding/json"
	"net/http"
)

type Hello struct {
	Message string
}

type Age struct {
	Age int
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/age", ageHandler)
	http.ListenAndServe(":8080", mux)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	data := Hello{"Hello World"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Write(resp)
}

func ageHandler(w http.ResponseWriter, r *http.Request) {
	data := Age{35}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Write(resp)
}