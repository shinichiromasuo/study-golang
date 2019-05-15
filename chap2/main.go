package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"net/http/fcgi"
	"strings"
)

type User struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Password string      `json:"password"`
	Array    interface{} `json:"array"`
}

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		return
	}
	http.HandleFunc("/wep", handlerMain)
	fcgi.Serve(l, nil)
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

			fmt.Fprint(w, xform)

			tmp := template.Must(template.ParseFiles("/var/opt/index.html"))
			tmp.ExecuteTemplate(w, "index.html", xform)

			return
		}

	}

	w.Write([]byte("Other than GET,POST Request \n"))
	log.Printf("Other than GET,POST Request")
}
