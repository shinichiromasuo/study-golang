package main

import (
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/main", handlerOnlyPost)
	http.ListenAndServe(":8080", nil)
}

func handlerOnlyPost(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		w.Write([]byte("get request \n"))
		log.Printf("post request")
		return
	}

	if r.Method == http.MethodPost {
		w.Write([]byte("post request \n"))
		log.Printf("get request")
		return
	}

	w.Write([]byte("Other than GET,POST Request \n"))
	log.Printf("Other than GET,POST Request")
}
