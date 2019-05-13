package main

import (
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/onlyPost", handlerOnlyPost)
	http.ListenAndServe(":8080", nil)
}

func handlerOnlyPost(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		w.Write([]byte("to do post only \n"))
		log.Printf("get request")
		return
	}

	w.Write([]byte("OK \n"))
	log.Printf("post request")
}
