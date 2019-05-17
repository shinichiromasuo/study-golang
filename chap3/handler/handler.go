package handler

import (
	"encoding/json"
	"fmt"
	"github.com/shinichiromasuo/study-golang/chap3/model"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var db = model.DB

func Handler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		w.Write([]byte("get request \n"))
		log.Printf("get request")

		tmp := template.Must(template.ParseFiles("/var/opt/resistor.html"))
		tmp.ExecuteTemplate(w, "index.html", nil)

		return
	}

	if r.Method == http.MethodPost {
		w.Write([]byte("post request \n"))
		log.Printf("post request")

		if r.Header.Get("Content-Type") == "application/json" {
			dec := json.NewDecoder(r.Body)

			var user model.User
			dec.Decode(&user)

			encoder := json.NewEncoder(w)
			encoder.SetIndent("", "  ")
			encoder.Encode(user)

			model.Insert(user)
			return
		}

		if strings.Contains(r.Header.Get("Content-type"), "multipart/form-data") {
			mulForm := r.FormValue("name")
			fmt.Println(mulForm)

			model.Insert(model.User{Name: r.FormValue("name"),
				Birth: r.FormValue("birth"),
				Email: r.FormValue("email"),
				Tell:  r.FormValue("tell")})

			return
		}

		if r.Header.Get("Content-type") == "application/x-www-form-urlencoded" {
			xform := r.FormValue("name")
			fmt.Println(xform)

			fmt.Fprint(w, xform)

			model.Insert(model.User{Name: r.FormValue("name"),
				Birth: r.FormValue("birth"),
				Email: r.FormValue("email"),
				Tell:  r.FormValue("tell")})

			tmp := template.Must(template.ParseFiles("/var/opt/index.html"))
			tmp.ExecuteTemplate(w, "index.html", xform)

			return
		}

	}

	w.Write([]byte("Other than GET,POST Request \n"))
	log.Printf("Other than GET,POST Request")
}
