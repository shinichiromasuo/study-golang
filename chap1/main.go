package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func main() {
	http.Handle("/asset/", http.FileServer(http.Dir("../")))
	http.HandleFunc("/main", handlerMain)
	http.ListenAndServe(":8080", nil)
}

func handlerMain(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		w.Write([]byte("get request \n"))
		log.Printf("post request")
		return
	}

	if r.Method == http.MethodPost {

		if r.Header.Get("Content-Type") == "application/json" {
			dec := json.NewDecoder(r.Body)
			var reqJson map[string]interface{}
			dec.Decode(&reqJson)
			fmt.Printf("%v\n", reqJson)

			return
		}

		if strings.Contains(r.Header.Get("Content-type"), "multipart/form-data") {
			mulForm := r.FormValue("aaa")
			fmt.Println(mulForm)

			return
		}

		if r.Header.Get("Content-type") == "application/x-www-form-urlencoded" {
			xform := r.FormValue("aaa")
			fmt.Println(xform)

			return
		}

	}

	w.Write([]byte("Other than GET,POST Request \n"))
	log.Printf("Other than GET,POST Request")
}
