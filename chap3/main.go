package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name  string
	Birth string
	Email string
	Tell  string
}

func main() {
	var err error
	db, err = bootDB(
		os.Getenv("APP_DB_ENDPOINT"),
		os.Getenv("APP_DB_USER"),
		os.Getenv("APP_DB_PASSWORD"),
	)
	if err != nil {
		log.Print(err)
		return
	}

	db.AutoMigrate(&User{})

	defer db.Close()

	http.HandleFunc("/wep", handlerMain)
	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		return
	}
	fcgi.Serve(l, nil)

}

func handlerMain(w http.ResponseWriter, r *http.Request) {

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
			var user User
			dec.Decode(&user)

			encoder := json.NewEncoder(w)
			encoder.SetIndent("", "  ")
			encoder.Encode(user)

			db.Create(&user)

			return
		}

		if strings.Contains(r.Header.Get("Content-type"), "multipart/form-data") {
			mulForm := r.FormValue("name")
			fmt.Println(mulForm)

			db.Create(&User{Name: r.FormValue("name"),
				Birth: r.FormValue("birth"),
				Email: r.FormValue("email"),
				Tell:  r.FormValue("tell")})

			return
		}

		if r.Header.Get("Content-type") == "application/x-www-form-urlencoded" {
			xform := r.FormValue("name")
			fmt.Println(xform)

			fmt.Fprint(w, xform)

			db.Create(&User{Name: r.FormValue("name"),
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

func bootDB(host, user, pass string) (*gorm.DB, error) {
	var err error

	driver := "mysql"
	protocol := "tcp"
	port := 3306
	name := "study_golang"
	args := "?charset=utf8&parseTime=True&loc=Local"

	con, err := gorm.Open(driver,
		fmt.Sprintf("%s:%s@%s([%s]:%d)/%s%s", user, pass, protocol, host, port, name, args),
	)
	if err != nil {
		return nil, err
	}

	con.DB().SetConnMaxLifetime(time.Second * 10)

	err = con.DB().Ping()
	if err != nil {
		return nil, err
	}

	return con, nil
}
