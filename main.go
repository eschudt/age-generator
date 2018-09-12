package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
)

type Hello struct {
	Message string
}

type Age struct {
	Age      int
	YearBorn int
}

func main() {
	listenPort := os.Getenv("NOMAD_PORT_http")

	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/age", ageHandler)
	http.ListenAndServe(":"+listenPort, mux)
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
	age := rand.Intn(99)
	data := Age{
		Age:      age,
		YearBorn: 2018 - age,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Write(resp)
}
