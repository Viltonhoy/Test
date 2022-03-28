package main

import (
	"Test/handlers"
	"log"
	"net/http"
)

func main() {

	h := handlers.UserOne{
		Id:      3,
		Account: 400,
	}

	nh := handlers.TwoUsers{}

	http.HandleFunc("/add", h.Ref)
	http.HandleFunc("/del", h.Minus)
	http.HandleFunc("/rev", nh.Pere)

	port := ":9090"
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListernAndServe", err)
	}
}
