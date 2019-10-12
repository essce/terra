package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(`{"version":"1.0"}`))
	})
	http.ListenAndServe(":8100", nil)
	log.Println("initialize")
}
