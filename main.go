package main

import (
    "log"
    "net/http"
	"time"
	"strconv"

	"github.com/gorilla/mux"

)

func Handler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	codeString := r.URL.Query().Get("code")
	code := 307
	if url == "" {
		url = "http://127.0.0.1"
	}
	if codeString != "" {
		i, err := strconv.Atoi(codeString)
		if err != nil {
			code = 307
		} else {
			code = i
		}
	}
	log.Printf("Redirecting to %s with code %d", url, code)
	http.Redirect(w, r, url, code)
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", Handler)

    srv := &http.Server{
        Handler:      r,
        Addr:         ":8080",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }
	log.Println("Starting redirect server")
    log.Fatal(srv.ListenAndServe())
}