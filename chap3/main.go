package main

import (
	"github.com/shinichiromasuo/study-golang/chap3/handler"
	"github.com/shinichiromasuo/study-golang/chap3/model"
	"net"
	"net/http"
	"net/http/fcgi"
)

func main() {
	model.DBinit()

	http.HandleFunc("/wep", handler.Handler)
	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		return
	}
	fcgi.Serve(l, nil)

}
