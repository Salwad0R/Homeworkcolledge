package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

var args = os.Args[1:]

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")

		err := json.NewEncoder(w).Encode(args)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	if r.Method == "PUT" || r.Method == "POST" {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		pattern := regexp.MustCompile(`\s+`)
		newArgs := pattern.Split(string(reqBody), -1)

		args = append(args, newArgs...)
		w.WriteHeader(http.StatusNoContent)
	}
}

func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8099", nil))
}