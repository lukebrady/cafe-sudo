package main

import (
	"net/http"
)

func indexRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write([]byte("Welcome to Cafe Sudo!\n"))
		w.WriteHeader(http.StatusOK)
	}
}

func listCommands(w http.ResponseWriter, r *http.Request) {
	commandList := "1. Grow some beans\n2. Check bean status\n3. Roast some beans\n"
	if r.Method == "GET" {
		w.Write([]byte(commandList))
		w.WriteHeader(http.StatusOK)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexRoute)
	mux.HandleFunc("/commands", listCommands)
	http.ListenAndServe(":4000", mux)
}
