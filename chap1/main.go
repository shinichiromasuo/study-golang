package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type User struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Password string      `json:"password"`
	Array    interface{} `json:"array"`
}

func main() {
	http.Handle("/asset/", http.FileServer(http.Dir("../")))
	http.HandleFunc("/main", handlerMain)
	http.ListenAndServe(":8080", nil)
}

func handlerMain(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		w.Write([]byte("get request \n"))
		log.Printf("get request")
		return
	}

	if r.Method == http.MethodPost {
		w.Write([]byte("post request \n"))
		log.Printf("post request")

		if r.Header.Get("Content-Type") == "application/json" {
			dec := json.NewDecoder(r.Body)
			var user User
			dec.Decode(&user)

			encoder := json.NewEncoder(w)
			encoder.SetIndent("", "  ")
			encoder.Encode(user)

			return
		}

		if strings.Contains(r.Header.Get("Content-type"), "multipart/form-data") {
			mulForm := r.FormValue("aaa")
			fmt.Println(mulForm)

			fmt.Fprint(w, mulForm)
			return
		}

		if r.Header.Get("Content-type") == "application/x-www-form-urlencoded" {
			xform := r.FormValue("aaa")
			fmt.Println(xform)

			tmp := template.Must(template.ParseFiles("../asset/index.html"))
			tmp.ExecuteTemplate(w, "index.html", xform)

			return
		}

	}

	w.Write([]byte("Other than GET,POST Request \n"))
	log.Printf("Other than GET,POST Request")
}
