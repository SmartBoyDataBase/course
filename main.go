package main

import (
	"course/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", handler.PingPongHandler)
	http.HandleFunc("/course", handler.Handler)
	http.HandleFunc("/courses", handler.AllHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
