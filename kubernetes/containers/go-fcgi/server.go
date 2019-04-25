package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/fcgi"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("hc")
	fmt.Fprintf(w, "hc")
}

func dbHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("db_hc")
	fmt.Fprintf(w, "db_hc")
}

func main() {
	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		return
	}
	http.HandleFunc("/wep/systems.health.check", handler)
	http.HandleFunc("/wep/systems.db.health.check", dbHandler)
	fcgi.Serve(l, nil)
}
